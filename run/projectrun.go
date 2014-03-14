package run

import "os"
import "os/exec"
import "fmt"

func RunProject(command string) (error) {
	exec.Command("killall", "-9", command).Run()
	wd, _ := os.Getwd()
	cmd := exec.Command(wd+"/src/"+command+"/"+command)
	cmd.Dir = "src/"+command+"/"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Dir)
	fmt.Println(cmd.Path)
	err := cmd.Start()
	fmt.Println(err)
	return err
}
