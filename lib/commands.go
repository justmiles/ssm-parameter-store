package ssmParameterStore

import "fmt"
import "bufio"
import "os"

// CMDPull handles the command line's ssm-parameter-store get
func CMDPull(paths []string, format string, directory string) {
	p := NewParameterStatesFromSSM(paths)
	p.toDisk(directory)
}

// CMDPush executes ssm-parameter-store push from the CLI
func CMDPush(paths []string, format string, directory string, noInput bool) {
	desired := NewParameterStatesFromDisk(paths, format, directory)
	current := NewParameterStatesFromSSM(paths)
	diff, err := desired.diff(current)
	Check(err)
	if diff.String() != "\n" {
		fmt.Print(diff)
		// push without prompt
		if noInput {
			err = diff.commit()
		} else {
			// prompt user for confirmation
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(`Do you want to perform these actions?
      ssm-parameter-store will perform the actions described above.
      Only 'yes' will be accepted to approve.

      Enter a value: `)
			text, _ := reader.ReadString('\n')
			if text == "yes\n" {
				err = diff.commit()
			} else {
				fmt.Println("Error: Push cancelled.")
			}
		}
		Check(err)
		fmt.Println("Push completed.")
	} else {
		fmt.Println("Nothing to change.")
	}
}

// CMDDiff executes ssm-parameter-store diff from the CLI
func CMDDiff(paths []string, format string, directory string) {
	desired := NewParameterStatesFromDisk(paths, format, directory)
	current := NewParameterStatesFromSSM(paths)
	diff, err := desired.diff(current)
	Check(err)
	fmt.Print(diff)
}
