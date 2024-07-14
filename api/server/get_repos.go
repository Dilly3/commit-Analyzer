package server

import (
	"github.com/dilly3/houdini/api/server/response"
	"github.com/dilly3/houdini/pkg/github"
	"net/http"
)

func (h *Handler) GetReposHandler(w http.ResponseWriter, _ *http.Request) {
	repos, err := github.DefaultGHClient.GetRepos()
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to get repos")
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.RespondWithJson(w, "repos retrieved successfully", http.StatusOK, repos)
	return
}
