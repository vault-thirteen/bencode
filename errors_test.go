package bencode

import (
	"errors"
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_combineErrors(t *testing.T) {

	type TestData struct {
		e1                    error
		e2                    error
		expectedCombinedError error
	}

	var aTest = tester.New(t)
	var tests []TestData

	// Test #1.
	tests = append(tests, TestData{
		e1:                    nil,
		e2:                    nil,
		expectedCombinedError: nil,
	})

	// Test #2.
	tests = append(tests, TestData{
		e1:                    nil,
		e2:                    errors.New("qwe"),
		expectedCombinedError: errors.New("qwe"),
	})

	// Test #3.
	tests = append(tests, TestData{
		e1:                    errors.New("qwe"),
		e2:                    nil,
		expectedCombinedError: errors.New("qwe"),
	})

	// Test #4.
	tests = append(tests, TestData{
		e1:                    errors.New("aaa"),
		e2:                    errors.New("bbb"),
		expectedCombinedError: errors.New("aaa: bbb"),
	})

	// Run the Tests.
	for _, test := range tests {
		result := combineErrors(test.e1, test.e2)
		aTest.MustBeEqual(result, test.expectedCombinedError)
	}
}
