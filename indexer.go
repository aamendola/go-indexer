package indexer

// Indexer ...
type Indexer interface {
	Update2(index, id string, message interface{}) error
}
