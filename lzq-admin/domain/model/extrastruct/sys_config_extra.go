package extrastruct

/**
 * @Author  糊涂的老知青
 * @Date    2022/4/5
 * @Version 1.0.0
 */

type ExtraConfigObject struct {
	ExtraQiNiuConfig
}

type ExtraQiNiuConfig struct {
	AccessKey string `json:"accessKey" lab:"公钥"`  //公钥
	SecretKey string `json:"secretKey" lab:"私钥"`  //私钥
	Bucket    string `json:"bucket" lab:"空间名称"`   //空间名称
	Area      string `json:"area" lab:"存储区域"`     //存储区域
	BaseUrl   string `json:"baseUrl" lab:"域名"`    //域名
	Directory string `json:"directory" lab:"主目录"` //主目录
}
