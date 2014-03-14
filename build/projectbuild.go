package build

import "os"
import "os/exec"
import "path"
import "Thelonious/utils"
import "io"
import "fmt"

func BuildProject(url string, w io.Writer) (string, error) {
	fmt.Println(url)
	cmd := exec.Command("go", "build", path.Base(url))
	cmd.Dir = "src/"+path.Base(url)
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	cmd.Env = append(cmd.Env, "GOPATH="+pwd)
	cmdOutput, err := cmd.CombinedOutput()
	utils.PrintAndFlush(w, string(cmdOutput))	
	return path.Base(url), err
}
