package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

const (
	defaultHTTPPort               = "5023"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1

	defaultPostgresPort = "5432"
	defaultPostgresHost = "localhost"

	defaultLimiterRPS   = 10
	defaultLimiterBurst = 2
	defaultLimiterTTL   = 10 * time.Minute
)

type Mode string

const (
	DEVELOPMENT Mode = "development"
	PRODUCTION  Mode = "production"
)

type (
	Config struct {
		HTTP     HTTPConfig
		Database DatabaseConfig
		Limiter  LimiterConfig
		CacheTTL time.Duration `mapstructure:"ttl"`
	}

	HTTPConfig struct {
		Port               string        `mapstructure:"port"`
		Timeout            TimeoutConfig `mapstructure:"timeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderMegabytes"`
	}

	DatabaseConfig struct {
		Postgres PostgresConfig `mapstructure:"postgres"`
	}

	PostgresConfig struct {
		User     string `mapstructure:"user"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"dbName"`
		Password string `mapstructure:"password"`
		SSLMode  string `mapstructure:"sslMode"`
	}

	TimeoutConfig struct {
		Read  time.Duration `mapstructure:"read"`
		Write time.Duration `mapstructure:"write"`
	}

	LimiterConfig struct {
		RPS   int           `mapstructure:"rps"`
		Burst int           `mapstructure:"burst"`
		TTL   time.Duration `mapstructure:"ttl"`
	}
)

func Init(cfgDir string, mode Mode) (*Config, error) {
	setupDefaultValues()

	if err := parseConfigFile(cfgDir, mode); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshall(&cfg); err != nil {
		return nil, err
	}

	parseEnvFile(&cfg)

	return &cfg, nil
}

func parseConfigFile(folder string, mode Mode) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName(string(mode))

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func unmarshall(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("database", &cfg.Database); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("limiter", &cfg.Limiter); err != nil {
		return err
	}
	return viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL)
}

func parseEnvFile(cfg *Config) {
	cfg.Database.Postgres.Password = os.Getenv("DB_POSTGRES_PASSWORD")
}

func setupDefaultValues() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderMegabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)

	viper.SetDefault("limiter.rps", defaultLimiterRPS)
	viper.SetDefault("limiter.ttl", defaultLimiterTTL)
	viper.SetDefault("limiter.burst", defaultLimiterBurst)

	viper.SetDefault("database.postgres.host", defaultPostgresHost)
	viper.SetDefault("database.postgres.port", defaultPostgresPort)
}
