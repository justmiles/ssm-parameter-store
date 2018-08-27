package ssmParameterStore

import (
	"testing"
)

func TestRun(t *testing.T) {
	CMDPull([]string{"/dev", "/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch")
	// CMDPush([]string{"/dev", "/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch")
	// CMDDiff([]string{"/dev", "/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch")
}
