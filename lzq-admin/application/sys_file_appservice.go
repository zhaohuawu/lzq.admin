package application

import (
	"github.com/gin-gonic/gin"
	"lzq-admin/domain/domainservice"
	"lzq-admin/domain/model"
)

/**
 * @Author  糊涂的老知青
 * @Date    2022/5/15
 * @Version 1.0.0
 */

type sysFileAppService struct {
	BaseAppService
}

var ISysFileAppService = sysFileAppService{}

// Upload
// @Summary 单个上传文件
// @Tags SysFile
// @Description
// @Accept mpfd
// @Produce  json
// @Param file formData file true "文件"
// @Success 200 {object} model.SystemFile " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysfile/upload [POST]
func (app *sysFileAppService) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	rep, err := domainservice.SysFileDomainService.Insert(file, *header)
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	//var rep model.SystemFile
	app.ResponseSuccess(c, rep)
}

// BatchUpload
// @Summary 批量上传文件
// @Tags SysFile
// @Description
// @Accept mpfd
// @Produce  json
// @Param files formData file true "文件"
// @Success 200 {array} model.SystemFile " "
// @Failure 500 {object} ResponseDto
// @Router /api/app/sysfile/batchUpload [POST]
func (app *sysFileAppService) BatchUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		app.ResponseError(c, err)
		return
	}
	files := form.File["files"]
	reps := make([]model.SystemFile, 0)
	for _, header := range files {
		f, err := header.Open()
		rep, err := domainservice.SysFileDomainService.Insert(f, *header)
		if err != nil {
			app.ResponseError(c, err)
			return
		}
		reps = append(reps, rep)
	}

	app.ResponseSuccess(c, reps)
}
