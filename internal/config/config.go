package config

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Environment string
		Mongo       MongoConfig
		Redis       RedisConfig
		Auth        AuthConfig
		Http        HttpConfig
		Limiter     LimiterConfig
	}

	MongoConfig struct {
		URI      string
		User     string
		Password string
		Name     string `mapstructure:"databaseName"`
	}

	RedisConfig struct {
		Host     string `mapstructure:"Host"`
		Port     string `mapstructure:"Port"`
		DB       int    `mapstructure:"DB"`
		Password string
	}

	AuthConfig struct {
		JWT                    JWTConfig
		Bcrypt                 BcryptConfig
		VerificationCodeLength int `mapstructure:"verificationCodeLength"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		Key             string
	}

	BcryptConfig struct {
		MinCost     int
		DefaultCost int
		MaxCost     int
	}

	HttpConfig struct {
		Domain             string        `mapstructure:"domain"`
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	LimiterConfig struct {
		RPS   int           `mapstructure:"rps"`
		Burst int           `mapstructure:"burst"`
		TTL   time.Duration `mapstructure:"ttl"`
	}
)

func Init(configDir string) (*Config, error) {
	if err := parseConfigFile(configDir, os.Getenv("APP_ENV")); err != nil {
		return nil, err
	}

	var conf Config
	if err := unmarhal(&conf); err != nil {
		return nil, err
	}
	if err := setFromEnv(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func parseConfigFile(folder, env string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("dev")

	if env == "production" {
		viper.SetConfigName("prod")
	}

	return viper.MergeInConfig()
}

func unmarhal(conf *Config) error {
	if err := viper.UnmarshalKey("mongo", &conf.Mongo); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("redis", &conf.Redis); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("http", &conf.Http); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("limiter", &conf.Limiter); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth", &conf.Auth.JWT); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("auth.verificationCodeLength", &conf.Auth.VerificationCodeLength); err != nil {
		return err
	}

	return nil
}

func setFromEnv(conf *Config) error {
	if err := envconfig.Process("mongo", &conf.Mongo); err != nil {
		return err
	}
	if err := envconfig.Process("redis", &conf.Redis); err != nil {
		return err
	}
	if err := envconfig.Process("http", &conf.Http); err != nil {
		return err
	}
	if err := envconfig.Process("jwt", &conf.Auth.JWT); err != nil {
		return err
	}
	if err := envconfig.Process("bcrypt", &conf.Auth.Bcrypt); err != nil {
		return err
	}
	conf.Environment = os.Getenv("APP_ENV")

	return nil
}
