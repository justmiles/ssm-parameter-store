package ssmParameterStore

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/fatih/color"
)

// Diff represents data to be added and removed
type Diff struct {
	additions []*ssm.PutParameterInput
	deletes   []*ssm.DeleteParametersInput
	asVisual  []string
}

func (diff Diff) String() string {
	return strings.Join(diff.asVisual, "\n") + "\n"
}

func (diff Diff) commit() error {
	for _, addition := range diff.additions {
		_, err := svc.PutParameter(addition)
		if err != nil {
			return err
		}
	}

	for _, deleteInput := range diff.deletes {
		_, err := svc.DeleteParameters(deleteInput)
		if err != nil {
			return err
		}
	}

	return nil
}

// AppendDeleteChange appends a line to delete
func (diff *Diff) AppendDeleteChange(path string) error {
	diff.asVisual = append(diff.asVisual, color.RedString(fmt.Sprintf("-\t%s", path)))

	if diff.deletes == nil {
		diff.deletes = append(diff.deletes, &ssm.DeleteParametersInput{})
	}

	for i, input := range diff.deletes {
		if len(input.Names) < 10 {
			diff.deletes[i].Names = append(input.Names, aws.String(path))
		} else if len(diff.deletes) < (i + 2) {
			diff.deletes = append(diff.deletes, &ssm.DeleteParametersInput{
				Names: aws.StringSlice([]string{path}),
			})
		}
	}
	return nil
}

// AppendAddChange appends a line to delete
func (diff *Diff) AppendAddChange(path, desiredValue, currentValue string) error {
	if currentValue != "" {
		diff.asVisual = append(diff.asVisual, color.YellowString(fmt.Sprintf("~\t%s\t%s  -->  %s", path, currentValue, desiredValue)))
	} else {
		diff.asVisual = append(diff.asVisual, color.GreenString(fmt.Sprintf("+\t%s\t%s", path, desiredValue)))
	}

	diff.additions = append(diff.additions, &ssm.PutParameterInput{
		// KeyId: ""
		Name:      aws.String(path),
		Overwrite: aws.Bool(true),
		Type:      aws.String("String"),
		Value:     aws.String(desiredValue),
	})

	return nil
}
