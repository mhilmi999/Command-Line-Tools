package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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
				createLog := makeLogFile(convJsonCmd)
				if !createLog{
					fmt.Println("Gagal membuat log")
				}else{
					fmt.Println("Berhasil membuat log")
				}

			case convTextCmd:
				firstInput.Parse(inputArgs[4:])
				fmt.Println("Pilihan untuk konversi ke format text")
				createLog := makeLogFile(convTextCmd)
				if !createLog{
					fmt.Println("Gagal membuat log")
				}else{
					fmt.Println("Berhasil membuat log")
				}
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

func makeLogFile(params string) bool {
	msg := "ini error log"
	f, err := os.OpenFile("error", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}else{
		fmt.Println("proses membuat file")
		if params == "json"{
			convToJson(params)
		}else{
			convToText(params)
		}
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)

	return true
}

func convToJson(params string) bool {
	fmt.Println("Berhasil membuat file json")
	json, err := json.Marshal(params)
	err = ioutil.WriteFile("error.json", json, 0644)
	if err != nil {
		log.Println(err)
	}
	return true
}

func convToText(params string) bool {
	fmt.Println("Berhasil membuat file text")
	return true
}