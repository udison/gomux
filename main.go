package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Usage:\n\tgomux <session_name>\n")
		return
	}

	session_name := args[0]

	conf, err := os.ReadFile("./.conf.gomux")
	CheckErr(err, "Error opening the configuration file")

	fmt.Println(string(conf))

	CreateSession(session_name)
	// ChangeSessionCWD(session_name, "~/Dev/react/manybrew")
	RenameWindow(session_name, 1, "nvim")
	CreateWindow(session_name, "test")
	CreateWindow(session_name, "new_window")
	ClearSessionTerminal(session_name)
}

/* Steps to creating a session
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

*/
