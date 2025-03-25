package filter

// FTree represents a binary tree structure used for filtering log tokens.
// Each node can either be a leaf node containing a filter set (FilterSet) or
// an internal node with a logical condition (AND/OR) applied between two child nodes (Left and Right).
//
// Example Usage:
/*
Thus the FTree of

Thus the FTree ((FSet#1 OR FSet#2) AND FSet#3) looks like this code

  tree := &FTree{
    Left: &FTree{
        Left:      &FTree{FilterSet: FSet#1},
        Right:     &FTree{FilterSet: FSet#2},
        Condition: ConditionOr,
    },
    Right:     &FTree{FilterSet: FSet#3},
    Condition: ConditionAnd,
  }

  result := tree.Evaluate()
  fmt.Println("Filter result:", result)  true or false

*/
type FTree struct {
	Left, Right *FTree
	Condition   Condition
	FilterSet   Filterable
}

// Evaluate executes the filtering logic on the tree.
// It recursively evaluates the left and right subtrees if it's an internal node,
// or directly applies the filter set if it's a leaf node.
func (ft *FTree) Evaluate() bool {
	if ft.FilterSet != nil {
		return ft.FilterSet.filt()
	}

	if ft.Condition == ConditionAnd {
		return ft.Left.Evaluate() && ft.Right.Evaluate()
	}
	if ft.Condition == ConditionOr {
		return ft.Left.Evaluate() || ft.Right.Evaluate()
	}

	return false
}
