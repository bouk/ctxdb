package ctxdb

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_row_intoDBRow(t *testing.T) {
	err := errors.New("Hello")
	r := (&row{err: err}).intoDBRow()

	assert.Equal(t, err, r.Scan())
}
