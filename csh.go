package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
)

var cmd *exec.Cmd
var running bool

func Coolsh() {
	var username string

	user, err := user.Current()
	if err == nil {
		username = user.Username
	}

	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(username + " $ ")
		line, _ := in.ReadString('\n')
		trimLine := strings.TrimSuffix(string(line), "\n")
		split := strings.Split(trimLine, " ")

		if trimLine != "" {
			cmd = exec.Command(split[0], split[1:]...)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			c := make(chan os.Signal)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				if running {
					cmd.Process.Kill()
				}
			}()

			err := cmd.Start()
			if err != nil {
				fmt.Println("failed to execute command: ", err)
			}

			running = true
			cmd.Wait()
			running = false
		}
	}
}
