package secrets

import (
	"fmt"
	"testing"
)

func TestTokenBase64(t *testing.T) {
	s, _ := TokenHex(16)
	fmt.Println(s)
}
