package storage

var defaultStore IStore

func GetDefaultStore() IStore {
	return defaultStore
}

type IStore interface {
	ICommitStore
	IAuthorStore
	IRepoStore
}

type Store struct {
	ICommitStore
	IAuthorStore
	IRepoStore
}

func NewStore(storage *Storage) Store {
	ics := NewCommitStore(storage)
	ias := NewAuthorStore(storage)
	irs := NewRepoStore(storage)
	store := Store{
		ICommitStore: ics,
		IAuthorStore: ias,
		IRepoStore:   irs,
	}
	defaultStore = store
	return store
}
