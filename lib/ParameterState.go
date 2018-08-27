package ssmParameterStore

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

// ParameterState struct represents a parameter store's value based on path
type ParameterState struct {
	EncryptionKey *string           `json:"EncryptionKey,omitempty" yaml:"EncryptionKey,omitempty"`
	EncryptedKeys []string          `json:"EncryptedKeys,omitempty" yaml:"EncryptedKeys,omitempty"`
	Parameters    map[string]string `json:"Parameters,omitempty" yaml:"Parameters,omitempty"`
}

func (p *ParameterState) json() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

func (p *ParameterState) yaml() ([]byte, error) {
	return yaml.Marshal(p)
}
