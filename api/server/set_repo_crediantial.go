package server

import (
	"github.com/dilly3/houdini/api/server/response"
	"github.com/dilly3/houdini/internal/model"
	"net/http"
)

func (h *Handler) SetRepoCredentialHandler(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	repo := r.URL.Query().Get("repo")
	if owner == "" {
		response.RespondWithError(w, http.StatusBadRequest, "owner is required")
		return
	}
	if repo == "" {
		response.RespondWithError(w, http.StatusBadRequest, "repo is required")
		return
	}
	model.SetOwnerName(owner)
	model.SetRepoName(repo)
	response.RespondWithJson(w, "repo credentials set successfully", http.StatusOK, nil)
}
