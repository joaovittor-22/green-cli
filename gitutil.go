package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const file = "green_file.txt"

func MakeCommit(date time.Time, index int) {
	dateStr := date.Format("Mon Jan 2 15:04:05 2006 -0700")
	message := fmt.Sprintf("Backdated commit %d - %s\n", index, dateStr)

	// Append to file
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	defer f.Close()
	if _, err := f.WriteString(message); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	env := append(os.Environ(),
		"GIT_AUTHOR_DATE="+dateStr,
		"GIT_COMMITTER_DATE="+dateStr,
	)

	runCommand("git", "add", file)
	runCommandWithEnv(env, "git", "commit", "-m", fmt.Sprintf("Backdated commit %d", index))
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command failed: %s %v\nError: %v", name, args, err)
	}
}

func runCommandWithEnv(env []string, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Command (with env) failed: %s %v\nError: %v", name, args, err)
	}
}