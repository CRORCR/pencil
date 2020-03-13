package db

import (
	"errors"
	"fmt"
	"pencil/global"
	p_logger "pencil/utils/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DB_User = `user`
)

type mysql struct {
	DataBaseUser *gorm.DB
	DataBaseVIP  *gorm.DB
}

var (
	MysqlService *mysql
)

func init() {
	MysqlService = new(mysql)
	MysqlService.loadDataBase()
}
func (o *mysql) GetTransaction(db string) (*gorm.DB, error) {
	switch db {
	case DB_User:
		return o.DataBaseUser.Begin(), nil

	}
	return nil, errors.New("db not exist")
}

func (o *mysql) loadDataBase() {
	o.DataBaseVIP = CreateDB(global.GlobalConfig.Conf.DB.VIP.Username,
		global.GlobalConfig.Conf.DB.VIP.Password,
		global.GlobalConfig.Conf.DB.VIP.Host,
		global.GlobalConfig.Conf.DB.VIP.Port,
		global.GlobalConfig.Conf.DB.VIP.Database,
		global.GlobalConfig.Conf.DB.VIP.LogMode,
		global.GlobalConfig.Conf.DB.VIP.MaxIdleConns,
		global.GlobalConfig.Conf.DB.VIP.MaxOpenConns,
		global.GlobalConfig.Conf.DB.VIP.IsAutoMigrate)

	o.DataBaseUser = CreateDB(global.GlobalConfig.Conf.DB.User.Username,
		global.GlobalConfig.Conf.DB.User.Password,
		global.GlobalConfig.Conf.DB.User.Host,
		global.GlobalConfig.Conf.DB.User.Port,
		global.GlobalConfig.Conf.DB.User.Database,
		global.GlobalConfig.Conf.DB.User.LogMode,
		global.GlobalConfig.Conf.DB.User.MaxIdleConns,
		global.GlobalConfig.Conf.DB.User.MaxOpenConns,
		global.GlobalConfig.Conf.DB.User.IsAutoMigrate)

}

func CreateDB(userName string, passWord string, host string, port uint32, dbName string, log_mode bool, max_idle_conns int, max_open_conns int, is_auto_migrate bool) *gorm.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		userName,
		passWord,
		host,
		port,
		dbName,
	)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic("mysql conn fail," + host + ":" + dbName)
	}
	p_logger.Logger.Info(fmt.Sprintf("Connected to MySQL db : %s ,host : %s ,port: %d", dbName, host, port))
	db.LogMode(log_mode)

	db.DB().SetMaxIdleConns(max_idle_conns)
	db.DB().SetMaxOpenConns(max_open_conns)
	//db.DB().SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)
	if is_auto_migrate {
		db.AutoMigrate()
	}

	return db
}
