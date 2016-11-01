package julius_wrapper

import (
  "testing"
  "os/exec"
  "github.com/stretchr/testify/assert"
)

func TestFoundJulius(t *testing.T) {
  _, err := exec.LookPath("julius.dSYM")
  assert.Nil(t, err)
}
