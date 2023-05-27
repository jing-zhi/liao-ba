package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"liaoBa/settings"
)

// 定义一个全局对象db
var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	// DSN:Data Source Name
	//dsn := fmt.Sprintf("#{cfg.User}:#{cfg.Password}@tcp(#{cfg.Host}:#{cfg.Port})/#{cfg.DbName}?charset=utf8mb4&parseTime=True")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	return
}
func Close() {
	_ = db.Close()
}
