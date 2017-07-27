//+build darwin

package audio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOSXSupport(t *testing.T) {
	a, err := EncodeBytes(Encoding{
		[]byte{},
		Format{},
		CanLoop{},
	})
	assert.Nil(t, err)
	err = <-a.Play()
	assert.NotNil(t, err)
	err = a.Stop()
	assert.NotNil(t, err)
}
