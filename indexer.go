package indexer

// Indexer ...
type Indexer interface {
	Update(index, id, content string) error
	// Get(id string) (document interface{})
}
