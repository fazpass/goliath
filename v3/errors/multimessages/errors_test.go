package multimessages

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	var scenarios = []struct {
		desc             string
		messages         map[string]string
		messagesExpected map[string]string
	}{
		{
			desc: "happy path",
			messages: map[string]string{
				"en": "Bad Request",
				"id": "Permintaan yang Buruk",
			},
			messagesExpected: map[string]string{
				"en": "Bad Request",
				"id": "Permintaan yang Buruk",
			},
		},
		{
			desc: "unhappy path - get undefined key",
			messages: map[string]string{
				"en": "Bad Request",
			},
			messagesExpected: map[string]string{
				"en": "Bad Request",
				"id": DefaultMessage,
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.desc, func(t *testing.T) {
			ErrBadRequest := New(http.StatusBadRequest, "00", s.messages)

			for key, message := range s.messagesExpected {
				assert.Equal(t, message, ErrBadRequest.Get(key).Error())
			}

		})
	}
}
