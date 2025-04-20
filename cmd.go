package main

import (
	"fmt"
	"os/exec"
)

func RunCmd(args ...string) error {
	return exec.Command(args[0], args[1:]...).Run()
}

func CreateSession(session_name string) {
	fmt.Printf("Creating session \"%s\"\n", session_name)

	err := RunCmd("tmux", "new", "-d", "-s", session_name)
	CheckErr(err, "Error initializing session \"%s\"", session_name)
}

func RenameWindow(target_session string, target_window_index int16, new_name string) {
	fmt.Printf("Renaming window \"%s\" to \"%s\"\n", target_window_index, new_name)

	cmd := fmt.Sprintf("tmux rename-session -t %i %s", target_window_index, new_name)
	err := RunCmd("tmux", "send", "-t", target_session, cmd, "ENTER")
	CheckErr(err, "Error renaming window \"%s\"", target_window_index)
}

func CreateWindow(target_session string, window_name string) {
	fmt.Printf("Creating window \"%s\"\n", window_name)

	err := RunCmd("tmux", "send", "-t", target_session, "tmux neww -dn"+window_name, "ENTER")
	CheckErr(err, "Error creating window \"%s\"", window_name)
}

func ChangeSessionCWD(target_session string, cwd string) {
	// maybe doesnt work in windows?
	// tmux attach -t test -c ~/Dev/react/manybrew/ <&- 2>/dev/null
	// fmt.Printf("Current cwd: %s\n", cwd)
	fmt.Printf("Changing work directory to \"%s\"\n", cwd)

	err := RunCmd("tmux", "attach", "-t", target_session, "-c", cwd, "<&- 2>/dev/null")
	CheckErr(err, "Error changing cwd")
}

func ClearSessionTerminal(target_session string) {
	err := RunCmd("tmux", "send", "-t", target_session, "clear", "ENTER")
	CheckErr(err, "Error changing cwd")
}
