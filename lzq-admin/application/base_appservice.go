package application

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"lzq-admin/domain/dto"
	"lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/utility"
	"net/http"
	"reflect"
	"strings"
	"xorm.io/xorm"
)

type BaseAppService struct {
}

type ResponseDto struct {
	Code int         `json:"code"`                    //1：请求成功， 0：请求业务错误
	Msg  string      `json:"msg"`                     //错误提示信息
	Data interface{} `json:"data" swaggered:"object"` //接口返回的业务数据
}

func (Base *BaseAppService) Response(c *gin.Context, code int, msg string, err error) {
	var res ResponseDto
	res.Code = code
	if err != nil {
		res.Msg = err.Error()
	} else {
		res.Msg = msg
	}
	c.JSON(http.StatusOK, res)
}
func (Base *BaseAppService) ResponseSuccess(c *gin.Context, obj interface{}) {
	//var res ResponseDto
	//res.Code = 1
	//res.Msg = "success"
	//res.Data = obj
	c.JSON(http.StatusOK, obj)
}
func (Base *BaseAppService) ResponseError(c *gin.Context, err error) {
	ResponseError(c, err)
}
func (Base *BaseAppService) ResponseSingleDto(c *gin.Context, obj1 interface{}, obj2 interface{}) {
	resultMap := utility.StructToMap(obj1, true)
	if err := mapstructure.Decode(resultMap, obj2); err != nil {
		Base.ResponseError(c, err)
		return
	}
	Base.ResponseSuccess(c, obj2)
	return
}

type PageListDto struct {
	TotalCount int64       `json:"totalCount"` //总条数
	Data       interface{} `json:"data"`       //数据
}
type PageParamsDto struct {
	RequireTotalCount bool   `form:"requireTotalCount"` //是否返回总条数
	Skip              int    `form:"skip"`              //第几页，>=1开始
	Take              int    `form:"take"`              //每页多少条数据
	Sort              string `form:"sort"`              //排序字段
	Filter            string `form:"filter"`            //查询条件
}
type Filter struct {
	Selector   string   `json:"selector"`
	Operator   string   `json:"operator"`
	Value      string   `json:"value"`
	OrSelector []Filter `json:"orSelector"`
}
type Sort struct {
	Selector string `json:"selector"`
	Desc     bool   `json:"desc"`
}

