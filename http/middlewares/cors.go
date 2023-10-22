package middlewares

import (
	"net/http"
	
	"github.com/rs/cors"
)

func CORS(next http.Handler) http.Handler {
	allowedHeaders := []string{"Authorization", "Content-Type"}
	return cors.New(cors.Options{Debug: true, AllowedHeaders: allowedHeaders}).Handler(next)
}