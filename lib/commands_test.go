package ssmParameterStore

import (
	"testing"
)

func TestRun(t *testing.T) {
	CMDPull([]string{"/ops", "/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch")
	CMDPush([]string{"/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch", true)
	CMDDiff([]string{"/ops"}, "yaml", "/home/justmiles/go/src/github.com/justmiles/ssm-parameter-store/scratch")
}
