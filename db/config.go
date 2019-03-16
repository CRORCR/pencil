package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

/**
 * @desc    初始化mysql redis
 * @author Ipencil
 * @create 2019/2/27
 */
func InitConfig(conKey string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("dbs")
	viper.AddConfigPath("../pencil/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	subv := viper.Sub(conKey)
	db_user := subv.Get("db_user")
	db_home := subv.Get("db_home")
	db_name := subv.Get("db_name")
	db_pass := subv.Get("db_pass")
	max_conn := subv.Get("max_conn")
	max_open := subv.Get("max_open")
	initDB(db_user, db_home, db_name, db_pass, max_conn, max_open)
}

var g_dbHand []*xorm.Engine

func GetDBHand(nIndex int) *xorm.Engine {
	return g_dbHand[nIndex%len(g_dbHand)]
}

func initDB(db_user, db_home, db_name, db_pass, max_conn, max_open interface{}) {
	//root:root@tcp(127.0.0.1:3306)/test
	strConn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", db_user, db_pass, db_home, db_name)
	fmt.Println(strConn)
	//("mysql", "root:feidianDB*@#4@tcp(172.16.250.198:3306)/thp")
	dbHand, err := xorm.NewEngine("mysql", strConn)
	if err != nil {
		panic(fmt.Errorf("LIKE DB FAILED:%v\n", err))
		return
	}
	if err := dbHand.Ping(); err != nil {
		panic(fmt.Errorf("DB PING FAILED:%v\n", err))
		return
	}
	dbHand.SetMapper(core.GonicMapper{})
	dbHand.ShowSQL(true)
	dbHand.SetMaxIdleConns(max_conn.(int))
	dbHand.SetMaxOpenConns(max_open.(int))
	g_dbHand = append(g_dbHand, dbHand)
}
