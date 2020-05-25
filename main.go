package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func cmdExec(c *Client, userid string, args []string) error {
	endpoint := "https://noticehub.herokuapp.com/telegram/send/" + userid
	baseCmd := args[0]
	cmdArgs := args[1:]
	cmd := exec.Command(baseCmd, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	data := map[string]string{"msg": "Your program [" + strings.Join(args, " ") + "] has been executed successfully"}
	if err != nil {
		data = map[string]string{"msg": "Your programm returned non-zero result: " + err.Error()}
	}
	c.Post(endpoint, data)
	return nil
}
func main() {
	userid := os.Args[1]
	program := os.Args[2:]
	c := NewClient()
	fmt.Println(program)
	cmdExec(c, userid, program)
}
