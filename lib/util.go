package ssmParameterStore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	yaml "gopkg.in/yaml.v2"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc = ssm.New(sess)
)

// getSSMParameters recursively loops through AWS SSM for the given path and returns a slice of SSM Parameters
func getSSMParameters(path string) (parameters []ssm.Parameter) {

	ssmOpts := ssm.GetParametersByPathInput{
		Path:           aws.String(path),
		WithDecryption: aws.Bool(true),
		MaxResults:     aws.Int64(10),
		Recursive:      aws.Bool(true),
	}

	// Loop through the SSM GetParametersByPathInput call until Pagination is complete
	for {

		// perform the request
		ssmResponse, err := svc.GetParametersByPath(&ssmOpts)
		Check(err)

		// range over response and store results in memory
		for _, parameter := range ssmResponse.Parameters {
			parameters = append(parameters, *parameter)
		}

		// if pagination NextToken exists, set it and continue loop. otherwise break loop
		if ssmResponse.NextToken == nil {
			break
		} else {
			ssmOpts.NextToken = ssmResponse.NextToken
		}
	}

	return parameters
}

// pathAndKey takes a file path and returns the path and basename
// eg, /uat/myapp/mypropertyname becomes /uat/myapp mypropertyname
func pathAndKey(p *string) (string, string) {
	s := strings.Split(*p, "/")
	return strings.Join(s[0:len(s)-1], "/"), s[len(s)-1]
}

// NewParameterStatesFromDisk reads the saved parameter store and returns a ParameterStates
func NewParameterStatesFromDisk(paths []string, format, directory string) ParameterStates {
	var p = make(ParameterStates)

	for _, localpath := range paths {

		err := filepath.Walk(directory+localpath,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if !strings.Contains(path, format) {
					return nil
				}
				x := strings.Replace(path, directory, "", 1)
				var re = regexp.MustCompile(fmt.Sprintf(`\.%s$`, format))
				ssmPath := re.ReplaceAllString(x, "")
				if ssmPath == "/" {
					ssmPath = ""
				}
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				p[ssmPath] = &ParameterState{}

				err = yaml.Unmarshal(data, p[ssmPath])
				if err != nil {
					return err
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	return p
}

// Check errors if exist and exit
func Check(err error) {
	if err != nil {
		fmt.Printf("ERROR\t%s\n", err.Error())
		os.Exit(1)
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
