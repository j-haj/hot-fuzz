// hotfuzz is a command line fuzzing tool
package main

import (
	"fmt"
	"hotfuzz"
	"os"
	//	"os/exec"
)

func printUsage() {
	header := "hot-fuzz is a command line fuzzing tool for binaries. Fuzz patterns will be piped into the binary."
	mainUsage := "\nUSAGE:\thotfuzz [OPTIONS] <path to binary>\n"
	options := "OPTIONS:\n" +
		"-l --list <file>: specify word list used in generating fuzz sequences\n" +
		"-r --repeat: specify the number of repetitions. Repetitions\n" +
		"             represent the number of times the fuzz clusters\n" +
		"             are repeated. If a word list is used, this is the\n" +
		"             number of words concatenated together to create an\n" +
		"             input. If alphanumeric is used, this is the number\n" +
		"             of letters or numbers in the sequence.\n" +
		"-a --alphabet: use letters from the alphabet in the fuzz sequences\n" +
		"-A --Alphabet: use both lowercase and uppercase letters\n" +
		"-n --numbers: use numbers in the fuzz sequences\n"
	fmt.Println(header)
	fmt.Println(mainUsage)
	fmt.Println(options)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	fmt.Println("main")
}
