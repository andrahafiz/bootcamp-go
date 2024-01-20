package transactions

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Register(router *chi.Mux, db *sql.DB) {

	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Route("/api/transactions", func(r chi.Router) {
		r.Post("/", handler.Process)
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.DetailById)
	})
}
