package ssmParameterStore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	yaml "gopkg.in/yaml.v2"
)

// ParameterStates is the format written to or read from disk
type ParameterStates map[string]*ParameterState

func (p *ParameterStates) json() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

func (p *ParameterStates) yaml() ([]byte, error) {
	return yaml.Marshal(p)
}

func (p *ParameterStates) toDisk(directory string) error {
	for key, ps := range *p {
		path, file := pathAndKey(&key)
		fullpath := directory + path
		fullname := fmt.Sprintf("%s/%s.%s", fullpath, file, "yaml")
		fmt.Printf("Writing %s.yaml to %s\n", file, fullpath)
		err := os.MkdirAll(fullpath, os.ModePerm)
		if err != nil {
			return err
		}

		contents, err := ps.yaml()
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(fullname, contents, 0644)
		if err != nil {
			return err
		}

	}

	return nil
}

func (p *ParameterStates) diff(current ParameterStates) (diffs Diff, err error) {

	for path, ps := range *p {
		for key, value := range ps.Parameters {

			// Add because the path does not exist in current
			if current[path] == nil {
				diffs.AppendAddChange(fmt.Sprintf("%s/%s", path, key), value, "")

				// Add because the key does not exist in current
			} else if current[path].Parameters[key] == "" {
				diffs.AppendAddChange(fmt.Sprintf("%s/%s", path, key), value, "")

				// Add because the key is not up to date in current
			} else if value != current[path].Parameters[key] {
				diffs.AppendAddChange(fmt.Sprintf("%s/%s", path, key), value, current[path].Parameters[key])

			}
		}
	}

	for path, ps := range current {
		for key := range ps.Parameters {

			// Delete because desired path does not exist
			if (*p)[path] == nil {
				diffs.AppendDeleteChange(fmt.Sprintf("%s/%s", path, key))

				// Delete because desired key does not exist
			} else if (*p)[path].Parameters[key] == "" {
				diffs.AppendDeleteChange(fmt.Sprintf("%s/%s", path, key))
			}
		}
	}

	return
}

func (p *ParameterStates) buildFromSSMParameters(paths []string) {

	var ssmParams []ssm.Parameter
	for _, path := range paths {
		ssmParams = append(ssmParams, getSSMParameters(path)...)
	}

	if *p == nil {
		*p = make(ParameterStates)
	}

	for _, parameter := range ssmParams {
		path, key := pathAndKey(parameter.Name)
		if _, ok := (*p)[path]; !ok {
			(*p)[path] = &ParameterState{
				Parameters: make(map[string]string),
			}
		}
		(*p)[path].Parameters[key] = *parameter.Value
		if *parameter.Type == "SecureString" {
			(*p)[path].EncryptedKeys = append((*p)[path].EncryptedKeys, key)
		}
	}
}

// NewParameterStatesFromSSM reads the current parameter store in AWS returns a ParameterStates
func NewParameterStatesFromSSM(paths []string) (p ParameterStates) {
	p.buildFromSSMParameters(paths)
	return
}
