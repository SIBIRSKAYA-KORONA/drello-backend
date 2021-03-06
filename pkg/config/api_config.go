package config

import (
	"fmt"
	"os"

	"github.com/go-park-mail-ru/2020_1_SIBIRSKAYA_KORONA/pkg/logger"
	"github.com/spf13/viper"
)

type ApiConfigController struct {
	origins           []string
	serverIp          string
	serverPort        uint
	db                string
	dbConnection      string
	grpcUserClient    string
	grpcSessionClient string
	s3Bucket          string
	s3BucketRegion    string
	tlsCrtPath        string
	tlsKeyPath        string
	metricURL         string
	service           string
	tmplPath          string
}

func CreateApiConfigController() *ApiConfigController {
	dbHost := viper.GetString("database.host")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	dbMode := viper.GetString("database.sslmode")

	bucket, exists := os.LookupEnv("S3_BUCKET")
	if !exists {
		logger.Fatal("S3_BUCKET environment variable not exist")
	}

	region, exists := os.LookupEnv("S3_BUCKET_REGION")
	if !exists {
		logger.Fatal("S3_BUCKET_REGION environment variable not exist")
	}

	credentialsDir, exists := os.LookupEnv("TLS_CREDENTIALS_DIR")
	if !exists {
		logger.Fatal("TLS_CREDENTIALS_DIR environment variable not exist")
	}

	keyFile, exists := os.LookupEnv("TLS_KEY_FILE")
	if !exists {
		logger.Fatal("TLS_KEY_FILE environment variable not exist")
	}

	crtFile, exists := os.LookupEnv("TLS_CRT_FILE")
	if !exists {
		logger.Fatal("TLS_CRT_FILE environment variable not exist")
	}

	tmplPath_, exists := os.LookupEnv("BOARD_TEMPLATES_PATH")
	if !exists {
		logger.Fatal("BOARD_TEMPLATES_PATH environment variable not exist")
	}

	return &ApiConfigController{
		origins:           viper.GetStringSlice("cors.allowed_origins"),
		serverIp:          viper.GetString("server.ip"),
		serverPort:        viper.GetUint("server.port"),
		db:                viper.GetString("database.dbms"),
		dbConnection:      fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPass, dbName, dbMode),
		grpcUserClient:    viper.GetString("grpc_clients.user"),
		grpcSessionClient: viper.GetString("grpc_clients.session"),
		s3Bucket:          bucket,
		s3BucketRegion:    region,
		tlsCrtPath:        fmt.Sprintf("%s/%s", credentialsDir, crtFile),
		tlsKeyPath:        fmt.Sprintf("%s/%s", credentialsDir, keyFile),
		metricURL:         viper.GetString("metrics.url"),
		service:           viper.GetString("metrics.service"),
		tmplPath:          tmplPath_,
	}
}

func (cc *ApiConfigController) GetOriginsSlice() []string {
	return cc.origins
}

func (cc *ApiConfigController) GetServerIP() string {
	return cc.serverIp
}

func (cc *ApiConfigController) GetServerPort() uint {
	return cc.serverPort
}

func (cc *ApiConfigController) GetDB() string {
	return cc.db
}

func (cc *ApiConfigController) GetDBConnection() string {
	return cc.dbConnection
}

func (cc *ApiConfigController) GetUserClient() string {
	return cc.grpcUserClient
}

func (cc *ApiConfigController) GetSessionClient() string {
	return cc.grpcSessionClient
}

func (cc *ApiConfigController) GetS3Bucket() string {
	return cc.s3Bucket
}

func (cc *ApiConfigController) GetS3BucketRegion() string {
	return cc.s3BucketRegion
}

func (cc *ApiConfigController) GetTLSCrtPath() string {
	return cc.tlsCrtPath
}

func (cc *ApiConfigController) GetTLSKeyPath() string {
	return cc.tlsKeyPath
}

func (cc *ApiConfigController) GetMetricsURL() string {
	return cc.metricURL
}

func (cc *ApiConfigController) GetServiceName() string {
	return cc.service
}

func (cc *ApiConfigController) GetTemplatesPath() string {
	return cc.tmplPath
}
