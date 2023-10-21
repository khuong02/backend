package minio

import (
	"context"
	"fmt"
	"github.com/khuong02/backend/pkg/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	//"repository.centralretail.com.vn/application/go-pkg/logger"
	//"repository.centralretail.com.vn/application/integrate-go-app/go-app/pkg/codeerror"
	"strings"
)

type MinioConfig struct {
	PublicEndpoint string `env-required:"true" yaml:"PUBLIC_ENDPOINT" env:"MINIO_PUBLIC_ENDPOINT"`
	ServerEndpoint string `env-required:"true" yaml:"SERVER_ENDPOINT" env:"MINIO_SERVER_ENDPOINT"`
	AccessKey      string `env-required:"true" yaml:"ACCESS_KEY" env:"MINIO_ACCESS_KEY"`
	SecretKey      string `env-required:"true" yaml:"SECRET_KEY" env:"MINIO_SECRET_KEY"`
	UseSSL         bool   `env-required:"true" yaml:"USE_SSL" env:"MINIO_USE_SSL"`

	Location   string `yaml:"LOCATION" env:"MINIO_LOCATION"`
	BucketName string `yaml:"BUCKET_NAME" env:"MINIO_BUCKET_NAME"`
}

type MinioFPutObject struct {
	Name     string
	FilePath string
	Otps     minio.PutObjectOptions
}

type MinioPutObject struct {
	Name       string
	Reader     io.Reader
	ObjectSize int64
	Otps       minio.PutObjectOptions
}

type MinioClient struct {
	Client *minio.Client

	cfg    MinioConfig
	logger *logger.Logger

	bucketName string
	location   string
}

func NewMinioClient(cfg MinioConfig, logger *logger.Logger) *MinioClient {
	return &MinioClient{
		cfg:        cfg,
		logger:     logger,
		bucketName: cfg.BucketName,
		location:   cfg.Location,
	}
}

func (m *MinioClient) Connect() *MinioClient {
	minioClient, err := minio.New(m.cfg.ServerEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.cfg.AccessKey, m.cfg.SecretKey, ""),
		Secure: m.cfg.UseSSL,
	})
	if err != nil {
		m.logger.Error("Connect minio server fail", "err:", err)

		os.Exit(1)
	}

	m.Client = minioClient
	m.logger.Info("Connect minio server successfully")

	return m
}

func (m *MinioClient) SetBucketName(bucketName string) *MinioClient {
	m.bucketName = bucketName

	return m
}

func (m *MinioClient) SetLocation(location string) *MinioClient {
	m.location = location

	return m
}

func (m *MinioClient) CreateBucket(ctx context.Context) (*MinioClient, error) {
	if m.Client == nil {
		return nil, errors.New("Not connected to minio server")
	}

	if strings.TrimSpace(m.bucketName) == "" {
		return nil, errors.New("Doesn't exist bucket")
	}

	err := m.Client.MakeBucket(ctx, m.bucketName, minio.MakeBucketOptions{Region: m.location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := m.Client.BucketExists(ctx, m.bucketName)
		if errBucketExists == nil && exists {
			m.logger.Info(fmt.Sprintf("We already own %s", m.bucketName))

			return m, nil
		} else {
			log.Fatalln(err)
		}
	} else {
		m.logger.Info(fmt.Sprintf("Successfully created %s\n", m.bucketName))
	}

	return m, nil
}

func (m *MinioClient) FPutObject(ctx context.Context, obj MinioFPutObject) (*minio.UploadInfo, error) {
	info, err := m.Client.FPutObject(ctx, m.bucketName, obj.Name, obj.FilePath, obj.Otps)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (m *MinioClient) PutObject(ctx context.Context, obj MinioPutObject) (*minio.UploadInfo, error) {
	info, err := m.Client.PutObject(ctx, m.bucketName, obj.Name, obj.Reader, obj.ObjectSize, obj.Otps)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (m *MinioClient) MultipartFileHeaderToPutObj(ctx context.Context, file *multipart.FileHeader) (*MinioPutObject, error) {

	fileBuffer, err := file.Open()
	if err != nil {
		return nil, err
	}

	fileName := file.Filename
	contentType := file.Header["Content-Type"][0]
	size := file.Size

	return &MinioPutObject{
		Name:       fileName,
		Reader:     fileBuffer,
		ObjectSize: size,
		Otps: minio.PutObjectOptions{
			ContentType: contentType,
		},
	}, nil
}

func (m *MinioClient) UploadMedia(ctx context.Context, bucket string, file *multipart.FileHeader) (*minio.UploadInfo, error) {
	minioClient, err := m.SetBucketName(bucket).CreateBucket(ctx)
	if err != nil {
		m.logger.Error("create bucket error", "err:", err)

		return nil, ErrFileUpload(err)
	}

	configObj, err := minioClient.MultipartFileHeaderToPutObj(ctx, file)
	if err != nil {
		m.logger.Error("file error", "err:", err)

		return nil, ErrFileUpload(err)
	}

	info, err := m.PutObject(ctx, *configObj)
	if err != nil {
		m.logger.Error("upload file to minio error", "err:", err)

		return nil, ErrFileUploadMinio(err)
	}

	return info, nil
}
