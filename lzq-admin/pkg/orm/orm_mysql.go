package orm

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"errors"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	_ "github.com/go-sql-driver/mysql"
	"lzq-admin/config"
	"lzq-admin/domain/model"
	"lzq-admin/middleware"
	"lzq-admin/pkg/hsflogger"
	"reflect"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var DB *xorm.Engine

func DatabaseInit() {
	var err error
	// 拼接连接字符串
	DB, err = xorm.NewEngine(config.Config.Database.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.Config.Database.UserName,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Database))
	if err != nil {
		hsflogger.LogError("Mysql 数据库连接失败", err)
	}
	if config.Config.Database.Type == "mysql" {
		DB.StoreEngine("ENGINE=InnoDB")
		DB.Charset("utf8mb4")
	}
	DB.SetColumnMapper(names.SameMapper{})
	// 设置打开数据库连接的最大数量
	DB.SetMaxOpenConns(config.Config.Database.MaxOpenConn)
	// 设置空闲连接池中连接的最大数量
	DB.SetMaxIdleConns(config.Config.Database.MaxIdleConn)
	// 设置了连接可复用的最大时间
	DB.SetConnMaxLifetime(4 * time.Hour)
	// 打印SQL
	DB.ShowSQL(true)
	// 迁移数据库
	if config.Config.IsMigration {
		migration()
	}

}

func migration() {
	// 自动迁移模式
	models := getModels()

	err := DB.Sync2(models...)
	if err != nil {
		fmt.Println(2, err)
	}
}

func getModels() []interface{} {
	return []interface{}{
		new(model.Tenant),
		new(model.SystemUser),
		new(model.AuthModule),
		new(model.AuthMenu),
		new(model.AuthPermission),
		new(model.AuthRole),
		new(model.AuthRolePermission),
	}
}

func BeginTrans() (*xorm.Session, error) {
	DbSession := DB.NewSession()
	defer DbSession.Close()
	return DbSession, DbSession.Begin()
}

// QSession 查询DB
func QSession(useMultiTenancy bool, tAlias ...string) *xorm.Session {
	queryDB := DB.Where("1=1")
	tableAlias := make([]string, 0)
	if len(tAlias) > 0 && tAlias[0] != "" {
		for _, v := range tAlias {
			tableAlias = append(tableAlias, v+".")
			queryDB.And(v+"."+"IsDeleted=?", 0)
		}
	} else {
		queryDB.And("IsDeleted=?", 0)
	}

	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	if useMultiTenancy {
		if len(tableAlias) > 0 {
			for _, v := range tableAlias {
				queryDB.Where(v+"TenantId=?", middleware.TokenClaims.TenantId)
			}

		} else {
			queryDB.Where("TenantId=?", middleware.TokenClaims.TenantId)
		}
	}
	return queryDB
}

// ISession 插入DBSession
func ISession() *xorm.Session {
	// 插入DB
	iBefore := func(obj interface{}) {
		obj = model.BeforeInsert(config.Config.ServerConfig.UseMultiTenancy, obj)

	}
	return DB.Before(iBefore)
}

// ISessionWithTrans 事务插入DBSession
func ISessionWithTrans(dbSession *xorm.Session) *xorm.Session {
	// 插入DB
	iBefore := func(obj interface{}) {
		obj = model.BeforeInsert(config.Config.ServerConfig.UseMultiTenancy, obj)
	}
	return dbSession.Before(iBefore)
}

func InsertWithCreateId(objs []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range objs {
		obj := model.BeforeInsert(config.Config.ServerConfig.UseMultiTenancy, v)
		result = append(result, obj)
	}
	return result
}

func UpdateWithModityId(objs []interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range objs {
		obj, _, _ := model.BeforeUpdate(v)
		result = append(result, obj)
	}
	return result
}

// USession 修改DB
func USession(useMultiTenancy bool) *xorm.Session {
	isModityId := false
	isModityTime := false
	// 修改DB
	uBefore := func(obj interface{}) {
		t := reflect.TypeOf(obj)
		if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
			obj, isModityId, isModityTime = model.BeforeUpdate(obj)
		}
	}
	updateDB := DB.Before(uBefore).Where("IsDeleted=?", 0)
	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	if useMultiTenancy {
		updateDB.And("TenantId=?", middleware.TokenClaims.TenantId).Omit("TenantId")
	}
	return updateDB.Omit("Id")
}

