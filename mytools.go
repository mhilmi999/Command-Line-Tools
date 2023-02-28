package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	firstCmd, inputArg := userPref()
	convOption(firstCmd, inputArg)
}

func userPref() (*flag.FlagSet, []string) {
	firstCmd := flag.NewFlagSet("/var/log/nginx/error.log", flag.ExitOnError)
	inputArg := os.Args
	if len(os.Args) < 2 {
		fmt.Println("kurang parameter 'json' atau 'text' untuk konversi")
		os.Exit(1)
	}
	return firstCmd, inputArg
}

func convOption(firstInput *flag.FlagSet, inputArgs []string) {
	convCmd := "-t"
	convJsonCmd := "json"
	convTextCmd := "text"
	helpCmd := "-h"
	expectedPath := "/var/log/nginx/error.log"
	switch inputArgs[1] {
	case expectedPath:
		switch inputArgs[2] {
		case convCmd:
			switch inputArgs[3] {
			case convJsonCmd:
				firstInput.Parse(inputArgs[4:])
				fmt.Println("Pilihan untuk konversi ke format json")

			case convTextCmd:
				firstInput.Parse(inputArgs[4:])
				fmt.Println("Pilihan untuk konversi ke format text")
			}
		}

	case helpCmd:
		fmt.Println("      Secara default akan terkonversi menjadi PlainText")
		fmt.Println("	-t json Mengkonversi menjadi file json")
		fmt.Println("	-t text Mengkonversi menjadi file text")
		fmt.Println("	-o Melakukan kustomisasi tempat peletakan file hasil konversi")
		fmt.Println("	-h Menampilkan bantuan")

	default:
		fmt.Println("subcommand tidak dikenali")
		os.Exit(1)
	}
}
