package define

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"log"
)

type UserClaim struct {
	Id       int
	Identity string
	Username string
	jwt.StandardClaims
}

// CodeLength mail code length
var CodeLength = 6

// CodeExpireTime mail code expire time second
var CodeExpireTime = 10 * 60

// CosAddr Tencent COS addr
var CosAddr = "https://ryan-1254379222.cos.ap-beijing.myqcloud.com"

// PageSize file
var PageSize = 20

var DateTime = "YYYY-MM-DD HH:m:s"

// TokenExpireTime second
var TokenExpireTime = 60 * 60

var RefreshTokenExpireTime = 60 * 60 * 24 * 3

func GetJetKey() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    // optionally look for a config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	JetKey, ok := viper.Get("JWT_KEY").(string)
	if !ok {
		log.Fatalf("cat not read JWT_KEY")
	}
	return JetKey
}

func GetDatabaseInfo() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    // call multiple times to add many search paths
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	Database, ok := viper.Get("MYSQL").(string)
	if !ok {
		log.Fatalf("can not read MYSQL database")
	}
	return Database
}

func GetRedisInfo() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")    // call multiple times to add many search paths
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	rdb, ok := viper.Get("RDB").(string)
	if !ok {
		log.Fatalf("can not read RDB")
	}
	return rdb
}

func GetSentEmailInfo() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	SentEmailPassword, ok := viper.Get("MAIL_TOKEN").(string)
	if !ok {
		log.Fatalf("can not read MAIL_TOKEN")
	}

	return SentEmailPassword
}

func GetCosSecretId() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	CosSecretId, ok := viper.Get("COS_SECRET_ID").(string)
	if !ok {
		log.Fatalf("can not read COS_SECRET_ID")
	}

	return CosSecretId
}

func GetCosSecretKey() string {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Error reading config file, %s", err)
	}
	CosSecretKey, ok := viper.Get("COS_SECRET_KEY").(string)
	if !ok {
		log.Fatalf("can not read COS_SECRET_KEY")
	}

	return CosSecretKey
}
