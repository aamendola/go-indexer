package indexer

// Indexer ...
type Indexer interface {
	Update(index, id, content string)
	// Get(id string) (document interface{})
}
