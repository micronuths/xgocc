package da

import (
	"xgocc/pkg/cedar"
)

// Dict contains the Trie and dict values
type Dict struct {
	Trie   *cedar.Cedar
	Values [][]string
}
