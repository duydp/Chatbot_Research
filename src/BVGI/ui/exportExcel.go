package ui

import (
	//"fmt"
	"os"
	"time"

	//"github.com/Luxurioust/excelize"
	"github.com/Sirupsen/logrus"
	//"BVGI/db"
)

func DeleteFile(path string, delay int) {
	if isExists(path){
		time.Sleep(time.Duration(delay) * time.Second)
		logrus.Info("remove file: ", path)
		err := os.Remove(path)
		if err != nil {
			logrus.Error("failed to delete file: ", err.Error())
		}
		logrus.Info("File deleted.")
	}else{
		logrus.Info("File not found")
	}

}

func isExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Export(filePath string) error{
	//faqs, err := db.GetAllFAQs()

	//if err != nil {
	//	logrus.Error("Cannot get all faqs in Export", err.Error())
	//	return err
	//}

	//xlsx := excelize.CreateFile()

	/*
	xlsx.NewSheet(1, "Sheet1")
	for i1, v1 := range faqs {
		stt := fmt.Sprintf("A%d", i1+1)
		intentCell := fmt.Sprintf("B%d", i1+1)
		answerCell := fmt.Sprintf("C%d", i1+1)
		questionsCell := fmt.Sprintf("D%d", i1+1)
		xlsx.SetCellValue("Sheet1",stt, fmt.Sprintf("%d", i1+1))
		xlsx.SetCellValue("Sheet1",intentCell, v1.IntentName)
		xlsx.SetCellValue("Sheet1",answerCell, v1.Answer)

		questions := ""
		for i2, v2 := range  v1.Questions {
			questions = questions + "\r\n" +fmt.Sprintf("%d. %s", i2+1, v2.Question)
		}
		xlsx.SetCellValue("Sheet1",questionsCell, questions)
	}

	xlsx.SetActiveSheet(1)

	DeleteFile(filePath, 0)
	err = xlsx.WriteTo(filePath)
	if err != nil {
		logrus.Error("Error in exportExcel", err.Error())
		return err
	}

	logrus.Info("Done.")
	*/
	return nil
}