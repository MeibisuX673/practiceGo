package main

import (
	"fmt"
	"github.com/MeibisuX673/practiceGo/pkg/store/mysql"
)

func main(){

	store, err := mysql.New(mysql.ConfigMysql{
		DbName: "base",
		ServerName: "127.0.0.1",
		Password: "test",
		User: "root",
		Port: "33061",
	})

	defer store.Database.Close()

	if err != nil{
		panic(err)
	}
	
	fmt.Println(store.Database.Ping())

}