package domainservice

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io"
	"lzq-admin/domain/domainconsts"
	"lzq-admin/domain/model"
	"lzq-admin/domain/model/extrastruct"
	"lzq-admin/pkg/orm"
	"lzq-admin/pkg/utility"
	"mime/multipart"
	"path"
	"strings"
	"sync"
	"time"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/15
 * @Version 1.0.0
 */

type sysFileDomainService struct {
	wg sync.WaitGroup
}

var SysFileDomainService = sysFileDomainService{}

func (s *sysFileDomainService) Insert(file multipart.File, header multipart.FileHeader) (model.SystemFile, error) {
	var entity model.SystemFile
	entity.ID = utility.UuidCreate()
	entity.OriginalName = header.Filename
	entity.Size = header.Size
	entity.ContentType = header.Header["Content-Type"][0]
	entity.Extension = strings.ToLower(path.Ext(header.Filename)) // 去文件后缀
	entity.NewName = entity.ID + entity.Extension
	entity.Status = domainconsts.SysFileStatusUnused
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return entity, err
	}
	if config, err := SysConfigDomainService.GetSysConfigCacheByCode(model.ExtraQiNiuConfig, "QiNiuStorage"); err != nil {
		return entity, err
	} else {
		// 上传七牛云
		var qnconfig extrastruct.ExtraQiNiuConfig
		if err := jsoniter.UnmarshalFromString(config, &qnconfig); err != nil {
			return entity, err
		}
		fileUrl := utility.UrlJoint(qnconfig.Directory, time.Now().Format("20060102"), entity.NewName)
		entity.Url = strings.TrimPrefix(fileUrl, qnconfig.Directory)
		_, entity.ThirdPartyId, err = utility.QiniuUtil.UploadQiNiuByStream(fileUrl, buf.Bytes(),
			utility.QiniuConfig{
				QiNiuBucket:    qnconfig.Bucket,
				QiNiuSecretKey: qnconfig.SecretKey,
				QiNiuAccessKey: qnconfig.AccessKey,
				QiNiuArea:      qnconfig.Area,
				QiNiuBaseUrl:   qnconfig.BaseUrl,
			})
		if err != nil {
			return entity, err
		}
	}

	if _, err := orm.ISession().Insert(&entity); err != nil {
		return model.SystemFile{}, err
	}
	return entity, nil
}

func (s *sysFileDomainService) UpdateStatus(id, status string) (model.SystemFile, error) {
	var entity model.SystemFile
	isExit, err := orm.QSession(true).ID(id).Get(&entity)
	if err != nil {
		return entity, err
	}
	if !isExit {
		return entity, errors.New("该文件不已存在")
	}
	entity.Status = status
	if status == domainconsts.SysFileStatusDeleted {
		if _, err := orm.DSession(true).ID(entity.ID).Update(&entity); err != nil {
			return model.SystemFile{}, err
		}
	} else {
		if _, err := orm.USession(true).ID(entity.ID).Update(&entity); err != nil {
			return model.SystemFile{}, err
		}
	}
	return entity, nil
}
