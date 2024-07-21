package server

import (
	"github.com/dilly3/houdini/api/server/response"
	"github.com/dilly3/houdini/pkg/github"
	"net/http"
)

func (h *Handler) ListCommitsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	repo := params.Get("repo")
	if repo == "" {
		response.RespondWithError(w, http.StatusBadRequest, "repo is required")
		return
	}
	owner := params.Get("owner")
	if owner == "" {
		response.RespondWithError(w, http.StatusBadRequest, "owner is required")
		return
	}
	since := params.Get("since")
	getCommits, err := github.DefaultGHClient.ListCommits(owner, repo, since)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to list commits")
		response.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondWithJson(w, "commits retrieved successfully", http.StatusOK, getCommits)
	return
}
