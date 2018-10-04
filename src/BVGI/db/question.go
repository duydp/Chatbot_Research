package db

import (
	log "github.com/Sirupsen/logrus"
)

type Question struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	IntentID int    `json:"intent_id"`
}

func InsertNewQuestion(intent, question string) error {
	id, err := GetIntentId(intent)
	if err != nil {
		log.Error("Error in InsertNewQuestion ", err.Error)
		return err
	}

	query := "INSERT INTO sample_questions(question, intent_id) VALUES(?, ?);"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("Error in InsertNewQuestion ", err.Error)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(question, id)
	if err != nil {
		log.Error("Error in InsertNewQuestion ", err.Error)
		return err
	}

	return nil
}

func GetAllQuestionsByIntentId(id int) ([]Question, error) {
	var questions []Question

	query := "SELECT * FROM sample_questions WHERE intent_id=?;"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Get All Question query: ", err)
		return questions, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Error("failed to execute Get All Question query: ", err)
		return questions, err
	}
	defer rows.Close()

	for rows.Next() {
		var u Question
		if err := rows.Scan(&u.ID, &u.Question, &u.IntentID); err != nil {
			log.Error("failed to scan question: ", err)
			return questions, err
		}
		questions = append(questions, u)
	}

	return questions, nil
}

func DeleteQuestion(id int) error {
	query := "DELETE FROM sample_questions WHERE id= ?"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete Question by ID: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete Question by ID: ", err)
		return err
	}

	return nil
}

func DeleteOldQuestion(question string) error {
	query := "DELETE FROM sample_questions WHERE question=?"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete Question by question: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(question)
	if err != nil {
		log.Error("failed to execute Delete Question by ID: ", err)
		return err
	}

	return nil
}