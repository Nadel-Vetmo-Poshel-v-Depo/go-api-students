package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/Nadel-Vetmo-Poshel-v-Depo/go-api-students/internal/db"
	"github.com/Nadel-Vetmo-Poshel-v-Depo/go-api-students/internal/handlers"
)

func main() {
	// Загружаем .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	conn := db.InitDB()
	defer conn.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/students", handlers.GetStudentsHandler(conn))
	r.Post("/students", handlers.CreateStudentHandler(conn))
	r.Put("/students/{id}", handlers.UpdateStudentHandler(conn))
	r.Delete("/students/{id}", handlers.DeleteStudentHandler(conn))

	log.Println("Сервер запущен на http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
