package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/term"
	_ "modernc.org/sqlite"
)

const (
	connectionString = `file:nihongo.db?mode=rwc&_journal_mode=WAL&_busy_timeout=10000`
)

func readLine(prompt string) (string, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return "", fmt.Errorf("pipe not supported")
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", fmt.Errorf("failed setting stdin to raw mode: %w", err)
	}
	tty := term.NewTerminal(os.Stdin, prompt)

	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return "", fmt.Errorf("failed to get terminal size: %w", err)
	}
	tty.SetSize(width, height)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			tty.Write([]byte("test \n"))
		}
	}()

	line, err := tty.ReadLine()
	_ = term.Restore(int(os.Stdin.Fd()), oldState)

	if err != nil {
		return "", fmt.Errorf("failed to read from stdin: %w", err)
	}
	return line, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	db, err := sqlx.Open("sqlite", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	line, err := readLine("> ")
	if err != nil {
		log.Printf("failed to read line: %v\n", err)
	}
	fmt.Printf("%q\n", line)
}
