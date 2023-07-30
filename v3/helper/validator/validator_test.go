package validator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Struct struct {
	A any    `json:"a"`
	B string `json:"b"`
}

func TestValidateRequiredExist(t *testing.T) {

	scenarios := []struct {
		desc           string
		dataStr        string
		expectedResult Struct
	}{
		{
			desc:    "test atrribute empty",
			dataStr: `{}`,
			expectedResult: Struct{
				A: nil,
				B: "",
			},
		},
		{
			desc:    "test atrribute empty",
			dataStr: `{"a": ""}`,
			expectedResult: Struct{
				A: "",
				B: "",
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			var (
				data Struct
			)

			err := json.Unmarshal([]byte(s.dataStr), &data)

			assert.Equal(t, nil, err)
			assert.Equal(t, s.expectedResult, data)
		})
	}

}
