package assert

import (
	"testing"
)

func True(t *testing.T, result bool, testCase string) {
	Boolean(t, result, true, testCase)
}

func False(t *testing.T, result bool, testCase string) {
	Boolean(t, result, false, testCase)
}

func Boolean(t *testing.T, result, expected bool, testCase string) {
	if result != expected {
		t.Errorf("returned unexpected result for %v: \ngot:  `%v` \nwant: `%v`\n", testCase, result, expected)
	}
}

func Equal(t *testing.T, expected, actual interface{}, testCase string) {
	if expected != actual {
		t.Errorf("returned unexpected result for %v: \ngot:  `%v` \nwant: `%v`\n", testCase, actual, expected)
	}
}
