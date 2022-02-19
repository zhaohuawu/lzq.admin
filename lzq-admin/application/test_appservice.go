package application

/**
 * @Author  糊涂的老知青
 * @Date    2021/11/30
 * @Version 1.0.0
 */

import (
	"github.com/gin-gonic/gin"
)

type TestAppService struct {
	BaseAppService
}

var ITestAppService = TestAppService{}

type TestDto struct {
	ID   string
	Name string
}

// TestReflectSetValue 测试接口
func (s *TestAppService) TestReflectSetValue(c *gin.Context) {
	var obj *TestDto
	//t := reflect.TypeOf(obj)
	//v := reflect.ValueOf(obj)

	//immutable := reflect.ValueOf(&obj).Elem()
	//fmt.Println("BeforeInsert-b", obj)
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println(t.Field(i).Name, v.Field(i).Interface())
	//	immutable.FieldByName("ID").Set(reflect.ValueOf("middleware.TokenClaims.Id"))
	//
	//}
	//hsflogger.LogInformation("BeforeInsert-a",obj)

	s.ResponseSuccess(c, obj)
}

func BeforeInsert(obj interface{}) {

}
