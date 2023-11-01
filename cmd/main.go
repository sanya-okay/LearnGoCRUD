package main

import (
	"LearnGoPersonGinPsql/internal/database"
	"LearnGoPersonGinPsql/internal/service"
	"LearnGoPersonGinPsql/internal/transport"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {

	connStr := "user=db-user password=db-pass dbname=db-name sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("db error:%s\n", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	//initializing dependencies (инициализация зависимостей)
	pDB := database.NewDB(db)
	pService := service.NewService(pDB)
	handler := transport.NewHandler(pService)

	//инициализация сервера, стандартная библа http
	server := &http.Server{
		Addr:         ":1234",
		Handler:      handler.InitRouter(),
		ReadTimeout:  100 * time.Second,
		WriteTimeout: 100 * time.Second,
	}
	log.Println("Server is starting, но это не точно'")
	//запуск сервера
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln("Server cannot work")
	}

}
