package models

import (
	"errors"
	"io"

	"github.com/99designs/gqlgen/graphql"
	uuid "github.com/satori/go.uuid"
)

func MarshalUUID(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte("\""))
		w.Write([]byte(u.String()))
		w.Write([]byte("\""))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.FromString(v)

	default:
		return uuid.Nil, errors.New("UUID must be a string")
	}
}
