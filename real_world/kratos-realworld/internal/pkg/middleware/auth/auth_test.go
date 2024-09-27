package auth

import (
	"github.com/spewerspew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tokenString := GenerateToken("b25cbd2b77454e275ce1236860839130f5fb615020bdb4f025308c6b0f0dc5ed", "aaa")
	spew.Dump(tokenString)
}
