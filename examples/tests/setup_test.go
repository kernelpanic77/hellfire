package tests

import (
	"os"
	"testing"

	"github.com/kernelpanic77/hellfire/pkg/hellfire"
)

func TestMain(m *testing.M) {
	exitcode := hellfire.Main(m)
	os.Exit(exitcode)
}


// func TestMain(m *testing.M) {
//     //fmt.Println("TestMain is running...")
//     // Perform setup


//     // Run tests
//     exitCode := hellfire.Main(m)

//     // Perform teardown

//     os.Exit(exitCode)
// }