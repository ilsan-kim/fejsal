# Fejsal Filter Library

Fejsal is a Go-based filtering library designed to simplify data filtering operations using intuitive and powerful filter expressions built as binary trees.

This library allows users to define complex filtering logic for structured data such as CSV, JSONL, or simple text logs, enabling efficient data querying even in concurrent and distributed environments.

> Fejsal was humorously inspired by the Serbian footballer Fejsal Mulić, known for his less-than-stellar moments on the field.


## Overview

The Fejsal library provides:

- **Generic Filters:** Define filters for different data types (string, number, datetime) with operators like `CONTAIN`, `EQUAL`, `LESS_THAN`, etc.
- **Filter Sets (`FSet`):** Combine multiple filters with logical conditions (`AND`/`OR`).
- **Filter Trees (`FTree`):** Create advanced nested logical expressions by combining filter sets in a binary tree structure.

## Key Components

### Filter

The smallest unit of filtering logic, representing conditions for specific data types:

- String: `CONTAIN`, `EQUAL`, `NOT_EQUAL`
- Number and Datetime: `EQUAL`, `NOT_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`

### FSet

Combines multiple filters and evaluates them using logical conditions (`AND`, `OR`).

Example:

```go
fset := FSet[string]{
	DataGetter: func() (string, bool) { return "banana", true },
	Filters: []Filter[string]{
		{operator: OperatorContain, valueType: ValueTypeString, value: "ana"},
		{operator: OperatorNotEqual, valueType: ValueTypeString, value: "tomato"},
	},
	Condition: ConditionOr,
}
```

### FTree

Constructs complex logical conditions as binary trees:

```go
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
```
## Testing

Run tests to validate functionality:

```bash
go test ./...
```

## License

MIT License.

