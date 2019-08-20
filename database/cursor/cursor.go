package cursor

type Cursor interface {
	Serialize() string
	GenerateAscendingSQLConditions(tablePrefix string, cursorColumn string, tiebreakerColumn string) (string, []interface{})
}
