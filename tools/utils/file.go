package utils

import (
	"errors"
	"fmt"
	"github.com/pengcainiao/zero/core/env"
	"github.com/pengcainiao/zero/tools"
	"github.com/pengcainiao/zero/tools/syncer"
	"mime/multipart"
	"path"
	"strconv"
	"time"
)

func ByteToMB(i int) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", float64(i)/1024/1024), 64)
	return value
}

func UploadImage(file *multipart.FileHeader, filePath string) (url string, err error) {
	// 判断格式
	if !tools.InArray(path.Ext(file.Filename), []string{".gif", ".jpg", ".jpeg", ".png"}) {
		return "", errors.New("")
	}
	return UploadFile(file, filePath)
}

// file 文件信息
// filePath 文件路径，完整路径，非文件夹路径
func UploadFile(file *multipart.FileHeader, filePath string) (url string, err error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()
	bucket, err := syncer.Oss().Bucket(env.OssBucketName)
	if err != nil {
		return url, err
	}
	//s,_ :=os.Stat(filePath)
	//if s.IsDir() {
	//	filePath = fmt.Sprintf("%s/%s", filePath, path.Ext(file.Filename))
	//}
	// 开始上传文件
	err = bucket.PutObject(filePath, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s?t=%d", env.OssBucketHost, filePath, time.Now().Unix()), nil
}
