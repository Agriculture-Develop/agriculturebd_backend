package upload

import (
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func CheckFile(file *multipart.FileHeader) error {
	if file == nil {
		return errors.New("文件为空")
	}

	// 检查文件大小
	if int(file.Size) > 1024*1024*config.Get().File.MaxSize {
		return errors.New("文件大小超过2MB")
	}

	// 打开并读取前 512 字节判断 MIME 类型
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buf := make([]byte, 512)
	if _, err := src.Read(buf); err != nil {
		return err
	}

	// 文件类型校验
	contentType := http.DetectContentType(buf)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/jpg":  true,
		"image/png":  true,
	}

	if !allowedTypes[contentType] {
		return fmt.Errorf("不支持的文件类型：%s", contentType)
	}

	return nil
}

func UploadFile(file *multipart.FileHeader, types string) (string, error) {
	// 可根据项目路径结构构造目标存储路径
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	dir := filepath.Join(config.Get().File.Path, types)

	// 创建目标目录
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// 打开上传文件内容
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// 创建目标文件
	out, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 拷贝内容
	if _, err = io.Copy(out, src); err != nil {
		return "", err
	}

	return filename, nil
}

func DeleteFile(fileName, types string) error {
	baseDir := filepath.Join(config.Get().File.Path, types)
	filePath := filepath.Join(baseDir, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("删除文件失败：%w", err)
	}

	return nil
}
