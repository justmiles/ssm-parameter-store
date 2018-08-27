package ssmParameterStore

import "fmt"

// CMDPull handles the command line's ssm-parameter-store get
func CMDPull(paths []string, format string, directory string) {
	p := NewParameterStatesFromSSM(paths)
	p.toDisk(directory)
}

// CMDPush executes ssm-parameter-store push from the CLI
func CMDPush(paths []string, format string, directory string) {
	desired := NewParameterStatesFromDisk(paths, format, directory)
	current := NewParameterStatesFromSSM(paths)
	diff, err := desired.diff(current)
	Check(err)
	fmt.Print(diff)
	err = diff.commit()
	Check(err)
}

// CMDDiff executes ssm-parameter-store diff from the CLI
func CMDDiff(paths []string, format string, directory string) {
	desired := NewParameterStatesFromDisk(paths, format, directory)
	current := NewParameterStatesFromSSM(paths)
	diff, err := desired.diff(current)
	Check(err)
	fmt.Print(diff)
}