// USessionWithTrans 事务修改DBSession
func USessionWithTrans(useMultiTenancy bool, dbSession *xorm.Session) *xorm.Session {
	// 修改DB
	isModityId := false
	isModityTime := false
	uBefore := func(obj interface{}) {
		t := reflect.TypeOf(obj)
		if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
			obj, isModityId, isModityTime = model.BeforeUpdate(obj)
		}
	}
	dbSession.Before(uBefore).Where("IsDeleted=?", 0)
	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	if useMultiTenancy {
		dbSession.And("TenantId=?", middleware.TokenClaims.TenantId).Omit("TenantId")
	}
	return dbSession.Omit("Id")
}

// DSession 删除DB
func DSession(useMultiTenancy bool) *xorm.Session {
	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	isDeleterId := false
	isDeletionTime := false
	// 删除DB
	dBefore := func(obj interface{}) {
		obj, isDeleterId, isDeletionTime = model.BeforeDelete(obj)
	}
	deleteDB := DB.Before(dBefore).UseBool("IsDeleted").Where("IsDeleted=?", 0)
	if isDeleterId {
		deleteDB.Cols("DeleterId")
	}
	if isDeletionTime {
		deleteDB.Cols("DeletionTime")
	}
	if useMultiTenancy {
		deleteDB.Where("TenantId=?", middleware.TokenClaims.TenantId).Omit("TenantId")
	}
	return deleteDB.Omit("Id")
}

// DSessionWithTrans 事务删除DBSession
func DSessionWithTrans(useMultiTenancy bool, dbSession *xorm.Session) *xorm.Session {
	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	isDeleterId := false
	isDeletionTime := false
	// 删除DB
	dBefore := func(obj interface{}) {
		obj, isDeleterId, isDeletionTime = model.BeforeDelete(obj)
	}
	deleteDB := dbSession.Before(dBefore).UseBool("IsDeleted").Where("IsDeleted=?", 0)
	if isDeleterId {
		deleteDB.Cols("DeleterId")
	}
	if isDeletionTime {
		deleteDB.Cols("DeletionTime")
	}
	if useMultiTenancy {
		deleteDB.Where("TenantId=?", middleware.TokenClaims.TenantId).Omit("TenantId")
	}
	return deleteDB.Omit("Id")
}

// 是否使用多租户
func getUseMultiTenancy(useMultiTenancy bool) bool {
	if config.Config.ServerConfig.UseMultiTenancy && useMultiTenancy {
		useMultiTenancy = true
	} else {
		useMultiTenancy = false
	}
	return useMultiTenancy
}

// GetUpdateFields 得到需要修改的字段
func GetUpdateFields(obj interface{}, omitFields ...string) ([]string, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Check type error not Struct.")
	}
	omitFields = append(omitFields, "Id", "ID")
	fieldNum := t.NumField()
	var fieldNames = []string{"LastModificationTime", "LastModifierId"}
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Type.Kind() == reflect.Struct {
			if fieldC, err := GetUpdateFields(v.Field(i).Interface(), omitFields...); err != nil {
				return nil, err
			} else {
				fieldNames = append(fieldNames, fieldC...)
			}
		} else {
			fieldName := t.Field(i).Name
			isOmit := linq.From(omitFields).AnyWith(func(i interface{}) bool {
				return i == fieldName
			})
			if isOmit == false {
				fieldNames = append(fieldNames, fieldName)
			}
		}
	}
	return fieldNames, nil
}

// GetOptionFields 得到需要操作的字段
func GetOptionFields(obj interface{}, omitFields ...string) ([]string, error) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("Check type error not Struct.")
	}
	fieldNum := t.NumField()
	var fieldNames = make([]string, 0)
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Type.Kind() == reflect.Struct {
			if fieldC, err := GetUpdateFields(v.Field(i).Interface(), omitFields...); err != nil {
				return nil, err
			} else {
				fieldNames = append(fieldNames, fieldC...)
			}
		} else {
			fieldName := t.Field(i).Name
			isOmit := linq.From(omitFields).AnyWith(func(i interface{}) bool {
				return i == fieldName
			})
			if isOmit == false {
				fieldNames = append(fieldNames, fieldName)
			}
		}
	}
	return fieldNames, nil
}

func ConditionWithDeletedOrTenantId(useMultiTenancy bool, condition, tAlias string) string {
	condition = fmt.Sprintf("%v and %v.IsDeleted=%v", condition, tAlias, 0)
	useMultiTenancy = getUseMultiTenancy(useMultiTenancy)
	if useMultiTenancy {
		condition = fmt.Sprintf("%v and %v.TenantId=%v", condition, tAlias, middleware.TokenClaims.TenantId)
	}
	return condition
}
