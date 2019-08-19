package cursor

import (
	"testing"
	"time"

	"github.com/gofrs/uuid"
)

func Test_deserializeAndSerialize_TimeUUIDCursor(t *testing.T) {

	cursor := TimeUUIDCursor{
		Time: time.Now(),
		UUID: uuid.Must(uuid.NewV4()),
	}

	serialized := cursor.Serialize()

	deserialized, err := DeserializeCursor(&serialized)
	if err != nil {
		t.Fatalf("Failed to deserialize cursor, %s", err.Error())
	}

	if !deserialized.Time.Equal(cursor.Time) {
		t.Fatal("Times do not match")
	}

	if deserialized.UUID != cursor.UUID {
		t.Fatal("UUIDs do not match")
	}

}

func Test_deserialize_handlesNilString(t *testing.T) {

	deserialized, err := DeserializeCursor(nil)
	if err != nil {
		t.Fatalf("Failed to deserialize cursor, %s", err.Error())
	}

	if deserialized != nil {
		t.Fatal("Deserialized should be nil")
	}
}
