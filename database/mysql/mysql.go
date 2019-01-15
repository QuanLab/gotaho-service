package mysql

import (
	"database/sql"
	"fmt"
	"github.com/QuanLab/go-service/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB *sql.DB
)

func init() {
	Connect()
}

func DataSourceName() string {
	return config.Get().MySQL.Username +
		":" +
		config.Get().MySQL.Password +
		"@tcp(" +
		config.Get().MySQL.Host +
		":" +
		fmt.Sprintf("%d", config.Get().MySQL.Port) +
		")/" +
		config.Get().MySQL.Database + config.Get().MySQL.Parameter
}

//open connection to MySQL database, it 's self contains pool init
func Connect() {
	var err error
	DB, err = sql.Open("mysql", DataSourceName())
	if err != nil {
		log.Printf("Cannot connect to MySQL server %s", err)
	}
}
