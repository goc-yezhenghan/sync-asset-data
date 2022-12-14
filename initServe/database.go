package initServe

import (
	"fmt"
	"syncData/config"
	"syncData/global"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMySQL() {

	// 获取mysql配置
	var mysqlConfig config.Mysql
	if err := viper.UnmarshalKey("mysql", &mysqlConfig); err != nil {
		fmt.Println("failed to read mysql config", err.Error())
	}

	// 获取gorm连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(),
		DefaultStringSize:         120,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // 在控制台打印执行的sql语句
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("gorm connect failed")
	} else {
		global.DB = db
	}

	

}
