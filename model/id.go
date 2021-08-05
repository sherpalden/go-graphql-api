package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"io"

	"github.com/google/uuid"
)

//ID -> new datatype
type ID uuid.UUID

// StringToID -> parse string to id
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

// String -> String Representation of ID
func (binary ID) String() string {
	return uuid.UUID(binary).String()
}

//GormDataType -> sets type to binary(16)
func (binary *ID) GormDataType() string {
	return "binary(16)"
}

// Scan --> From DB
func (binary *ID) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal  value:", value))
	}
	uuidVal, err := uuid.Parse(strVal)
	if err != nil {
		return errors.New(fmt.Sprint("Failed to parse  value:", value))
	}
	*binary = ID(uuidVal)
	return err
}

// Value -> TO DB
func (binary ID) Value() (driver.Value, error) {
	return uuid.UUID(binary).MarshalBinary()
}

// MarshalJSON -> convert to json string
func (binary ID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(binary)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

// UnmarshalJSON -> convert from json string
func (binary *ID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*binary = ID(s)
	return err
}

// MarshalGQL implements the graphql.Marshaler interface
func (binary ID) MarshalGQL(w io.Writer) {
	idString := uuid.UUID(binary).String()
	w.Write([]byte(idString))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (binary *ID) UnmarshalGQL(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("id must be a string")
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	*binary = ID(id)
	return nil
}
