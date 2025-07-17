package service

import (
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/infrastructure/utils/upload"
	"mime/multipart"
	"os"
)

type IUploadSvc interface {
	UploadFile(file *multipart.FileHeader, types string) (string, error)
	UploadFiles(files []*multipart.FileHeader, types string) ([]string, error)
	DeleteFile(filePath, types string) error
}

type UploadSvc struct {
}

func NewUploadSvc() IUploadSvc {
	return &UploadSvc{}
}

func (s *UploadSvc) UploadFile(file *multipart.FileHeader, types string) (string, error) {
	// 文件检查
	if err := upload.CheckFile(file); err != nil {
		return "", err
	}

	// 上传文件
	path, err := upload.UploadFile(file, types)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (s *UploadSvc) UploadFiles(files []*multipart.FileHeader, types string) ([]string, error) {
	var paths []string

	for _, file := range files {
		if err := upload.CheckFile(file); err != nil {
			return nil, fmt.Errorf("文件 %s 检查失败：%w", file.Filename, err)
		}

		path, err := upload.UploadFile(file, types)
		if err != nil {
			return nil, fmt.Errorf("文件 %s 上传失败：%w", file.Filename, err)
		}

		paths = append(paths, path)
	}

	return paths, nil
}

func (s *UploadSvc) DeleteFile(filePath, types string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("文件不存在：%s", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("删除文件失败：%w", err)
	}

	return nil
}
