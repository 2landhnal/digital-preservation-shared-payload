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

type OAuthConfig struct {
	AccessKey string `mapstructure:"ACCESS_KEY_ID"`
	SecretKey string `mapstructure:"SECRET_ACCESS_KEY"`
}

type ServiceConfig struct {
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	HTTPServerPort int    `mapstructure:"HTTP_SERVER_PORT"`
	GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`
}

type NoSQLConfig struct {
	ServerAddress                   string `mapstructure:"SERVER_ADDRESS"`
	Database                        string `mapstructure:"DATABASE"`
	UserCollection                  string `mapstructure:"USER_COLLECTION"`
	VerifyEmailCollection           string `mapstructure:"VERIFY_EMAIL_COLLECTION"`
	SessionCollection               string `mapstructure:"SESSION_COLLECTION"`
	TmpObjectCollection             string `mapstructure:"TMP_OBJECT_COLLECTION"`
	ObjectCollection                string `mapstructure:"OBJECT_COLLECTION"`
	ObjectPermissionCollection      string `mapstructure:"OBJECT_PERMISSION_COLLECTION"`
	FolderPermissionCollection      string `mapstructure:"FOLDER_PERMISSION_COLLECTION"`
	ObjectRelationshipCollection    string `mapstructure:"OBJECT_RELATIONSHIP_COLLECTION"`
	FolderCollection                string `mapstructure:"FOLDER_COLLECTION"`
	EventCollection                 string `mapstructure:"EVENT_COLLECTION"`
	AgentCollection                 string `mapstructure:"AGENT_COLLECTION"`
	RightCollection                 string `mapstructure:"RIGHT_COLLECTION"`
	GroupCollection                 string `mapstructure:"GROUP_COLLECTION"`
	VirusScanResultCollection       string `mapstructure:"VIRUS_SCAN_RESULT_COLLECTION"`
	ExtractMetadataResultCollection string `mapstructure:"EXTRACT_METADATA_RESULT_COLLECTION"`
	EventObjectCollection           string `mapstructure:"EVENT_OBJECT_COLLECTION"`
	GroupUserCollection             string `mapstructure:"GROUP_USER_COLLECTION"`
}

type Cors struct {
	AllowOrigins []string `mapstructure:"ALLOW_ORIGINS"`
	AllowMethods []string `mapstructure:"ALLOW_METHODS"`
	AllowHeaders []string `mapstructure:"ALLOW_HEADERS"`
}

type Config struct {
	Environment   string      `mapstructure:"ENVIRONMENT"`
	DBSource      string      `mapstructure:"DB_SOURCE"`
	MongoDBConfig NoSQLConfig `mapstructure:"MONGO_DB_CONFIG"`
	HttpPort      int         `mapstructure:"HTTP_PORT"`
	GrpcPort      int         `mapstructure:"GRPC_PORT"`

	AuthService            ServiceConfig `mapstructure:"AUTH_SERVICE"`
	CoreService            ServiceConfig `mapstructure:"CORE_SERVICE"`
	CollectService         ServiceConfig `mapstructure:"COLLECT_SERVICE"`
	NotifyService          ServiceConfig `mapstructure:"NOTIFY_SERVICE"`
	MigrateService         ServiceConfig `mapstructure:"MIGRATE_SERVICE"`
	ConsumeSerivce         ServiceConfig `mapstructure:"CONSUME_SERVICE"`
	ScanVirussSerivce      ServiceConfig `mapstructure:"SCAN_VIRUSS_SERVICE"`
	ExtractMetadataSerivce ServiceConfig `mapstructure:"EXTRACT_METADATA_SERVICE"`

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

	OAuths            map[string]OAuthConfig `mapstructure:"OAUTH_CONFIGS"`
	OAuthKeyGoogle    string                 `mapstructure:"OAUTH_GOOGLE"`
	OAuthKeyDiscord   string                 `mapstructure:"OAUTH_DISCORD"`
	OAuthKeyFacebook  string                 `mapstructure:"OAUTH_FACEBOOK"`
	OAuthKeyGithub    string                 `mapstructure:"OAUTH_GITHUB"`
	OAuthKeyLinkedin  string                 `mapstructure:"OAUTH_LINKEDIN"`
	OAuthKeyMicrosoft string                 `mapstructure:"OAUTH_MICROSOFT"`
	OAuthServerURL    string                 `mapstructure:"OAUTH_SERVER_URL"`

	MQTopicDebeziumReplicateConfig       string `mapstructure:"MQ_TOPIC_DEBEZIUM_REPLICATE_CONFIG"`
	MQTopicDebeziumUser                  string `mapstructure:"MQ_TOPIC_DEBEZIUM_USER"`
	MQTopicDebeziumVerifyEmail           string `mapstructure:"MQ_TOPIC_DEBEZIUM_VERIFY_EMAIL"`
	MQTopicDebeziumSession               string `mapstructure:"MQ_TOPIC_DEBEZIUM_SESSION"`
	MQTopicDebeziumTmpObject             string `mapstructure:"MQ_TOPIC_DEBEZIUM_TMP_OBJECT"`
	MQTopicDebeziumObject                string `mapstructure:"MQ_TOPIC_DEBEZIUM_OBJECT"`
	MQTopicDebeziumObjectPermission      string `mapstructure:"MQ_TOPIC_DEBEZIUM_OBJECT_PERMISSION"`
	MQTopicDebeziumFolderPermission      string `mapstructure:"MQ_TOPIC_DEBEZIUM_FOLDER_PERMISSION"`
	MQTopicDebeziumObjectRelationship    string `mapstructure:"MQ_TOPIC_DEBEZIUM_OBJECT_RELATIONSHIP"`
	MQTopicDebeziumFolder                string `mapstructure:"MQ_TOPIC_DEBEZIUM_FOLDER"`
	MQTopicDebeziumEvent                 string `mapstructure:"MQ_TOPIC_DEBEZIUM_EVENT"`
	MQTopicDebeziumAgent                 string `mapstructure:"MQ_TOPIC_DEBEZIUM_AGENT"`
	MQTopicDebeziumRight                 string `mapstructure:"MQ_TOPIC_DEBEZIUM_RIGHT"`
	MQTopicDebeziumGroup                 string `mapstructure:"MQ_TOPIC_DEBEZIUM_GROUP"`
	MQTopicDebeziumVirusScanResult       string `mapstructure:"MQ_TOPIC_DEBEZIUM_VIRUS_SCAN_RESULT"`
	MQTopicDebeziumExtractMetadataResult string `mapstructure:"MQ_TOPIC_DEBEZIUM_EXTRACT_METADATA_RESULT"`
	MQTopicDebeziumEventObject           string `mapstructure:"MQ_TOPIC_DEBEZIUM_EVENT_OBJECT"`
	MQTopicDebeziumGroupUser             string `mapstructure:"MQ_TOPIC_DEBEZIUM_GROUP_USER"`

	S3Config      S3Config `mapstructure:"S3_CONFIG"`
	TmpFolderPath string   `mapstructure:"TMP_FOLDER_PATH"`

	// CORS configuration
	BackEndCors Cors `mapstructure:"BACKEND_CORS"`
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
