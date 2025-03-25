package filter

import (
	"errors"
	"math"
	"strings"
	"time"
)

// The Filter struct represents the smallest unit of a filter.
// It is used to perform filtering based on the specified ValueType and Operator.
//
// Supported operators per ValueType:
// - String ValueType:
//   - Contains
//   - Equal
//   - NotEqual
//
// - Number ValueType:
//   - Equal
//   - NotEqual
//   - LessThan
//   - LessThanOrEqual
//   - MoreThan
//   - MoreThanOrEqual
//
// - Datetime ValueType:
//   - Equal
//   - NotEqual
//   - LessThan
//   - LessThanOrEqual
//   - MoreThan
//   - MoreThanOrEqual
//
// T represents the type of the Value and must match the specified ValueType.
type Filter[T Value] struct {
	operator  Operator
	valueType ValueType
	value     T
}

func NewFilter[T Value](operator Operator, valueType ValueType, value T) (Filter[T], error) {
	f := Filter[T]{
		operator:  operator,
		valueType: valueType,
		value:     value,
	}

	err := f.Validate()
	if err != nil {
		return Filter[T]{}, err
	}
	return f, nil
}

// Validate checks the validity of the Filter.
// It verifies that the actual Value of the Filter matches the specified ValueType
// and ensures that the assigned Operator is valid for the given ValueType.
func (f Filter[T]) Validate() error {
	if !validateValueType(f.valueType, f.value) {
		return errors.New("invalid value type")
	}
	if !validateOperator(f.operator, f.valueType) {
		return errors.New("invalid operator")
	}
	return nil
}

// validateValueType checks if the actual type of the Value matches the specified ValueType.
func validateValueType[T Value](valueType ValueType, value T) bool {
	switch any(value).(type) {
	case int, float64, float32:
		if valueType != ValueTypeNumber {
			return false
		}
	case string:
		if valueType != ValueTypeString {
			return false
		}
	case time.Time:
		if valueType != ValueTypeDatetime {
			return false
		}
	}
	return true
}

// validateOperator checks if the specified Operator is valid for the given ValueType.
// It ensures that:
// - ValueTypeNumber and ValueTypeDatetime do not use the OperatorContain.
// - ValueTypeString does not use operators such as LessThan, LessThanOrEqual, MoreThan, or MoreThanOrEqual.
func validateOperator(operator Operator, valueType ValueType) bool {
	switch valueType {
	case ValueTypeNumber, ValueTypeDatetime:
		if operator == OperatorContain {
			return false
		}
	case ValueTypeString:
		if operator == OperatorLessThan || operator == OperatorLessThanOrEqual {
			return false
		}
		if operator == OperatorGreaterThan || operator == OperatorGreaterThanOrEqual {
			return false
		}
	}
	return true
}

// filtData applies the filter's operator to compare the filter's value with the provided data.
// It handles various operators such as Equal, NotEqual, Contains, LessThan, LessThanOrEqual, GreaterThan, and GreaterThanOrEqual.
// Parameters:
// - data: The data to be compared against the filter's value.
// Returns:
// - bool: True if the data satisfies the filter condition, otherwise false.
func (f Filter[T]) filtData(data T) bool {
	switch f.operator {
	case OperatorEqual:
		return filtEqual(f.value, data)
	case OperatorNotEqual:
		return !filtEqual(f.value, data)
	case OperatorContain:
		return filtContains(f.value, data)
	case OperatorLessThan:
		return compareComparable(f.value, data, OperatorLessThan)
	case OperatorLessThanOrEqual:
		return compareComparable(f.value, data, OperatorLessThanOrEqual)
	case OperatorGreaterThan:
		return compareComparable(f.value, data, OperatorGreaterThan)
	case OperatorGreaterThanOrEqual:
		return compareComparable(f.value, data, OperatorGreaterThanOrEqual)
	default:
		return false
	}
}

