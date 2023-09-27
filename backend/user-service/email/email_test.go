package email

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsChulaStudentEmail(t *testing.T) {
	emailListsOK := []string{"6330203521@student.chula.ac.th"}
	for _, email := range emailListsOK {
		result := IsChulaStudentEmail(email)
		require.Equal(t, true, result)
	}

	emailNotOK := []string{"123@student.chula.ac.th", "6330203521@gmail.com", "string", "63302035aa@student.chula.ac.th"}
	for _, email := range emailNotOK {
		result := IsChulaStudentEmail(email)
		require.Equal(t, false, result)
	}

}

func TestEmailFormat(t *testing.T) {
	emailListsOK := []string{"abc@gmail.com", "def542@admin.com", "123456789@student.chula.ac.th", "comp34any2w@company.com", "som2wer@email.th"}
	for _, email := range emailListsOK {
		result := IsCorrectEmailFormat(email)
		require.Equal(t, true, result)
	}

	emailNotOK := []string{"abc@gmail", "123456789@student", "som2wer@.th", "som2wer@.com", "som2wer@.co.th", "string", "...@..."}
	for _, email := range emailNotOK {
		result := IsCorrectEmailFormat(email)
		require.Equal(t, false, result)
	}
}
