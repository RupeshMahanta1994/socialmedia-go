package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rupeshmahanta/socialmedia-go/cmd/api"
	"github.com/rupeshmahanta/socialmedia-go/configs"
	"github.com/rupeshmahanta/socialmedia-go/db"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err := db.NewMySQLStorage(cfg)
	// db, err := sql.Open("mysql", "root:9437!Nrupendra@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("Error in connecting to database", err)
	}
	initStorage(db)
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal("Problem in running the server", err)
	}

}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Not able to ping DB", err)
	}
	log.Println("DB: Successfully connected!")
}
