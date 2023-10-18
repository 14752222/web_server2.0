package client

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

type FileManager struct {
	Enpoint   string
	Bucket    string
	AccessKey string
	SecretKey string
	UseSSL    bool //true 使用https ，false 使用 http
	Ctx       context.Context
	client    *minio.Client
}

type FileDesc struct {
	Bucket      string
	FileName    string
	FilePath    string
	ContentType string
}

var fileManagerInstance *FileManager

//var fileManagerClient *minio.Client

func init() {
	var err error
	fileManagerInstance = &FileManager{
		Enpoint:   os.Getenv("ENPOINT"),
		Bucket:    os.Getenv("MINIO_BUCKET"),
		AccessKey: os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey: os.Getenv("MINIO_SECRET_KEY"),
		UseSSL:    false,
		Ctx:       context.Background(),
	}
	options := &minio.Options{
		Creds:  credentials.NewStaticV4(fileManagerInstance.AccessKey, fileManagerInstance.SecretKey, ""),
		Secure: fileManagerInstance.UseSSL,
	}

	fileManagerClient, err := minio.New(fileManagerInstance.Enpoint, options)
	if err != nil {
		fmt.Println("minio client Error: ", err)
	}

	fileManagerInstance.client = fileManagerClient
}

func GetfileManagerInstance() *FileManager {
	return fileManagerInstance
}

func (fm *FileManager) HasBucket(bucketName string) (bool, error) {
	exists, errBucketExists := fm.client.BucketExists(fm.Ctx, bucketName)
	if errBucketExists == nil && exists {
		return false, nil
	}
	if errBucketExists != nil {
		return false, errBucketExists
	}
	return true, nil
}

//CreateBucket
/**
* @method  CreateBucket 创建bucket
* @param {bucketName} string bucket名称
* @param {location} string bucket存储位置
* @brief  创建bucket
* @return {error} 返回错误信息
**/
func (fm *FileManager) CreateBucket(bucketName, location string) error {
	return fm.client.MakeBucket(fm.Ctx, bucketName, minio.MakeBucketOptions{Region: location})
}

//UploadFile
/**
* @method UploadFile 上传文件
* @param  filedesc *FileDesc
* @brief  Upload file to minio server
**/
func (fm *FileManager) UploadFile(filedesc *FileDesc) (minio.UploadInfo, error) {
	contentType := minio.PutObjectOptions{
		ContentType: filedesc.ContentType,
	}
	fileInfo, err := fm.client.
		FPutObject(fm.Ctx, filedesc.Bucket, filedesc.FileName, filedesc.FilePath, contentType)
	if err != nil {
		fmt.Println("upload info Error: ", err)
		return fileInfo, err
	}
	return fileInfo, err
}

func (fm *FileManager) GetFileInfo(bucketName, fileName string) (minio.ObjectInfo, error) {
	return fm.client.StatObject(fm.Ctx, bucketName, fileName, minio.StatObjectOptions{})
}

// RemoveFile 删除文件
func (fm *FileManager) RemoveFile(bucketName, fileName string) error {
	return fm.client.RemoveObject(fm.Ctx, bucketName, fileName, minio.RemoveObjectOptions{})
}
