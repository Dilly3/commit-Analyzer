package server

import (
	"github.com/dilly3/houdini/api/server/response"
	errs "github.com/dilly3/houdini/internal/error"
	"github.com/dilly3/houdini/internal/storage"
	"github.com/dilly3/houdini/pkg/github"
	"net/http"
)

func (h *Handler) GetRepoHandler(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	repo := params.Get("repo")
	if repo == "" {
		http.Error(w, "repo is required", http.StatusBadRequest)
		return
	}
	owner := params.Get("owner")
	if owner == "" {
		http.Error(w, "owner is required", http.StatusBadRequest)
		return
	}
	getRepo, err := github.DefaultGHClient.GetRepo(owner, repo)
	if err != nil {
		h.Logger.Error().Err(err).Msg("failed to get repo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = storage.GetDefaultStore().SaveRepo(r.Context(), getRepo)
	if err != nil {
		errq := errs.NewAppError("save repo:", err)
		response.RespondWithError(w, http.StatusInternalServerError, errq.Error())
		return
	}
	response.RespondWithJson(w, "repo retrieved successfully", http.StatusOK, getRepo)
	return
}
