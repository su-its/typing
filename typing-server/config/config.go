package config

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Environment string
	DBAddr      string
	DBUser      string
	DBPassword  string
	DBName      string
	Location    *time.Location
	MySQLConfig *mysql.Config
}

func New() (*Config, error) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "production"
	}

	dbAddr := os.Getenv("DB_ADDR")
	if dbAddr == "" {
		dbAddr = "db:3306" // docker-composeでのデフォルト
	}

	dbUser := "user"
	dbPassword := "password"
	dbName := "typing-db"

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	mysqlCfg := &mysql.Config{
		DBName:    dbName,
		User:      dbUser,
		Passwd:    dbPassword,
		Net:       "tcp",
		Addr:      dbAddr,
		ParseTime: true,
		Loc:       jst,
	}

	return &Config{
		Environment: environment,
		DBAddr:      dbAddr,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBName:      dbName,
		Location:    jst,
		MySQLConfig: mysqlCfg,
	}, nil
}

// GetMySQLDSN はMySQL接続用のDSN文字列を返します。
func (c *Config) GetMySQLDSN() string {
	return c.MySQLConfig.FormatDSN()
}

// GetLocation は設定されたタイムゾーン Location を返します。
func (c *Config) GetLocation() *time.Location {
	return c.Location
}
