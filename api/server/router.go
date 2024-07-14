package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func NewMuxRouter(h *Handler, limiterDuration time.Duration) http.Handler {
	router := httprouter.New()
	limiter := NewRateLimiter(limiterDuration)
	router.HandlerFunc(http.MethodGet, "/v1/repo", h.GetRepoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/commits", h.ListCommitsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/repos", h.GetReposHandler)
	router.HandlerFunc(http.MethodGet, "/v1/set/repo/credential", h.SetRepoCredentialHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allows all origins
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})
	return limiter.IPRateLimit(c.Handler(router))
}
