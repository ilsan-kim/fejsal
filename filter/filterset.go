package filter

type Filterable interface {
	filt() bool
}

// FSet implements the Filterable interface which allows it to be used in the FTree.
// FSet is a generic type that holds a value and a set of filters that can be applied to the value.
// It also has a condition field that determines whether the filters should be evaluated using an AND/OR logic.
//
// Example Usage:
/*
An FSet with dataGetter retrieving the value "banana", and filters checking if this value contains "ana" or not equal to "tomato" looks like:
  fset := FSet[string]{
    DataGetter: func() (string, bool) {return "banana", true},
    Filters: []Filter[string]{
      {operator: OperatorContain, valueType: ValueTypeString, value: "ana"},
      {operator: OperatorNotEqual, valueType: ValueTypeString, value: "tomato"},
    },
    Condition: ConditionOr,
  }
*/
type FSet[T Value] struct {
	DataGetter func() (T, bool)
	Filters    []Filter[T]
	Condition  Condition
}

func NewFilterSet[T Value](dataGetter func() (T, bool), filters []Filter[T], condition Condition) FSet[T] {
	return FSet[T]{DataGetter: dataGetter, Filters: filters, Condition: condition}
}

func (f FSet[T]) filt() bool {
	if f.DataGetter == nil {
		return false
	}

	data, ok := f.DataGetter()
	if !ok {
		return false
	}

	hasFiltered := false
	allFiltered := true

	for _, filter := range f.Filters {
		filtered := filter.filtData(data)
		hasFiltered = filtered
		if f.Condition == ConditionOr {
			if hasFiltered {
				return hasFiltered
			}
		}

		if !filtered {
			allFiltered = false
		}
	}

	return allFiltered
}
