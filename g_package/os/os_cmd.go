package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {

	{
		cmd := exec.Command("date")
		dateOut, _ := cmd.Output()

		fmt.Println("bash > date ")
		fmt.Println(string(dateOut))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		dateout, err := exec.Command("date", "-I").Output()
		if err != nil {
			switch e := err.(type) {
			case *exec.Error:
				fmt.Println("err:", err)
			case *exec.ExitError:
				fmt.Println("command exit", e.ExitCode())
			default:
				panic(err)
			}
		}
		fmt.Println("bash > date -I ")
		fmt.Println(string(dateout))
		fmt.Println(strings.Repeat("##", 20))
	}
	{
		cmd := exec.Command("grep", "hello")

		grepIn, _ := cmd.StdinPipe()
		grepOut, _ := cmd.StdoutPipe()

		cmd.Start()
		grepIn.Write([]byte("hello grep\ngoodbye grep"))
		grepIn.Close()
		grepBytes, _ := io.ReadAll(grepOut)
		cmd.Wait()

		fmt.Println("bash > grep hello")
		fmt.Println(string(grepBytes))
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		cmd := exec.Command("bash", "-c", "ls -a -l -h")
		out, _ := cmd.Output()

		fmt.Println("bash > ls -a -l -h")
		fmt.Println(string(out))
	}
}
