package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func create() {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS goblog")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE goblog")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Employee` (`id` int(6) unsigned NOT NULL AUTO_INCREMENT,`name` varchar(30) NOT NULL,`city` varchar(30) NOT NULL,PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;")
	if err != nil {
		panic(err)
	}
}

func main() {
	create()
}
