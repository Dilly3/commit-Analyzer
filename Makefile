.PHONY: mocks test up prune

mocks:
	mockgen -destination=api/server/mocks/repo_mock.go -package=mocks github.com/dilly3/houdini/internal/storage IRepoStore
	mockgen -destination=api/server/mocks/commit_mock.go -package=mocks github.com/dilly3/houdini/internal/storage ICommitStore
	mockgen -destination=api/server/mocks/store_mock.go -package=mocks github.com/dilly3/houdini/internal/storage IStore

prune:
	docker image prune -a -f
up:
	docker compose up --build
test: mocks
	go test -count=1 ./api/server