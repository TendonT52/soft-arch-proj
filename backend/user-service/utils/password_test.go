package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyPassword(t *testing.T) {
	password := "123456"

	hashedPassword := HashPassword(password, 1)

	err := VerifyPassword(hashedPassword, password, 1)
	require.NoError(t, err)

	err = VerifyPassword(hashedPassword, password, 2)
	require.Error(t, err)

	err = VerifyPassword(hashedPassword, "1234567", 1)
	require.Error(t, err)
}
