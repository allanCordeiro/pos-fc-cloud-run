package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCep(t *testing.T) {
	scenarios := []struct {
		name string
		cep  string
		err  error
	}{
		{
			name: "given a valid cep code when call validate should return true",
			cep:  "01211-100",
			err:  nil,
		},
		{
			name: "given a valid cep code but less than 8 numbers when call validate should return true",
			cep:  "4266-060",
			err:  nil,
		},
		{
			name: "given a cep code with more than 8 numbers when call validate should return false",
			cep:  "401211-100",
			err:  ErrInvalidZipCode,
		},
		{
			name: "given a cep code with invalid chars when call validate should return false",
			cep:  "01211-*100",
			err:  ErrInvalidZipCode,
		},
		{
			name: "given an invalid cep code when call validate should return false",
			cep:  "12345a",
			err:  ErrInvalidZipCode,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			cep := NewCep(test.cep)
			assert.Equal(t, test.err, cep.Validate())
		})
	}
}
