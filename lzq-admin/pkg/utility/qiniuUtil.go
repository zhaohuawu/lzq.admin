package utility

import (
	"bytes"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/**
 *
 * @author  糊涂的老知青
 * @date  2021/1/5 2:26 下午
* @version 1.0
*/

type QiniuConfig struct {
	QiNiuAccessKey string
	QiNiuSecretKey string
	QiNiuBucket    string
	QiNiuArea      string
	QiNiuBaseUrl   string
}

type qiniuUtil struct{}

func (qiniuUtil) UploadQiNiu(fileName string, localFilePath string, qiNiuConfig QiniuConfig) string {
	accessKey := qiNiuConfig.QiNiuAccessKey
	secretKey := qiNiuConfig.QiNiuSecretKey
	bucket := qiNiuConfig.QiNiuBucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = setQiNiuArea(qiNiuConfig.QiNiuArea)
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFilePath, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return ret.Key
}
func (qiniuUtil) UploadQiNiuByStream(fileName string, data []byte, qiNiuConfig QiniuConfig) (string, string, error) {
	accessKey := qiNiuConfig.QiNiuAccessKey
	secretKey := qiNiuConfig.QiNiuSecretKey
	bucket := qiNiuConfig.QiNiuBucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = setQiNiuArea(qiNiuConfig.QiNiuArea)
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}

	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, fileName, bytes.NewReader(data), dataLen, nil)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	// fmt.Println(fmt.Sprintf("QiNiu上传：Key:%v，Hash：%v，PersistentID：%v", ret.Key, ret.Hash, ret.PersistentID))
	return ret.Key, ret.Hash, nil

}
func setQiNiuArea(area string) *storage.Region {
	var zone *storage.Region
	switch {
	case area == "z0":
		zone = &storage.ZoneHuadong
	case area == "z1":
		zone = &storage.ZoneHuabei
	case area == "z2":
		zone = &storage.ZoneHuanan
	case area == "na0":
		zone = &storage.ZoneBeimei
	case area == "as0":
		zone = &storage.ZoneXinjiapo
	}
	return zone
}

func (qiniuUtil) DeleteFileList(fileNameList []string, qiNiuConfig QiniuConfig) bool {
	accessKey := qiNiuConfig.QiNiuAccessKey
	secretKey := qiNiuConfig.QiNiuSecretKey
	bucket := qiNiuConfig.QiNiuBucket
	successCount := 0
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = false
	bucketManager := storage.NewBucketManager(mac, &cfg)
	for _, fileName := range fileNameList {
		key := fileName
		err := bucketManager.Delete(bucket, key)
		if err != nil {
			fmt.Println(err)
			return false
		}

		successCount += 1
	}
	return true
}

func (qiniuUtil) DeleteFile(fileName string, qiNiuConfig QiniuConfig) int {
	accessKey := qiNiuConfig.QiNiuAccessKey
	secretKey := qiNiuConfig.QiNiuSecretKey
	bucket := qiNiuConfig.QiNiuBucket
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{}
	// 是否使用https域名
	cfg.UseHTTPS = false
	bucketManager := storage.NewBucketManager(mac, &cfg)
	key := fileName
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return -1
}

var QiniuUtil = &qiniuUtil{}
