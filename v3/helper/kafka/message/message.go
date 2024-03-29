package queue

import "encoding/json"

type Message struct {
	Data  interface{} `json:"data"`
	Event string      `json:"event"`

	encoded []byte
	err     error
}

func (msg *Message) ensureEncoded() {
	if msg.encoded == nil && msg.err == nil {
		msg.encoded, msg.err = json.Marshal(msg)
	}
}

func (msg *Message) Length() int {
	msg.ensureEncoded()
	return len(msg.encoded)
}

func (msg *Message) Encode() ([]byte, error) {
	msg.ensureEncoded()
	return msg.encoded, msg.err
}

func Build(data interface{}, event string) *Message {
	return &Message{
		Data:  data,
		Event: event,
	}
}
