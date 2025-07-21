package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Nadel-Vetmo-Poshel-v-Depo/go-api-students/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func GetStudentsHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.Background(), "SELECT id, name, age, enrolled FROM students")
		if err != nil {
			http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var students []models.Student

		for rows.Next() {
			var s models.Student
			err := rows.Scan(&s.ID, &s.FirstName, &s.Age, &s.Enrolled)
			if err != nil {
				http.Error(w, "Ошибка чтения строки", http.StatusInternalServerError)
				return
			}
			students = append(students, s)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}

func CreateStudentHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s models.Student

		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, "Невалидный JSON", http.StatusBadRequest)
			return
		}

		sql := `INSERT INTO students (name, age, enrolled) VALUES ($1, $2, $3) RETURNING id`
		err = db.QueryRow(context.Background(), sql, s.FirstName, s.Age, s.Enrolled).Scan(&s.ID)
		if err != nil {
			http.Error(w, "Ошибка вставки", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	}
}

func UpdateStudentHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Некорректный ID", http.StatusBadRequest)
			return
		}

		var s models.Student
		err = json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, "Невалидный JSON", http.StatusBadRequest)
			return
		}

		sql := `UPDATE students SET name=$1, age=$2, enrolled=$3 WHERE id=$4`
		_, err = db.Exec(context.Background(), sql, s.FirstName, s.Age, s.Enrolled, id)
		if err != nil {
			http.Error(w, "Ошибка обновления", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteStudentHandler(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			http.Error(w, "Некорректный ID", http.StatusBadRequest)
			return
		}

		sql := `DELETE FROM students WHERE id=$1`
		_, err = db.Exec(context.Background(), sql, id)
		if err != nil {
			http.Error(w, "Ошибка удаления", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