// filtEqual checks if the filter value and the data are equal.
// It supports comparisons for numbers, strings, and time values.
// Parameters:
// - filterValue: The value defined in the filter.
// - data: The value to be compared against.
// Return:
// - bool: True if the values are considered equal, otherwise false.
func filtEqual[T Value](filterValue, data T) bool {
	switch v := any(filterValue).(type) {
	case int:
		fv := v
		intData := any(data).(int)
		if fv == intData {
			return true
		}
		return false
	case float32:
		fv := float64(v)
		floatedData := float64(any(data).(float32))
		if approximatelyEqual(fv, floatedData) {
			return true
		}
		return false
	case float64:
		floatedData := any(data).(float64)
		if approximatelyEqual(v, floatedData) {
			return true
		}
		return false
	case time.Time:
		fv := any(data).(time.Time)
		timedData := any(data).(time.Time)
		if fv.Equal(timedData) {
			return true
		}
		return false
	case string:
		fv := any(filterValue).(string)
		searchTarget := any(data).(string)
		if fv == searchTarget {
			return true
		}
		return false
	}
	return false
}

// filtContains checks if the filter value is contained within the data string.
// This function is specifically used for string comparisons.
// Parameters:
// - filterValue: The substring to search for within the data.
// - data: The string to be searched.
// Returns:
// - bool: True if the data contains the filter value, otherwise false.
func filtContains[T Value](filterValue, data T) bool {
	fv := any(filterValue).(string)
	searchTarget := any(data).(string)
	if strings.Contains(searchTarget, fv) {
		return true
	}
	return false
}

// compareComparable compares two comparable values based on the specified operator.
// It handles comparison operators like LessThan, LessThanOrEqual, GreaterThan, and GreaterThanOrEqual for numbers and time values.
// Parameters:
// - filterValue: The filter value to compare.
// - data: The data value to be compared against.
// - operator: The comparison operator to use.
// Returns:
// - bool: True if the comparison condition is met, otherwise false.
func compareComparable[T Value](filterValue, data T, operator Operator) bool {
	switch v := any(filterValue).(type) {
	case int:
		fv := float64(v)
		floatedData := float64(any(data).(int))
		return compareFloat64(operator, fv, floatedData)
	case float32:
		fv := float64(v)
		floatedData := float64(any(data).(float32))
		return compareFloat64(operator, fv, floatedData)
	case float64:
		floatedData := any(data).(float64)
		return compareFloat64(operator, v, floatedData)
	case time.Time:
		fv := any(filterValue).(time.Time)
		timedData := any(data).(time.Time)
		// TODO: 별도 함수로 분리
		if operator == OperatorLessThan {
			return fv.UnixNano() > timedData.UnixNano()
		}
		if operator == OperatorLessThanOrEqual {
			return fv.UnixNano() > timedData.UnixNano() || fv.UnixNano() == timedData.UnixNano()
		}
		if operator == OperatorGreaterThan {
			return fv.UnixNano() < timedData.UnixNano()
		}
		if operator == OperatorGreaterThanOrEqual {
			return fv.UnixNano() < timedData.UnixNano() || fv.UnixNano() == timedData.UnixNano()
		}
	}
	return false
}

func compareFloat64(operator Operator, filterValue, data float64) bool {
	if operator == OperatorEqual {
		return approximatelyEqual(filterValue, data)
	}
	if operator == OperatorNotEqual {
		return !approximatelyEqual(filterValue, data)
	}
	if operator == OperatorLessThan {
		return data < filterValue
	}
	if operator == OperatorLessThanOrEqual {
		return data < filterValue || approximatelyEqual(filterValue, data)
	}
	if operator == OperatorGreaterThan {
		return data > filterValue
	}
	if operator == OperatorGreaterThanOrEqual {
		return data > filterValue || approximatelyEqual(filterValue, data)
	}
	return false
}

// approximatelyEqual determines if two floating-point numbers are approximately equal.
// Due to the precision limitations of floating-point types, direct equality checks can be unreliable.
// This helper function checks if the two numbers are within a small threshold of each other.
// Parameters:
// - x1: The first floating-point number.
// - x2: The second floating-point number.
// Returns:
// - bool: True if the numbers are approximately equal within the defined threshold, otherwise false.
func approximatelyEqual(x1, x2 float64) bool {
	threshold := 0.00001
	if math.Abs(x1-x2) > threshold {
		return false
	}
	return true
}
