package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed web/swagger-ui
var SwaggerUIFS embed.FS

//go:embed api/openapi.yaml
var OpenAPISpec []byte

func setupSwaggerUI(r chi.Router) {
	r.Get("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		// w.Header().Set("Content-Type", "text/yaml")
		w.Write(OpenAPISpec)
	})

	swaggerUIRoot, err := fs.Sub(SwaggerUIFS, "web/swagger-ui")
	if err != nil {
		log.Fatalf("Error creating sub-filesystem for swagger-ui: %v", err)
	}

	r.Get("/swagger-ui/*", http.StripPrefix("/swagger-ui/", http.FileServer(http.FS(swaggerUIRoot))).ServeHTTP)
	
	r.Get("/swagger-ui", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger-ui/", http.StatusMovedPermanently)
	})

	// Redirect root to Swagger UI for convenience
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger-ui/", http.StatusMovedPermanently)
	})
}
