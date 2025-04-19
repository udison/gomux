package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	session_name := "mb"
	cwd := "~/Dev/react/manybrew/"

	fmt.Printf("Creating session %s\n", session_name)
	cmd := exec.Command("tmux", "new", "-d", "-s", session_name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	cd_cmd := "cd " + cwd
	cmd = exec.Command("tmux", "send", "-t", session_name, cd_cmd, "ENTER")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current cwd: %s\n", cwd)

	cmd = exec.Command("tmux", "send", "-t", session_name, "tmux neww -dnnvim", "ENTER")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	// tmux send-keys -t mb:nvim nvim C-m
	cmd = exec.Command("tmux", "send-keys", "-t", session_name+":nvim", "nvim", "C-m")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
