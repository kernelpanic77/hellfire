package tests

import (
	"os"
	"testing"

	"github.com/kernelpanic77/hellfire/pkg/hellfire"
)

func TestMain(m *testing.M) {
	os.Exit(hellfire.Main(m))
}
