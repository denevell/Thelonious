package build

import "os"
import "os/exec"
import "path"
import "Thelonious/utils"
import "io"

func BuildProject(url string, w io.Writer) (string, error) {
	cmd := exec.Command("go", "build", path.Base(url))
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd.Env = append(cmd.Env, "GOPATH="+pwd)
	cmdOutput, err := cmd.CombinedOutput()
	utils.PrintAndFlush(w, string(cmdOutput))	
	return path.Base(url), err
}
