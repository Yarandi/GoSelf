package helpers

import (
	"fmt"
	"os"
)

// VarDump will print out any number of variables given to it
func VarDump(myVar ...interface{}) {
	fmt.Printf("\n\n\n %v \n", myVar)
}

// DD will print out variables given to it (like varDump()) but it will stop execution from continuing
func DD(myVar ...interface{}) {
	VarDump(myVar...)
	os.Exit(1)
}
