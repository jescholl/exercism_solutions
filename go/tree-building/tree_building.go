package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// sort input records to ease validation
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	if len(records) == 0 {
		return nil, nil
	}
	if err := ValidateRecords(records); err != nil {
		return nil, err
	}

	node := BuildNode(records)
	return node, nil
}

func ValidateRecords(records []Record) error {
	switch {
	case records[0].ID != 0:
		return errors.New("No root node")
	case records[0].Parent != 0:
		return errors.New("Root node must be its own parent")
	}
	for i := 0; i < len(records); i++ {
		if records[i].ID != i {
			return errors.New("Missing node")
		}
		if records[i].ID <= records[i].Parent && i != 0 {
			return errors.New("Node cannot have ID less than parent")
		}
	}
	return nil
}

func BuildNode(records []Record) *Node {
	// Build tree starting with first node in []records
	tree, records := &Node{ID: records[0].ID}, records[1:]

	for i := 0; i < len(records); i++ {
		if records[i].Parent == tree.ID {
			// Recursively call BuildNode to generate children
			tree.Children = append(tree.Children, BuildNode(records[i:]))
		}
	}
	return tree
}
