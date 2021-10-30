package quay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTags(t *testing.T) {
	act, err := Tags()
	if err != nil {
		t.Error(err)
	}
	for _, a := range act {
		assert.Equal(t, "1.2.3", a.Str)
	}

}
