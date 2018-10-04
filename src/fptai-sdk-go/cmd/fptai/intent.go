package main 

import (
	"fmt"

	"fptai-sdk-go"

	"github.com/michlabs/nlu"
	log "github.com/Sirupsen/logrus"
)

func TrainIntent(client *fptai.Client, inputFP string) error {
	ius, err := nlu.ReadIntentsFromFile(inputFP)
	if err != nil {
		return err
	}

	labelCodeMap, err := app.LabelCodeMap()
	if err != nil {
		log.Error("failed to get label to code map: ", err)
		return err
	}

	total := len(ius)
	for i, iu := range ius {
		fmt.Printf("%d/%d\t: %s %s\n", i, total, iu.Intent, iu.Utterance)
		code, ok := labelCodeMap[iu.Intent]
		if !ok {
			intent, err := app.CreateIntent(iu.Intent, iu.Intent)
			if err != nil {
				log.Error("failed to create new intent: ", err)
				return err
			}
			labelCodeMap[intent.Label] = intent.Code
			code = intent.Code
		}

		// if err := app.AddSampleByCode(code, iu.Utterance); err != nil {
		// 	log.Error("failed to add sample by code: ", err)
		// 	return err
		// }

		_ = code
		if err := app.AddSampleByLabel(iu.Intent, iu.Utterance); err != nil {
			log.Error("failed to add sample by label: ", err)
			return err
		}
	}

	return app.Train()	
}

func TestIntent(app *fptai.Application, inputFP string) error {
	ius, err := nlu.ReadIntentsFromFile(inputFP)
	if err != nil {
		return err
	}

	total := len(ius)
	correct_counter := 0
	for i, iu := range ius {
		ir, err := app.Recognize(iu.Utterance)
		if err != nil {
			fmt.Printf("%d\tError: %s", i, err.Error())
			continue
		}

		if iu.Intent == ir.Intent {
			correct_counter++
			fmt.Printf("%d/%d:\tcorrect\t%d/%d\n", i, total, correct_counter, total)
		} else {
			fmt.Printf("%d/%d:\tincorrect\n", i, total)
		}
	}
	fmt.Printf("Result: %f\n", float64(correct_counter)/float64(total))

	return nil
}