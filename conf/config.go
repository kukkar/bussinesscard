package config

import (
	"errors"
	"fmt"

	"github.com/kukkar/common-golang/pkg/config"
	"github.com/kukkar/common-golang/pkg/factory/cache"
	"github.com/kukkar/common-golang/pkg/factory/sql"
)

type AppConfig struct {
	MySql              *sql.MysqlConfig   `json:"Mysql"`
	Cache              *cache.CacheConfig `json:"Cache"`
	UseClientValidator bool               `json:"UseClientValidator"`
	SentryDSN          string             `json:"SentryDSN"`
}

func GetAppConfig() (*AppConfig, error) {
	c := config.GlobalAppConfig.ApplicationConfig
	appConfig, ok := c.(*AppConfig)
	if !ok {
		msg := fmt.Sprintf("Example APP Config Not correct %+v", c)
		return nil, errors.New(msg)
	}
	return appConfig, nil
}

func GetGlobalConfig() (*config.AppConfig, error) {
	return config.GlobalAppConfig, nil
}

func EnvUpdateMap() map[string]string {
	m := make(map[string]string)

	m["Mysql.User"] = "bussinesscard_MYSQL_USER"
	m["Mysql.Password"] = "bussinesscard_MYSQL_PASSWORD"
	m["Mysql.DbName"] = "bussinesscard_MYSQL_DBNAME"
	m["Mysql.MaxOpenConnections"] = "bussinesscard_MYSQL_MAXOPENCONNECTIONS"
	m["Mysql.MaxIdleConnections"] = "bussinesscard_MYSQL_MAXIDLECONNECTIONS"
	m["Mysql.DefaultTimeZone"] = "bussinesscard_MYSQL_DEFAULTTIMEZONE"
	m["Mysql.Host"] = "bussinesscard_MYSQL_HOST"
	m["Mysql.Port"] = "bussinesscard_MYSQL_PORT"
	m["Cache.Use"] = "bussinesscard_CACHE_USE"
	m["Cache.Redis.Addr"] = "bussinesscard_CACHE_REDIS_ADDRESS"
	m["Cache.Redis.PoolSize"] = "bussinesscard_CACHE_REDIS_POOLSIZE"

	m["SentryDSN"] = "bussinesscard_SENTRY_DSN"
	return m
}
