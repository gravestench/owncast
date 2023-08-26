package persistence

import (
	"bytes"
	"encoding/gob"
)

// KeyVal is the actual object saved to the database.
// The Value is encoded using encoding/gob.
type KeyVal struct {
	Key   string
	Value interface{}
}

func (c *KeyVal) GetStringSlice() ([]string, error) {
	decoder := c.GetDecoder()
	var result []string
	err := decoder.Decode(&result)
	return result, err
}

func (c *KeyVal) GetStringMap() (map[string]string, error) {
	decoder := c.GetDecoder()
	var result map[string]string
	err := decoder.Decode(&result)
	return result, err
}

func (c *KeyVal) GetString() (string, error) {
	decoder := c.GetDecoder()
	var result string
	err := decoder.Decode(&result)
	return result, err
}

func (c *KeyVal) GetNumber() (float64, error) {
	decoder := c.GetDecoder()
	var result float64
	err := decoder.Decode(&result)
	return result, err
}

func (c *KeyVal) GetBool() (bool, error) {
	decoder := c.GetDecoder()
	var result bool
	err := decoder.Decode(&result)
	return result, err
}

func (c *KeyVal) GetObject(result interface{}) error {
	decoder := c.GetDecoder()
	err := decoder.Decode(result)
	return err
}

func (c *KeyVal) GetDecoder() *gob.Decoder {
	valueBytes := c.Value.([]byte)
	decoder := gob.NewDecoder(bytes.NewBuffer(valueBytes))
	return decoder
}
