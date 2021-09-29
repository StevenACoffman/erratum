package erratum_test

import (
	"fmt"
	"testing"

	"github.com/StevenACoffman/erratum"
	"github.com/cockroachdb/errors/testutils"
)

func TestHTTP(t *testing.T) {
	err := fmt.Errorf("hello")
	fields := erratum.Fields{"one": "oneval", "two": "twoval"}
	err = erratum.WrapWithFields(err, fields)

	tt := testutils.T{T: t}

	// It's possible to extract the Fields
	tt.CheckEqual(fmt.Sprintf("%v", erratum.GetFields(err)), fmt.Sprintf("%v", fields))

	// If there are multiple codes, the most recent one wins.
	otherFields := erratum.Fields{"three": "threeval", "four": "fourval"}
	otherErr := erratum.WrapWithFields(err, otherFields)
	tt.CheckEqual(fmt.Sprintf("%v", erratum.GetFields(otherErr)), fmt.Sprintf("%v", otherFields))

	// The code is hidden when the error is printed with %v.
	tt.CheckStringEqual(fmt.Sprintf("%v", err), `hello`)
	// The code appears when the error is printed verbosely.
	tt.CheckStringEqual(fmt.Sprintf("%+v", err), `hello
(1) fields: [one:oneval,two:twoval]
Wraps: (2) hello
Error types: (1) *erratum.withFields (2) *errors.errorString`)
}
