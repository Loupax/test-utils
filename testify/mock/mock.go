package mock

import (
	"github.com/stretchr/testify/mock"
	"sort"
)

// StringsInAnyOrder creates an assertion that guarantees the expected string array matches the actual string array,
// while ignoring their order.
//
// * Example
//
// mockObject.On("Append", mock.InAnyOrder([]string{"three", "one", "two"}))
func StringsInAnyOrder(expectation []string) interface{} {
	e := make([]string, len(expectation))
	copy(e, expectation) // Avoid altering the original arg because you never know...

	return mock.MatchedBy(func(actual []string) bool {
		if len(expectation) != len(actual) {
			return false
		}
		a := make([]string, len(actual)) // Avoid altering the original arg because you really never know.
		copy(a, actual)

		sort.Strings(e)
		sort.Strings(a)

		for k := range e {
			if e[k] != a[k] {
				return false
			}
		}
		return true
	})
}
