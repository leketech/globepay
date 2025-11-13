package config

import (
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// Config holds application configuration
type Config struct {
	Environment   string              `mapstructure:"environment"`
	ServiceName   string              `mapstructure:"service_name"`
	Version       string              `mapstructure:"version"`
	Server        ServerConfig        `mapstructure:"server"`
	Database      DatabaseConfig      `mapstructure:"database"`
	Redis         RedisConfig         `mapstructure:"redis"`
	JWT           JWTConfig           `mapstructure:"jwt"`
	AWS           AWSConfig           `mapstructure:"aws"`
	Observability ObservabilityConfig `mapstructure:"observability"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"db_name"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	MigrationsPath  string        `mapstructure:"migrations_path"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string        `mapstructure:"secret"`
	Expiration time.Duration `mapstructure:"expiration"`
}

type AWSConfig struct {
	Region       string `mapstructure:"region"`
	S3Bucket     string `mapstructure:"s3_bucket"`
	SQSQueueURL  string `mapstructure:"sqs_queue_url"`
	SESFromEmail string `mapstructure:"ses_from_email"`
}

type ObservabilityConfig struct {
	PrometheusPort int    `mapstructure:"prometheus_port"`
	JaegerEndpoint string `mapstructure:"jaeger_endpoint"`
	LogLevel       string `mapstructure:"log_level"`
	EnableTracing  bool   `mapstructure:"enable_tracing"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Enable environment variable override
	viper.AutomaticEnv()

	// Set defaults
	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("environment", "development")
	viper.SetDefault("service_name", "globepay-api")
	viper.SetDefault("version", "0.0.0")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", 30*time.Second)
	viper.SetDefault("server.write_timeout", 30*time.Second)
	viper.SetDefault("server.idle_timeout", 120*time.Second)
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 5)
	viper.SetDefault("database.conn_max_lifetime", 5*time.Minute)
	viper.SetDefault("jwt.expiration", 24*time.Hour)
}

// SetConfigFile sets the config file to use
func SetConfigFile(file string) {
	viper.SetConfigFile(file)
}

// IsDevelopment checks if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsStaging checks if the environment is staging
func (c *Config) IsStaging() bool {
	return c.Environment == "staging"
}

// IsProduction checks if the environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// GetDatabaseDSN returns the database DSN (Data Source Name)
func (c *Config) GetDatabaseDSN() string {
	return "host=" + c.Database.Host +
		" port=" + strconv.Itoa(c.Database.Port) +
		" user=" + c.Database.User +
		" password=" + c.Database.Password +
		" dbname=" + c.Database.DBName +
		" sslmode=" + c.Database.SSLMode
}

// GetRedisAddress returns the Redis address
func (c *Config) GetRedisAddress() string {
	return c.Redis.Host + ":" + strconv.Itoa(c.Redis.Port)
}

// GetServerAddress returns the server address
func (c *Config) GetServerAddress() string {
	return strconv.Itoa(c.Server.Port)
}

// GetJWTExpirySeconds returns the JWT expiry in seconds
func (c *Config) GetJWTExpirySeconds() int {
	return int(c.JWT.Expiration.Seconds())
}
