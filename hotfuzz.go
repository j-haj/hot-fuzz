// hotfuzz is a command line fuzzing tool
package hotfuzz

import (
	"exec"
	"fmt"
	"os"
)

func printUsage() {
	header := "hot-fuzz is a command line fuzzing tool for binaries"
	mainUsage := "Usage: hotfuzz [options] <path to binary>"
	options := "-l --list: specify word list used in generating fuzz sequences" +
						 "-r --repeat: specify the number of repetitions. Repetitions " +
						 "             represent the number of times the fuzz clusters " +
						 "             are repeated. If a word list is used, this is the " +
						 "             number of words concatenated together to create an " +
						 "             input. If alphanumeric is used, this is the number " +
						 "             of letters or numbers in the sequence." +
						 "-a --alphabet: use letters from the alphabet in the fuzz sequences" +
						 "-A --Alphabet: use both lowercase and uppercase letters" +
						 "-n --numbers: use numbers in the fuzz sequences" +
}

func main() {

	fmt.Println("main")
}
