package lumberjack

import (
	"reflect"
	"testing"
)

// assert fails the test if condition is false.
func assert(condition bool, t testing.TB, msg string, v ...any) {
	t.Helper()
	if !condition {
		t.Fatalf(msg, v...)
	}
}

// equals fails the test if the two values are not equal according to
// reflect.DeepEqual.
func equals(exp, act any, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("exp: %v (%T), got: %v (%T)", exp, exp, act, act)
	}
}

// isNil reports a failure if the given value is not nil.  Note that values
// which cannot be nil will always fail this check.
func isNil(obtained any, t testing.TB) {
	t.Helper()
	if !_isNil(obtained) {
		t.Fatalf("expected nil, got: %v", obtained)
	}
}

// notNil reports a failure if the given value is nil.
func notNil(obtained any, t testing.TB) {
	t.Helper()
	if _isNil(obtained) {
		t.Fatalf("expected non-nil, got: %v", obtained)
	}
}

// _isNil is a helper function for isNil and notNil, and should not be used
// directly.
func _isNil(obtained any) bool {
	if obtained == nil {
		return true
	}

	switch v := reflect.ValueOf(obtained); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}

	return false
}
