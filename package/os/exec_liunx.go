package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func main() {

	{
		dateCmd := exec.Command("date")
		dateOut, _ := dateCmd.Output()
		fmt.Println("bash > date")
		fmt.Println(string(dateOut))
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		// include argument
		_, err := exec.Command("date", "-x").Output()
		if err != nil {
			switch e := err.(type) {
			case *exec.Error:
				fmt.Println("failed executing:", err)
			case *exec.ExitError:
				fmt.Println("command exit rc =", e.ExitCode())
			default:
				panic(err)
			}
		}
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		grepCmd := exec.Command("grep", "hello")

		grepIn, _ := grepCmd.StdinPipe()
		grepOut, _ := grepCmd.StdoutPipe()

		grepCmd.Start()
		grepIn.Write([]byte("hello grep\ngoodbye grep"))
		grepIn.Close()
		grepBytes, _ := io.ReadAll(grepOut)
		grepCmd.Wait()

		fmt.Println("bash > grep hello")
		fmt.Println(string(grepBytes))
		fmt.Println(strings.Repeat("*#", 20))
	}
	{
		lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
		lsOut, _ := lsCmd.Output()

		fmt.Println("bash > ls -a -l -h")
		fmt.Println(string(lsOut))
	}
}
