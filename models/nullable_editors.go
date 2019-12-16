package models

import "github.com/gofrs/uuid"

// Allows editing of strings that can be null
// If the value in this type is null, then the outer value will be set to null
// However, if the outer value is null (meaning this whole data structure is omitted), nothing will change
type NullableStringEditor struct {
	Value *string `json:"value"`
}

func (n *NullableStringEditor) ContainsValue() bool {
	return n != nil
}

// Allows editing of IDs that can be null
// If the value in this type is null, then the outer value will be set to null
// However, if the outer value is null (meaning this whole data structure is omitted), nothing will change
type NullableIDEditor struct {
	Value *uuid.UUID `json:"value"`
}

func (n *NullableIDEditor) ContainsValue() bool {
	return n != nil
}

// Allows editing of ints that can be null
// If the value in this type is null, then the outer value will be set to null
// However, if the outer value is null (meaning this whole data structure is omitted), nothing will change
type NullableIntEditor struct {
	Value *int `json:"value"`
}

func (n *NullableIntEditor) ContainsValue() bool {
	return n != nil
}
