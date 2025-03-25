package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockFilterable struct {
	result bool
}

func (m mockFilterable) filt() bool {
	return m.result
}

func TestFTree_Evaluate(t *testing.T) {
	tests := []struct {
		name     string
		ftree    *FTree
		expected bool
	}{
		{
			name: "Filtered single node",
			ftree: &FTree{
				FilterSet: mockFilterable{result: true},
			},
			expected: true,
		},
		{
			name: "Unfiltered single node",
			ftree: &FTree{
				FilterSet: mockFilterable{result: false},
			},
			expected: false,
		},
		{
			name: "Filtered And Unfiltered -> Unfiltered",
			ftree: &FTree{
				Left:      &FTree{FilterSet: mockFilterable{result: true}},
				Right:     &FTree{FilterSet: mockFilterable{result: false}},
				Condition: ConditionAnd,
			},
			expected: false,
		},
		{
			name: "Filtered Or Unfiltered -> Filtered",
			ftree: &FTree{
				Left:      &FTree{FilterSet: mockFilterable{result: true}},
				Right:     &FTree{FilterSet: mockFilterable{result: false}},
				Condition: ConditionOr,
			},
			expected: true,
		},
		{
			name: "(Filtered OR Unfiltered) And Filtered -> Filtered",
			ftree: &FTree{
				Left: &FTree{
					Left:      &FTree{FilterSet: mockFilterable{result: true}},
					Right:     &FTree{FilterSet: mockFilterable{result: false}},
					Condition: ConditionOr,
				},
				Right:     &FTree{FilterSet: mockFilterable{result: true}},
				Condition: ConditionAnd,
			},
			expected: true,
		},
		{
			name: "(Unfiltered OR UnFiltered) And Filtered) -> Unfiltered",
			ftree: &FTree{
				Left: &FTree{
					Left:      &FTree{FilterSet: mockFilterable{result: false}},
					Right:     &FTree{FilterSet: mockFilterable{result: false}},
					Condition: ConditionOr,
				},
				Right:     &FTree{FilterSet: mockFilterable{result: true}},
				Condition: ConditionAnd,
			},
			expected: false,
		},
		{
			name: "(Unfiltered OR Filtered) And (Filtered Or Unfiltered) -> Filtered",
			ftree: &FTree{
				Left: &FTree{
					Left:      &FTree{FilterSet: mockFilterable{result: false}},
					Right:     &FTree{FilterSet: mockFilterable{result: true}},
					Condition: ConditionOr,
				},
				Right: &FTree{
					Left:      &FTree{FilterSet: mockFilterable{result: true}},
					Right:     &FTree{FilterSet: mockFilterable{result: false}},
					Condition: ConditionOr,
				},
				Condition: ConditionAnd,
			},
			expected: true,
		},
		{
			name: "(Filtered Or (Unfiltered And Filtered)) And (Filtered And Filtered)",
			ftree: &FTree{
				Left: &FTree{
					Left: &FTree{FilterSet: mockFilterable{result: true}},
					Right: &FTree{
						Left:      &FTree{FilterSet: mockFilterable{result: false}},
						Right:     &FTree{FilterSet: mockFilterable{result: true}},
						Condition: ConditionAnd,
					},
					Condition: ConditionOr,
				},
				Right: &FTree{
					Left:      &FTree{FilterSet: mockFilterable{result: true}},
					Right:     &FTree{FilterSet: mockFilterable{result: false}},
					Condition: ConditionOr,
				},
				Condition: ConditionAnd,
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := tc.ftree.Evaluate()
			assert.Equal(t, tc.expected, res)
		})
	}

}
