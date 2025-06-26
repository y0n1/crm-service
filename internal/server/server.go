package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/y0n1/crm-service/internal/server/slices/customers/create"
	"github.com/y0n1/crm-service/internal/server/slices/customers/delete"
	"github.com/y0n1/crm-service/internal/server/slices/customers/get"
	"github.com/y0n1/crm-service/internal/server/slices/customers/list"
	"github.com/y0n1/crm-service/internal/server/slices/customers/update"
	"github.com/y0n1/crm-service/internal/store"
)

type Server struct {
	ctx context.Context
}

func New() *Server {
	return &Server{
		ctx: context.Background(),
	}
}

func (s Server) Run() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db := store.NewMemoryStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	setupSwaggerUI(r)
	
	r.Route("/v1", func(r chi.Router) {
		r.Get(list.UrlPattern, list.MakeHandler(s.ctx, db, logger))
		r.Get(get.UrlPattern, get.MakeHandler(s.ctx, db, logger))
		r.Post(create.UrlPattern, create.MakeHandler(s.ctx, db, logger))
		r.Patch(update.UrlPattern, update.MakeHandler(s.ctx, db, logger))
		r.Delete(delete.UrlPattern, delete.MakeHandler(s.ctx, db, logger))
	})

	logger.Info("Server starting on port 8888")
	if err := http.ListenAndServe(":8888", r); err != nil {
		logger.Error(err.Error())
	}
}
