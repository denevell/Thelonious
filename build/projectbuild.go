package build

import "os"
import "os/exec"
import "path"
//import "fmt"

func BuildProject(url string) (string, error) {
	cmd := exec.Command("go", "build", path.Base(url))
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd.Env = append(cmd.Env, "GOPATH="+pwd)
	_, err = cmd.Output()
	return path.Base(url), err
}
