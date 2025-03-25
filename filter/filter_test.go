package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFilter_Validate(t *testing.T) {
	testsString := []struct {
		name      string
		filter    Filter[string]
		wantValid bool
		wantErr   bool
	}{
		// Test Value & ValueType
		{
			name:      "Invalid Number Type for String",
			filter:    Filter[string]{operator: OperatorEqual, valueType: ValueTypeNumber, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid DateTime Type for String",
			filter:    Filter[string]{operator: OperatorEqual, valueType: ValueTypeDatetime, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Valid String Type for String",
			filter:    Filter[string]{operator: OperatorEqual, valueType: ValueTypeString, value: "test"},
			wantValid: true,
			wantErr:   false,
		},
		// Test Operator
		{
			name:      "Invalid LessThan Operator for String",
			filter:    Filter[string]{operator: OperatorLessThan, valueType: ValueTypeString, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid LessThanOrEqual Operator for String",
			filter:    Filter[string]{operator: OperatorLessThanOrEqual, valueType: ValueTypeString, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid GreaterThan Operator for String",
			filter:    Filter[string]{operator: OperatorGreaterThan, valueType: ValueTypeString, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid GreaterThanOrEqual Operator for String",
			filter:    Filter[string]{operator: OperatorGreaterThanOrEqual, valueType: ValueTypeString, value: "test"},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Valid Contain Operator for String",
			filter:    Filter[string]{operator: OperatorContain, valueType: ValueTypeString, value: "test"},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid Equal Operator for String",
			filter:    Filter[string]{operator: OperatorEqual, valueType: ValueTypeString, value: "test"},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid NotEqual Operator for String",
			filter:    Filter[string]{operator: OperatorNotEqual, valueType: ValueTypeString, value: "test"},
			wantValid: true,
			wantErr:   false,
		},
	}

	testsInt := []struct {
		name      string
		filter    Filter[int]
		wantValid bool
		wantErr   bool
	}{
		// Tests Value & ValueType
		{
			name:      "Invalid DateTime Type for Number",
			filter:    Filter[int]{operator: OperatorEqual, valueType: ValueTypeDatetime, value: 123},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid String Type for Number",
			filter:    Filter[int]{operator: OperatorEqual, valueType: ValueTypeString, value: 123},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Valid Number Type for Number",
			filter:    Filter[int]{operator: OperatorEqual, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		// Test Operator
		{
			name:      "Valid Equal Operator for Integer",
			filter:    Filter[int]{operator: OperatorEqual, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid NotEqual Operator for Integer",
			filter:    Filter[int]{operator: OperatorNotEqual, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid LessThan Operator for Integer",
			filter:    Filter[int]{operator: OperatorLessThan, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid LessThanOrEqual Operator for Integer",
			filter:    Filter[int]{operator: OperatorLessThanOrEqual, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid GreaterThan Operator for Integer",
			filter:    Filter[int]{operator: OperatorGreaterThan, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Valid GreaterThanOrEqual Operator for Integer",
			filter:    Filter[int]{operator: OperatorGreaterThanOrEqual, valueType: ValueTypeNumber, value: 123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Invalid Contain Operator for Integer",
			filter:    Filter[int]{operator: OperatorContain, valueType: ValueTypeNumber, value: 123},
			wantValid: false,
			wantErr:   true,
		},
	}

	testsFloat := []struct {
		name      string
		filter    Filter[float64]
		wantValid bool
		wantErr   bool
	}{
		// Test Value & ValueType
		{
			name:      "Valid Number Type for Number(float)",
			filter:    Filter[float64]{operator: OperatorEqual, valueType: ValueTypeNumber, value: 123.123},
			wantValid: true,
			wantErr:   false,
		},
		{
			name:      "Invalid String Type for Number(float)",
			filter:    Filter[float64]{operator: OperatorEqual, valueType: ValueTypeString, value: 123.123},
			wantValid: false,
			wantErr:   true,
		},
		{
			name:      "Invalid DateTime Type for Number(float)",
			filter:    Filter[float64]{operator: OperatorEqual, valueType: ValueTypeDatetime, value: 123.123},
			wantValid: false,
			wantErr:   true,
		},
	}

	testsDateTime := []struct {
		name    string
		filter  Filter[time.Time]
		wantErr bool
	}{
		// Test Value & ValueType
		{
			name:    "Invalid Number Type for Datetime",
			filter:  Filter[time.Time]{operator: OperatorEqual, valueType: ValueTypeNumber, value: time.Now()},
			wantErr: true,
		},
		{
			name:    "Invalid String Type for Datetime",
			filter:  Filter[time.Time]{operator: OperatorEqual, valueType: ValueTypeString, value: time.Now()},
			wantErr: true,
		},
		{
			name:    "Valid DateTime Type for Datetime",
			filter:  Filter[time.Time]{operator: OperatorEqual, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		// Test Operator
		{
			name:    "Valid Equal Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorEqual, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Valid NotEqual Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorNotEqual, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Valid LessThan Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorLessThan, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Valid GreaterThan Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorGreaterThan, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Valid GreaterThanOrEqual Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorGreaterThanOrEqual, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Valid LessThanOrEqual Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorLessThanOrEqual, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: false,
		},
		{
			name:    "Invalid Contain Operator for Datetime",
			filter:  Filter[time.Time]{operator: OperatorContain, valueType: ValueTypeDatetime, value: time.Now()},
			wantErr: true,
		},
		{
			name:    "Invalid Equal Type for Datetime",
			filter:  Filter[time.Time]{operator: OperatorEqual, valueType: ValueTypeString, value: time.Now()},
			wantErr: true,
		},
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.filter.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.filter.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}

	for _, tt := range testsFloat {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.filter.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}

	for _, tt := range testsDateTime {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.filter.Validate()
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestFilter_filtData(t *testing.T) {
	testsString := []struct {
		name   string
		filter Filter[string]
		data   string
		want   bool
	}{
		{
			name:   "Equal Strings",
			filter: Filter[string]{operator: OperatorEqual, valueType: ValueTypeString, value: "hello"},
			data:   "hello",
			want:   true,
		},
		{
			name:   "Not Equal Strings",
			filter: Filter[string]{operator: OperatorNotEqual, valueType: ValueTypeString, value: "hello"},
			data:   "world",
			want:   true,
		},
		{
			name:   "Contains String",
			filter: Filter[string]{operator: OperatorContain, valueType: ValueTypeString, value: "test"},
			data:   "this is a test",
			want:   true,
		},
		{
			name:   "Failed Filtering Equal String",
			filter: Filter[string]{operator: OperatorEqual, valueType: ValueTypeString, value: "hello"},
			data:   "hello2",
			want:   false,
		},
		{
			name:   "Failed Filtering Not Equal String",
			filter: Filter[string]{operator: OperatorNotEqual, valueType: ValueTypeString, value: "hello"},
			data:   "hello",
			want:   false,
		},
		{
			name:   "Failed Filtering Contains String",
			filter: Filter[string]{operator: OperatorContain, valueType: ValueTypeString, value: "test"},
			data:   "i'm not contain anything",
			want:   false,
		},
	}

	testsInt := []struct {
		name   string
		filter Filter[int]
		data   int
		want   bool
	}{
		{
			name:   "Not Equal Int",
			filter: Filter[int]{operator: OperatorNotEqual, valueType: ValueTypeNumber, value: 10},
			data:   11,
			want:   true,
		},
		{
			name:   "Equal Int",
			filter: Filter[int]{operator: OperatorEqual, valueType: ValueTypeNumber, value: 11},
			data:   11,
			want:   true,
		},
		{
			name:   "LessThan Int",
			filter: Filter[int]{operator: OperatorLessThan, valueType: ValueTypeNumber, value: 11},
			data:   10,
			want:   true,
		},
		{
			name:   "LessThanOrEqual Int (less)",
			filter: Filter[int]{operator: OperatorLessThanOrEqual, valueType: ValueTypeNumber, value: 12},
			data:   10,
			want:   true,
		},
		{
			name:   "LessThanOrEqual Int (equal)",
			filter: Filter[int]{operator: OperatorLessThanOrEqual, valueType: ValueTypeNumber, value: 12},
			data:   12,
			want:   true,
		},
		{
			name:   "GreaterThan Int",
			filter: Filter[int]{operator: OperatorGreaterThan, valueType: ValueTypeNumber, value: 12},
			data:   13,
			want:   true,
		},
		{
			name:   "GreaterThanOrEqual Int (greater)",
			filter: Filter[int]{operator: OperatorGreaterThanOrEqual, valueType: ValueTypeNumber, value: 12},
			data:   13,
			want:   true,
		},
		{
			name:   "GreaterThanOrEqual Int (equal)",
			filter: Filter[int]{operator: OperatorGreaterThanOrEqual, valueType: ValueTypeNumber, value: 12},
			data:   12,
			want:   true,
		},
		{
			name:   "Failed Filtering Not Equal Int",
			filter: Filter[int]{operator: OperatorNotEqual, valueType: ValueTypeNumber, value: 12},
			data:   12,
			want:   false,
		},
		{
			name:   "Failed Filtering Equal Int",
			filter: Filter[int]{operator: OperatorEqual, valueType: ValueTypeNumber, value: 12},
			data:   13,
			want:   false,
		},
		{
			name:   "Failed Filtering LessThan Int",
			filter: Filter[int]{operator: OperatorLessThan, valueType: ValueTypeNumber, value: 12},
			data:   13,
			want:   false,
		},
		{
			name:   "Failed Filtering LessThanOrEqual Int",
			filter: Filter[int]{operator: OperatorLessThan, valueType: ValueTypeNumber, value: 12},
			data:   13,
			want:   false,
		},
		{
			name:   "Failed Filtering GreaterThan Int",
			filter: Filter[int]{operator: OperatorGreaterThan, valueType: ValueTypeNumber, value: 12},
			data:   11,
			want:   false,
		},
		{
			name:   "Failed Filtering GreaterThanOrEqual Int",
			filter: Filter[int]{operator: OperatorGreaterThan, valueType: ValueTypeNumber, value: 12},
			data:   11,
			want:   false,
		},
	}

	testFloat32 := []struct {
		name   string
		filter Filter[float32]
		data   float32
		want   bool
	}{
		{
			name:   "Not Equal Float32",
			filter: Filter[float32]{operator: OperatorNotEqual, valueType: ValueTypeNumber, value: float32(12.23)},
			data:   13.24,
			want:   true,
		},
	}

	for _, tt := range testsString {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filter.filtData(tt.data), tt.want)
		})
	}

	for _, tt := range testsInt {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filter.filtData(tt.data), tt.want)
		})
	}

	for _, tt := range testFloat32 {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.filter.filtData(tt.data), tt.want)
		})
	}
}
