package files

import (
	"os"
	"os/exec"
)

func Mkdir(s string) {
	os.Mkdir(s, 0755)
}

func Rmdir(s string) {
	exec.Command("rm", "-rf", s).CombinedOutput()
}
