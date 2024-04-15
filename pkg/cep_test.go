package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCep(t *testing.T) {
	scenarios := []struct {
		name    string
		cep     string
		isValid bool
	}{
		{
			name:    "given a valid cep code when call validate should return true",
			cep:     "01211-100",
			isValid: true,
		},
		{
			name:    "given a valid cep code but less than 8 numbers when call validate should return true",
			cep:     "4266-060",
			isValid: true,
		},
		{
			name:    "given a cep code with more than 8 numbers when call validate should return false",
			cep:     "401211-100",
			isValid: false,
		},
		{
			name:    "given a cep code with invalid chars when call validate should return false",
			cep:     "01211-*100",
			isValid: false,
		},
		{
			name:    "given an invalid cep code when call validate should return false",
			cep:     "12345a",
			isValid: false,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			cepCode := NewCep(test.cep)
			assert.Equal(t, test.isValid, cepCode.IsCepCodeValid())
		})
	}
}
