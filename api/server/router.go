package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func NewChiRouter(h *Handler, limiterDuration time.Duration) *chi.Mux {
	router := chi.NewRouter()
	limiter := NewRateLimiter(limiterDuration)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	router.Use(c.Handler)
	router.Use(h.loggingMiddleware)
	router.Patch("/v1/settings", h.UpdateSettingsHandler)
	limitRoutes := router.With(limiter.IPRateLimit)

	limitRoutes.Get("/v1/repo", h.GetRepoHandler)
	limitRoutes.Get("/v1/commits", h.ListCommitsHandler)
	limitRoutes.Get("/v1/repos", h.GetReposHandler)

	limitRoutes.Get("/v1/repo/{name}", h.GetRepoByName)
	limitRoutes.Get("/v1/commits/{name}/{limit}", h.GetCommitsByRepoName)
	limitRoutes.Get("/v1/repos/{language}/{limit}", h.GetReposByLanguage)
	limitRoutes.Get("/v1/repos-stars/{limit}", h.GetRepoByStarsCount)

	return router

}
