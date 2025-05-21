package util

import (
	"time"

	"github.com/spf13/viper"
)

type Topic struct {
	Partition int32    `mapstructure:"PARTITION"`
	RF        int32    `mapstructure:"RF"`
	Groups    []string `mapstructure:"GROUPS"`
}

type S3Config struct {
	AccessKey string `mapstructure:"ACCESS_KEY_ID"`
	SecretKey string `mapstructure:"SECRET_ACCESS_KEY"`
	Region    string `mapstructure:"REGION"`
	Bucket    string `mapstructure:"BUCKET"`
}

type ServiceConfig struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	HTTPServerPort int    `mapstructure:"HTTP_SERVER_PORT"`
	GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`
}

type Config struct {
	Environment            string           `mapstructure:"ENVIRONMENT"`
	DBSource               string           `mapstructure:"DB_SOURCE"`
	MongoDBSource          string           `mapstructure:"MONGO_DB_SOURCE"`
	HttpPort               int              `mapstructure:"HTTP_PORT"`
	GrpcPort               int              `mapstructure:"GRPC_PORT"`
	CollectSerivce         ServiceConfig    `mapstructure:"COLLECT_SERVICE"`
	NotifySerivce          ServiceConfig    `mapstructure:"NOTIFY_SERVICE"`
	ConsumeSerivce         ServiceConfig    `mapstructure:"CONSUME_SERVICE"`
	ScanVirussSerivce      ServiceConfig    `mapstructure:"SCAN_VIRUSS_SERVICE"`
	ExtractMetadataSerivce ServiceConfig    `mapstructure:"EXTRACT_METADATA_SERVICE"`
	TokenSymmetricKey      string           `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration    time.Duration    `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   time.Duration    `mapstructure:"REFRESH_TOKEN_DURATION"`
	MigrationURL           string           `mapstructure:"MIGRATION_URL"`
	CacheAddress           string           `mapstructure:"CACHE_ADDRESS"`
	EmailSenderName        string           `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress     string           `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword    string           `mapstructure:"EMAIL_SENDER_PASSWORD"`
	AllowedOrigins         []string         `mapstructure:"ALLOWED_ORIGINS"`
	MQAddress              []string         `mapstructure:"MQ_ADDRESS"`
	MQTopics               map[string]Topic `mapstructure:"MQ_TOPICS"`
	MQTopicVerifyEmail     string           `mapstructure:"MQ_TOPIC_VERIFY_EMAIL"`
	MQTopicScanViruss      string           `mapstructure:"MQ_TOPIC_SCAN_VIRUSS"`
	MQTopicExtractMetadata string           `mapstructure:"MQ_TOPIC_EXTRACT_METADATA"`
	S3Config               S3Config         `mapstructure:"S3_CONFIG"`
	TmpFolderPath          string           `mapstructure:"TMP_FOLDER_PATH"`
}

func LoadConfig(path string, config *Config) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
