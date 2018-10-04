package main 

import (
	"flag"
	"fmt"
	"os"
	"log"

	"fptai-sdk-go"
)

var target string
var inputFP string
var token string

func main() {
	trainCmd := flag.NewFlagSet("train", flag.ExitOnError)
	AddSharedFlags(trainCmd)

	testCmd := flag.NewFlagSet("test", flag.ExitOnError)
	AddSharedFlags(testCmd)

	if len(os.Args) < 2 {
		fmt.Println("Error: Input is not enough")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "train":
		trainCmd.Parse(os.Args[2:])
	case "test":
		testCmd.Parse(os.Args[2:])
	case "help":
		fmt.Println(helpMessage)
		os.Exit(0)
	default:
		fmt.Printf("Error: Do not have command %s \n", command)
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	if target != "intent" && target != "entity" {
		fmt.Println("Error: You must choose intent or entity")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	Require(inputFP, "input file is required")
	Require(token, "application token is required")

	client, err := fptai.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	if trainCmd.Parsed() {
		if target == "intent" {
			err := TrainIntent(client, inputFP)	
			if err != nil {
				log.Fatal(err)
			}
		} else {
			TrainEntity()
		}
	}

	if testCmd.Parsed() {
		if target == "intent" {
			if err := TestIntent(client, inputFP); err != nil {
				log.Fatal(err)
			}
		} else {
			TestEntity()
		}
	}
}

func Require(f, errMsg string) {
	if f == "" {
		fmt.Println("Error: ", errMsg)
		fmt.Println(helpMessage)
		os.Exit(1)
	}
}

func AddSharedFlags(fs *flag.FlagSet) {
	fs.StringVar(&target, "t", "", "required, intent or entity")
	fs.StringVar(&inputFP, "i", "", "required, path to the input file")
	fs.StringVar(&token, "token", "", "required, your application token")
}

const helpMessage string = `
fptai is CLI tool that helps you train and test FPT.AI in terminal

Usage: fptai <command> <option>
Available commands and corresponding options:
	train
	  -t string
	    	required, type of training (intent, entity)
	  -i string
	    	required, path to your input file
	  -token string 
	  		required, your application token

	test
	  -t string
	    	required, type of training (intent, entity)
	  -i string
	    	required, path to your input file
	  -token string 
	  		required, your application token

	help
`