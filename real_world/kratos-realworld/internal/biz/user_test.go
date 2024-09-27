package biz

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hashPassword("aaa")
}

func TestVerifyPassword(t *testing.T) {
	hash := hashPassword("aaa")

	b1 := verifyPassword(hash, "aaa")
	require.True(t, b1)

	b2 := verifyPassword(hash, "aaj")
	require.False(t, b2)
}
