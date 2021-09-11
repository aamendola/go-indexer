package indexer

// Indexer ...
type Indexer interface {
	Update(index, id, content string) error
	Update2(index, id string, message interface{}) error
}
