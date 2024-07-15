package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func NewChiRouter(h *Handler, limiterDuration time.Duration) http.Handler {
	router := chi.NewRouter()
	limiter := NewRateLimiter(limiterDuration)
	//router.HandlerFunc(http.MethodGet, "/v1/repo", h.GetRepoHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/commits", h.ListCommitsHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/repos", h.GetReposHandler)
	router.Get("/v1/set/repo/credential", h.SetRepoCredentialHandler)
	router.Get("/v1/repo/{name}", h.GetRepoByName)
	router.Get("/v1/commits/{name}/{limit}", h.GetCommitsByRepoName)
	router.Get("/v1/repos/{language}/{limit}", h.GetReposByLanguage)
	router.Get("/v1/repos-stars/{limit}", h.GetRepoByStarsCount)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})
	return limiter.IPRateLimit(c.Handler(router))
}
