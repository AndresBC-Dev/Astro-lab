package service

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Configurar cliente S3
func UploadFileToS3(file *multipart.FileHeader, bucketName, key string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("error cargando configuración AWS: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	// Abrir archivo
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Leer archivo en memoria
	buffer := bytes.NewBuffer(nil)
	_, err = buffer.ReadFrom(src)
	if err != nil {
		return "", err
	}

	// Subir archivo a S3
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &key,
		Body:        bytes.NewReader(buffer.Bytes()),
		ContentType: &file.Header["Content-Type"][0],
	})

	if err != nil {
		return "", fmt.Errorf("error subiendo archivo a S3: %v", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	return url, nil
}
