package erratum_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/errors/testutils"
	"github.com/StevenACoffman/erratum"
)

func TestHTTP(t *testing.T) {
	err := fmt.Errorf("hello")
	err = erratum.WrapWithFields(err, 302)

	// Simulate a network transfer.
	enc := errors.EncodeError(context.Background(), err)
	otherErr := errors.DecodeError(context.Background(), enc)

	tt := testutils.T{T: t}

	// Error is preserved through the network.
	tt.CheckDeepEqual(otherErr, err)

	// It's possible to extract the HTTP code.
	tt.CheckEqual(erratum.GetHTTPCode(otherErr, 100), 302)

	// If there are multiple codes, the most recent one wins.
	otherErr = erratum.WrapWithFields(otherErr, 404)
	tt.CheckEqual(erratum.GetHTTPCode(otherErr, 100), 404)

	// The code is hidden when the error is printed with %v.
	tt.CheckStringEqual(fmt.Sprintf("%v", err), `hello`)
	// The code appears when the error is printed verbosely.
	tt.CheckStringEqual(fmt.Sprintf("%+v", err), `hello
(1) http code: 302
Wraps: (2) hello
Error types: (1) *erratum.withHTTPCode (2) *errors.errorString`)
}
