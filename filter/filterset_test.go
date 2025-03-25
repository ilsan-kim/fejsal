package filter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFSetFilt(t *testing.T) {
	testNumberInt := []struct {
		name       string
		filters    []Filter[int]
		condition  Condition
		dataGetter func() (int, bool)
		want       bool
	}{
		{
			name: "NumberInt Filtering with AND Condition (15 is GreaterThan 10 AND 15 is LessThanOrEqual 20",
			filters: []Filter[int]{
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 10),
				mustNewFilter(OperatorLessThanOrEqual, ValueTypeNumber, 20),
			},
			condition:  ConditionAnd,
			dataGetter: func() (int, bool) { return 15, true },
			want:       true,
		},
		{
			name: "NumberInt Filtering with AND Condition (30 is GreaterThanOrEqual 30 AND 30 is LessThan 40",
			filters: []Filter[int]{
				mustNewFilter(OperatorGreaterThanOrEqual, ValueTypeNumber, 30),
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 40),
			},
			condition:  ConditionAnd,
			dataGetter: func() (int, bool) { return 30, true },
			want:       true,
		},
		{
			name: "NumberInt Filtering With OR Condition with multiple filters if only one filter is satisfied",
			filters: []Filter[int]{
				mustNewFilter(OperatorEqual, ValueTypeNumber, 3158),
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 10000),
				mustNewFilter(OperatorEqual, ValueTypeNumber, 0),
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 3158),
			},
			condition:  ConditionOr,
			dataGetter: func() (int, bool) { return 3158, true },
			want:       true,
		},
		{
			name: "NumberInt Filtering with OR Condition if only one filter is satisfied (15 is LessThan 20 OR 15 is GreaterThan 25)",
			filters: []Filter[int]{
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 20),
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 25),
			},
			condition:  ConditionOr,
			dataGetter: func() (int, bool) { return 15, true },
			want:       true,
		},
		{
			name: "NumberInt Filtering failed with AND Condition (20 is LessThan 20 AND 20 is GreaterThan 30)",
			filters: []Filter[int]{
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 20),
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 30),
			},
			condition:  ConditionAnd,
			dataGetter: func() (int, bool) { return 20, true },
			want:       false,
		},
	}

	testNumberFloat64 := []struct {
		name       string
		filters    []Filter[float64]
		condition  Condition
		dataGetter func() (float64, bool)
		want       bool
	}{
		{
			name: "NumberFloat64 Filtering with AND Condition (3.14 is GreaterThan 3.10 AND 3.14 is LessThan 4.77)",
			filters: []Filter[float64]{
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 3.10),
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 4.77),
			},
			condition:  ConditionAnd,
			dataGetter: func() (float64, bool) { return 3.14, true },
			want:       true,
		},
		{
			name: "NumberFloat64 Filtering with AND Condition (3.14 is LessThan 1.25 ANd 3.14 is GreaterThan 1.25)",
			filters: []Filter[float64]{
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 1.25),
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 1.25),
			},
			condition:  ConditionAnd,
			dataGetter: func() (float64, bool) { return 3.14, true },
			want:       false,
		},
		{
			name: "NumberFloat64 Filtering with AND Condition but approximately equal numbers (3.1400002 is approximately equal with 3.1400005 AND 3.1400002 is approximately equal with 3.1400001)",
			filters: []Filter[float64]{
				mustNewFilter(OperatorEqual, ValueTypeNumber, 3.1400005),
				mustNewFilter(OperatorEqual, ValueTypeNumber, 3.1400001),
			},
			condition:  ConditionAnd,
			dataGetter: func() (float64, bool) { return 3.1400002, true },
			want:       true,
		},
		{
			name: "NumberFloat64 Filtering with OR Condition if only one filter is satisfied",
			filters: []Filter[float64]{
				mustNewFilter(OperatorLessThan, ValueTypeNumber, 3.6133),
				mustNewFilter(OperatorLessThanOrEqual, ValueTypeNumber, 6.1241),
				mustNewFilter(OperatorGreaterThan, ValueTypeNumber, 1031.2125),
			},
			condition:  ConditionOr,
			dataGetter: func() (float64, bool) { return 6.1241, true },
			want:       true,
		},
	}

	testNumberDateTime := []struct {
		name       string
		filters    []Filter[time.Time]
		condition  Condition
		dataGetter func() (time.Time, bool)
		want       bool
	}{
		{
			name: "DateTime Filtering if pass one filter (now - 1 hour is earlier(less in unixtime) than now)",
			filters: []Filter[time.Time]{
				mustNewFilter(OperatorLessThanOrEqual, ValueTypeDatetime, time.Now()),
			},
			condition:  ConditionOr,
			dataGetter: func() (time.Time, bool) { return time.Now().Add(-time.Hour), true },
			want:       true,
		},
		// TODO: and more....
	}

	testString := []struct {
		name       string
		filters    []Filter[string]
		condition  Condition
		dataGetter func() (string, bool)
		want       bool
	}{
		{
			name: "String Filtering if pass one filter (banana contains anana)",
			filters: []Filter[string]{
				mustNewFilter(OperatorContain, ValueTypeString, "anana"),
			},
			condition:  ConditionAnd,
			dataGetter: func() (string, bool) { return "banana", true },
			want:       true,
		},
	}

	for _, tt := range testNumberInt {
		t.Run(tt.name, func(t *testing.T) {
			fset := NewFilterSet[int](tt.dataGetter, tt.filters, tt.condition)
			assert.Equal(t, tt.want, fset.filt())
		})
	}

	for _, tt := range testNumberFloat64 {
		t.Run(tt.name, func(t *testing.T) {
			fset := NewFilterSet[float64](tt.dataGetter, tt.filters, tt.condition)
			assert.Equal(t, tt.want, fset.filt())

		})
	}

	for _, tt := range testNumberDateTime {
		t.Run(tt.name, func(t *testing.T) {
			fset := NewFilterSet[time.Time](tt.dataGetter, tt.filters, tt.condition)
			assert.Equal(t, tt.want, fset.filt())
		})
	}

	for _, tt := range testString {
		t.Run(tt.name, func(t *testing.T) {
			fset := NewFilterSet[string](tt.dataGetter, tt.filters, tt.condition)
			assert.Equal(t, tt.want, fset.filt())
		})
	}
}

// mustNewFilter is a helper function to create a Filter and panic if an error occurs.
func mustNewFilter[T Value](operator Operator, valueType ValueType, value T) Filter[T] {
	f, err := NewFilter(operator, valueType, value)
	if err != nil {
		panic(err)
	}
	return f
}
