package main

import (
	"LearnGoCRUD/internal/database"
	"LearnGoCRUD/internal/models"
	"LearnGoCRUD/internal/service"
	"LearnGoCRUD/internal/transport"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {

	db, err := sqlx.Open("postgres", models.ConnStr) //установка соединения с бд
	if err != nil {
		log.Fatalf("db error sqlx.Open:%s\n", err)
	}
	defer func(db *sqlx.DB) { //закрытие соединения с бд
		err := db.Close()
		if err != nil {
		}
	}(db)

	_, err = db.Exec(models.CreateTableQuery) //создает таблицу, если ее не было
	if err != nil {
		log.Fatalf("db error createTableQuery:%s\n", err)
	}

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
