package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"fptai-sdk-go"
	"github.com/tealeg/xlsx"
	"BVGI/db"
	"BVGI/config"
)

func main() {
	if err := config.LoadFromEnv(); err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	if err := db.Init(&config.DB); err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

	log.Println("Delete all intent")

	client := fptai.NewClient("wTVnDbO3XI9Vm4fQ5yhyX5pLUfCsRJ9s")

	intents, err := client.GetIntents()
	if err != nil {
		log.Fatal("Failed to get all intents")
	}

	for _, v := range intents {
		err := client.DeleteIntent(v.Name)
		if err != nil {
			log.Fatal("Failed to delete intent: ", v)
		}

		err = db.DeleteAllFAQ()
		if err != nil {
			log.Fatal("Failed to delete all faqs")
		}
	}

	log.Println("Deleted success")

	excelFileName := "./ftel_faq.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("Failed to open file")
	}
	cnt := 0
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			col := 0
			samples := 0
			var intent string
			var utts []string
			var questions []db.Question
			var answer string

			for _, cell := range row.Cells {
				col = col + 1
				text, _ := cell.String()
				text = strings.Replace(text, "\n"," ", -1)
				if col == 2 {
					intent = text
					_, err := client.CreateIntent(text, text)
					if err != nil {
						log.Fatal("Failed to create Intent: ", intent)
					}
				}
				if col == 3 {
					answer = text
				}
				if col == 4 {
					arr := strings.SplitN(text, ".", -1)
					for _, v:= range arr {
						var re = regexp.MustCompile(` [0-9]+`)
						content := re.ReplaceAllString(v, ``)
						content = strings.TrimSpace(content)
						if content != "" && content !="1" {
							samples ++
							utts = append(utts, content)

							question := db.Question{
								ID: 0,
								Question:content,
								IntentID: 0,
							}

							questions = append(questions, question)

							fmt.Printf("#%d: %s|\n",samples, content)
						}
					}
					err := client.CreateUtterances(intent, utts)
					if err != nil {
						log.Println("Error: ", err.Error())
						log.Fatalf("Failed to create utts: %s %s\n", intent, utts)

					}
					faq := db.FAQ{
						ID:0,
						IntentName: intent,
						Answer:answer,
						Questions:questions,
					}

					err = db.InsertFAQ(faq)

					if err != nil {
						log.Fatal("Failed to insert faq: ", faq)
					}
				}
			}
			cnt = cnt + 1
			fmt.Println("\n\n")
		}
	}

	fmt.Println("Done.")
}