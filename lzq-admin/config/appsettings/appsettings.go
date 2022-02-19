package appsettings

/**
 * @Author  糊涂的老知青
 * @Date    2021/10/30
 * @Version 1.0.0
 */

import (
	"encoding/json"
	"lzq-admin/pkg/hsflogger"
	"lzq-admin/pkg/utility"
	"os"
)

type Configuration struct {
	TenantInfos []TenantInfo
}

type TenantInfo struct {
	TenantId string
	Code     string
	Name     string
	Status   string
}

var AppSettings = &Configuration{}

func Init() {
	path, err := utility.GetRootPath()
	if err != nil {
		hsflogger.LogError("get root path err", err)
	}
	// 打开文件
	file, _ := os.Open(path + "/config/appsettings/appsettings.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	error := decoder.Decode(&AppSettings)
	if error != nil {
		hsflogger.LogError("AppsettingsInit_Error:", error)
	}
}
