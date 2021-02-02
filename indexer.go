package indexer

import "io"

// Indexer ...
type Indexer interface {
	Update(index, id string, body io.Reader)
	// Get(id string) (document interface{})
}
