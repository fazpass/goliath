package uuid

import "github.com/gofrs/uuid"

func Generate() string {
	var uuidV4, _ = uuid.NewV4()
	return uuidV4.String()
}
