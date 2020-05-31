package handler

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"github.com/thhy/ginblog/conf"
	"github.com/thhy/ginblog/logger"
)

//StaticFiles static file server
func StaticFiles(c *gin.Context) {
	url := c.Request.URL.Path
	prefix := path.Base(url)
	var requestPath string
	if prefix == "css" {
		requestPath = "static/css"
	} else if prefix == "js" {
		requestPath = "static/js"
	} else if prefix == "html" {
		requestPath = "static/html"
	}
	c.String(http.StatusOK, requestPath)
}

//ConvertToString 编码转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//Upload single file upload
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Log(logger.ERROR, err)
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	filename := ConvertToString(file.Filename, "gbk", "utf-8")
	uploadFilePath := conf.FILEPATH

	out, err := os.Create(uploadFilePath + filename)
	if err != nil {
		logger.Log(logger.ERROR, err)
		c.String(http.StatusCreated, "upload successful")
		return
	}
	defer out.Close()
	ioReader, _ := file.Open()
	io.Copy(out, ioReader)
	c.String(http.StatusCreated, "upload successful")
	// 单文件
	// file, err := c.FormFile("file")
	// if err != nil {
	// 	logger.Log(logger.ERROR, err)
	// 	c.String(http.StatusBadRequest, fmt.Sprintln(" upload failed"))
	// 	return
	// }
	// log.Println(file.Filename)

	// // 上传文件到指定的路径
	// // c.SaveUploadedFile(file, dst)

	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//UploadMultiFile upload multi files
func UploadMultiFile(c *gin.Context) {

}
