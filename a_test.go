package tr

import (
	"os"
	"os/exec"
	"testing"
)

func TestStringer(t *testing.T) {
	out, _ := exec.Command("go", "run", "golang.org/x/tools/cmd/stringer").CombinedOutput()
	t.Fatal(string(out))
}

func TestWd(t *testing.T) {
	out, _ := os.Getwd()
	t.Fatal(string(out))
}
