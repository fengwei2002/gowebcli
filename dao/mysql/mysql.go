package mysql

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql" // 导入驱动
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() (err error) {
	//dsn := "user:password:@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	// 数据库的连接信息由配置文件进行填充
	dsn := fmt.Sprintf("%s:%s:@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := sqlx.Connect("mysql", dsn)
	// 也可以使用 MustConnect 连接不成功就 panic
	if err != nil {
		//fmt.Printf("connect DB failed, err: %v\n", err)
		// 使用 zap 日志库记录 error
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
