package clone

import "os"
import "os/exec"
import "path"

func CloneProject(url string) (string, error) {
	err := os.RemoveAll("src/"+path.Base(url))
	if err != nil {
		return "", nil
	}
	cmd := exec.Command("git", "clone", url, "src/"+path.Base(url))
	err = cmd.Run()
	return path.Base(url), err
}
