package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func bootApp() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scaned := scanner.Scan()
		if !scaned {
			return fmt.Errorf("Scanner failed!")
		}

		input := scanner.Text()
		inputCleaned := cleanInput(input)

		if len(inputCleaned) <= 0 {
			continue
		}
		requestedCommand := inputCleaned[0]
		c, ok := pokedexCliCommandsMap[requestedCommand]

		if !ok {
			fmt.Printf("Command %s is invalid", requestedCommand)
			os.Exit(1)
		}

		c.callback()

	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	result := strings.Fields(input)
	return result
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
