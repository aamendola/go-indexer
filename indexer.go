package indexer

// Indexer ...
type Indexer interface {
	Update(index, id string, message interface{}) error
}
