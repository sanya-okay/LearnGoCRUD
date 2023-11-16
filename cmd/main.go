package main

import (
	"LearnGoCRUD/internal/database"
	"LearnGoCRUD/internal/service"
	"LearnGoCRUD/internal/transport"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {

	createTableQuery := `CREATE TABLE IF NOT EXISTS Person(
                                     id SERIAL NOT NULL PRIMARY KEY,
                                     name TEXT NOT NULL,
                                     age INTEGER NOT NULL,
                                     license BOOL NOT NULL
);`
	connStr := "user=db-user password=db-pass dbname=db-name sslmode=disable"

	db, err := sqlx.Open("postgres", connStr) //установка соединения с бд
	if err != nil {
		log.Fatalf("db error sqlx.Open:%s\n", err)
	}
	defer func(db *sqlx.DB) { //закрытие соединения с бд
		err := db.Close()
		if err != nil {
		}
	}(db)

	_, err = db.Exec(createTableQuery) //создает таблицу, если ее не было
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
