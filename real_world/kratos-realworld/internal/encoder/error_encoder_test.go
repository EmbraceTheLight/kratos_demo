package encoder

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHTTPErrorEncoder(t *testing.T) {
	a := &HTTPError{
		Errors: &Errors{
			Field: []string{"can't be empty"},
		},
	}

	b, err := json.Marshal(a)
	require.NoError(t, err)
	fmt.Printf("%s", string(b))
}
