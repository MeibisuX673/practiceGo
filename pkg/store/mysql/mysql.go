package mysql

import (

	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

type Store struct{
	
	Database *sql.DB
}

type ConfigMysql struct{

	DbName string
	User string
	Password string
	Port string
	ServerName string

}

func (configMysql *ConfigMysql) toDsn() string{

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", configMysql.User, configMysql.Password, configMysql.ServerName, configMysql.Port, configMysql.DbName)

	return dsn
}

func New(configMysql ConfigMysql) (*Store, error){

	db, err := sql.Open("mysql", configMysql.toDsn())

	if err != nil{
		return nil, err
	}

	return &Store{Database: db}, nil

}