func DBCondition(inputDto PageParamsDto, dbSession *xorm.Session, tAlias string, structObj interface{}) error {
	// 整理dto对应的数据库字段
	tagMap := make(map[string]reflect.StructTag)
	reflectStruct(structObj, tagMap)
	// 条件
	var filters = make([]Filter, 0)
	if len(inputDto.Filter) > 0 {
		// [["name","contains"," 菜单管理"],["code","contains"," 菜单管理","or"]]
		var filterArray [][]string
		if err := json.Unmarshal([]byte(inputDto.Filter), &filterArray); err != nil {
			return err
		}

		var lastFilter Filter
		for i := 0; i < len(filterArray); {
			lastFilter = Filter{
				Selector: filterArray[i][0],
				Operator: filterArray[i][1],
				Value:    filterArray[i][2],
			}
			i++
			for i < len(filterArray) {
				if i > 0 && len(filterArray[i]) > 3 && filterArray[i][3] == "or" {
					lastFilter.OrSelector = append(lastFilter.OrSelector, Filter{
						Selector: filterArray[i][0],
						Operator: filterArray[i][1],
						Value:    filterArray[i][2],
					})
				} else {
					break
				}
				i++
			}
			filters = append(filters, lastFilter)
		}
		operatorMap := make(map[string]string, 0)
		operatorMap["="] = "="
		operatorMap["in"] = "in"
		operatorMap["not in"] = "not in"
		operatorMap[">"] = ">"
		operatorMap["<"] = "<"
		operatorMap[">="] = ">="
		operatorMap["<="] = "<="
		operatorMap["contains"] = "like"

		for _, f := range filters {
			operator := ""
			var isExist bool
			if operator, isExist = operatorMap[strings.ToLower(f.Operator)]; !isExist {
				return errors.New("不存在该条件操作符")
			}
			switch operator {
			case "in", "not in":
				if len(f.OrSelector) > 0 {
					var v []interface{}
					wStr := fmt.Sprintf("%v %v (?)", sqlField(tagMap, tAlias, f.Selector), operator)
					sv := append(v, f.Value)
					v = sv
					if orStr, orV, err := operatorTree(tagMap, tAlias, f.OrSelector, operatorMap); err != nil {
						return err
					} else {
						wStr = wStr + orStr
						sv := append(v, orV...)
						v = sv
					}
					dbSession.And(wStr, v...)
				} else {
					dbSession.And(fmt.Sprintf("%v %v (?)", sqlField(tagMap, tAlias, f.Selector), operator), f.Value)
				}
			case "like":
				if len(f.OrSelector) > 0 {
					var v []interface{}
					wStr := fmt.Sprintf("%v %v ?", sqlField(tagMap, tAlias, f.Selector), operator)
					sv := append(v, "%"+f.Value+"%")
					v = sv
					if orStr, orV, err := operatorTree(tagMap, tAlias, f.OrSelector, operatorMap); err != nil {
						return err
					} else {
						wStr = wStr + orStr
						sv := append(v, orV...)
						v = sv
					}
					dbSession.And(wStr, v...)
				} else {
					dbSession.And(fmt.Sprintf("%v %v ?", sqlField(tagMap, tAlias, f.Selector), operator), "%"+f.Value+"%")
				}
			default:
				if len(f.OrSelector) > 0 {
					var v []interface{}
					wStr := fmt.Sprintf("%v %v ?", sqlField(tagMap, tAlias, f.Selector), operator)
					sv := append(v, f.Value)
					v = sv
					if orStr, orV, err := operatorTree(tagMap, tAlias, f.OrSelector, operatorMap); err != nil {
						return err
					} else {
						wStr = wStr + orStr
						sv := append(v, orV...)
						v = sv
					}
					dbSession.And(wStr, v...)
				} else {
					dbSession.And(fmt.Sprintf("%v %v ?", sqlField(tagMap, tAlias, f.Selector), operator), "%"+f.Value+"%")
				}
			}
		}
	}

	// 排序
	if inputDto.Sort != "" {
		var sorts []Sort
		if err := json.Unmarshal([]byte(inputDto.Sort), &sorts); err != nil {
			return err
		}
		for i := 0; i < len(sorts); i++ {
			if sorts[i].Desc {
				dbSession.Desc(sqlField(tagMap, tAlias, sorts[i].Selector))
			} else {
				dbSession.Asc(sqlField(tagMap, tAlias, sorts[i].Selector))
			}
		}
	}

	if inputDto.Take > 0 {
		dbSession.Limit(inputDto.Take, inputDto.Skip)
	}
	return nil
}
func reflectStruct(structObj interface{}, tagMap map[string]reflect.StructTag) {
	t := reflect.TypeOf(structObj)
	v := reflect.ValueOf(structObj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		hsflogger.LogError("Check type error not Struct", nil)
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		f := t.Field(i)
		tags := f.Tag
		if len(tags.Get("json")) > 0 {
			tagMap[tags.Get("json")] = tags
		} else if f.Type.Kind() == reflect.Struct && tags.Get("xorm") == "extends" {
			reflectStruct(v.Field(i).Interface(), tagMap)
		} else {
			tagMap[f.Name] = tags
		}
	}
}
func operatorTree(tagMap map[string]reflect.StructTag, tAlias string, filters []Filter, operatorMap map[string]string) (string, []interface{}, error) {
	var wStr string
	var v []interface{}
	for _, f := range filters {
		operator := ""
		var isExist bool
		if operator, isExist = operatorMap[strings.ToLower(f.Operator)]; !isExist {
			return "", nil, errors.New("不存在该条件操作符")
		}
		switch operator {
		case "in", "not in":
			wStr = wStr + fmt.Sprintf(" or %v %v (?)", sqlField(tagMap, tAlias, f.Selector), operator)
			v = append(v, f.Value)
		case "like":
			wStr = wStr + fmt.Sprintf(" or %v %v ?", sqlField(tagMap, tAlias, f.Selector), operator)
			v = append(v, "%"+f.Value+"%")
		default:
			wStr = fmt.Sprintf(" or %v %v ?", sqlField(tagMap, tAlias, f.Selector), operator)
			v = append(v, f.Value)
		}
	}
	return wStr, v, nil
}
func sqlField(tagMap map[string]reflect.StructTag, tAlias string, field string) string {
	if v, t := tagMap[field]; t {
		f := field
		if len(v.Get("tField")) > 0 {
			f = v.Get("tField")
		}
		if len(v.Get("tAlias")) > 0 {
			f = fmt.Sprintf("%v.%v", v.Get("tAlias"), f)
		} else if len(tAlias) > 0 {
			f = fmt.Sprintf("%v.%v", tAlias, f)
		}
		return f
	} else {
		return field
	}
}

func ResponseError(c *gin.Context, err error) {
	var res ResponseDto
	res.Code = 0
	res.Msg = err.Error()
	c.JSON(http.StatusInternalServerError, res)
	return
}

func GetCurrentUserGrantedOperation(operations []dto.OperationDto) []dto.OperationDto {
	//userId:=middleware.TokenClaims.Id
	return operations
}
