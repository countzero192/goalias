package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// path to config file
func get_rc_path(sh string) string {
	if sh == "bash" {
		return fmt.Sprintf("%s/.bashrc", os.Getenv("HOME"))
	} else if sh == "zsh" {
		return fmt.Sprintf("%s/.zshrc", os.Getenv("HOME"))
	} else if sh == "tcsh" {
		return fmt.Sprintf("%s/.tcshrc", os.Getenv("HOME"))
	} else {
		return fmt.Sprintf("%s/.bashrc", os.Getenv("HOME"))
	}
}

func make_alias_command(name, com, sh string) string {
	if sh == "zsh" || sh == "bash" {
		return fmt.Sprintf("alias %s=\"%s\"\n", name, com)
	} else if sh == "tcsh" {
		return fmt.Sprintf("alias %s \"%s\"\n", name, com)
	} else {
		return fmt.Sprintf("alias %s=\"%s\"\n", name, com)
	}
}

func main() {
	// parse shell args
	var shell = flag.String("s", "bash", "Your shell")
	var name_of_alias = flag.String("n", "", "Name of your alias")
	var command = flag.String("c", "", "Command")
	flag.Parse()
	// if user don't give name as argument, we ask him to enter name
	if *name_of_alias == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter name of alias -> ")
		text, _ := reader.ReadString('\n')
		*name_of_alias = strings.Replace(text, "\n", "", -1)
	}
	// also with command
	if *command == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter command -> ")
		text, _ := reader.ReadString('\n')
		*command = strings.Replace(text, "\n", "", -1)
	}

	path_to_file := get_rc_path(*shell)

	file, err := os.OpenFile(path_to_file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	com := make_alias_command(*name_of_alias, *command, *shell)

	if _, err := file.WriteString(com); err != nil {
		panic(err)
	}

	fmt.Println("Successful")
}
