package run

import "os/exec"
//import "fmt"

func RunProject(command string) (error) {
	exec.Command("killall", "-9", command).Run()
	cmd := exec.Command("./"+command)
	err := cmd.Start()
	//fmt.Println(string(output))
	return err
}
