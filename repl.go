package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func shouldContinue(text string) bool {
	if strings.EqualFold(".exit", text) {
		return false
	}
	return true
}

func printPrompt() {
	fmt.Print("db > ")
}

func printLine(text string) {
	printPrompt()
	fmt.Println(text)
}
func help() {
	printLine("Welcome to gosqlite!")
	fmt.Println(".help For help.")
	fmt.Println(".exit To quit.")
	fmt.Println("")
}

func cls() {
	cmd := exec.Command(".clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	commands := map[string]interface{}{
		".help": help,
		".cls":  cls,
	}
	reader := bufio.NewReader(os.Stdin)
	help()
	printPrompt()
	text := get(reader)
	for ; shouldContinue(text); text = get(reader) {
		if value, exists := commands[text]; exists {
			value.(func())()
		} else {
			printLine(text + " does not exist")
		}
		printPrompt()
	}
	fmt.Println("Bye!")

}
