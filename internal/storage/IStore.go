package storage

var defaultStore IStore

func GetDefaultStore() IStore {
	return defaultStore
}
func SetDefaultStore(store IStore) {
	defaultStore = store
}

type IStore interface {
	ICommitStore
	IRepoStore
}

type Store struct {
	ICommitStore
	IRepoStore
}

func NewStore(storage *Storage) Store {
	ics := NewCommitStore(storage)
	irs := NewRepoStore(storage)
	store := Store{
		ICommitStore: ics,
		IRepoStore:   irs,
	}
	defaultStore = store
	return store
}
