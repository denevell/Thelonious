package clone

import "os"
import "os/exec"
import "path"
import "Thelonious/utils"
import "io"

func CloneProject(url string, output io.Writer) (string, error) {
	err := os.RemoveAll("src/"+path.Base(url))
	if err != nil {
		return "", nil
	}
	cmd := exec.Command("git", "clone", url, "src/"+path.Base(url))
	cmdOutput, err := cmd.CombinedOutput()
	utils.PrintAndFlush(output, string(cmdOutput))
	return path.Base(url), err
}
