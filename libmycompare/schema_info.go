package libmycompare

import "fmt"

type SchemaInfo struct {
	FieldName   string
	FieldType   string
	FieldNull   string
	FieldKey    string
	FieldDefult string
	FieldExtra  string
}

func (s SchemaInfo) String() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		s.FieldName,
		s.FieldType,
		s.FieldNull,
		s.FieldKey,
		s.FieldDefult,
		s.FieldExtra)
}
