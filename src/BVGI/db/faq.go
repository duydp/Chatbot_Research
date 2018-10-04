package db

import (
	log "github.com/Sirupsen/logrus"
)

type FAQ struct {
	ID         int        `json:"id"`
	IntentName string     `json:"intent"`
	Answer     string     `json:"answer"`
	Questions  []Question `json:"questions"`
}


func InsertFAQ(faq FAQ) error {
	query := "INSERT INTO intent_answer(intent, answer) VALUES(?,?);"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Insert FAQ query: ", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(faq.IntentName, faq.Answer)
	if err != nil {
		log.Error("failed to execute Insert FAQ query: ", err)
		return err
	}

	intent_id, err := res.LastInsertId()
	if err != nil {
		log.Error("failed to get last insert ID: ", err)
		return err
	}

	for _, q := range faq.Questions {
		query := "INSERT INTO sample_questions(question, intent_id) VALUES(?, ?);"

		stmt, err := DB.Prepare(query)
		if err != nil {
			log.Error("failed to prepare Insert Question query: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(q.Question, intent_id)
		if err != nil {
			log.Error("failed to execute Insert Question query: ", err)
			return err
		}
	}

	return nil
}

func UpdateFAQ(faq FAQ) error {
	query := "UPDATE intent_answer SET intent=?, answer=? WHERE id=? ;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Update FAQ query: ", err)
		return err
	}
	defer stmt.Close()

	// TODO: Use transaction
	_, err = stmt.Exec(faq.IntentName, faq.Answer, faq.ID)
	if err != nil {
		log.Error("failed to execute Update FAQ query: ", err)
		return err
	}

	if err := DeleteAllSamples(faq.ID); err != nil {
		log.Error("failed to delete all samples ", err)
		return err
	}

	for _, q := range faq.Questions {
		query := "INSERT INTO sample_questions(question, intent_id) VALUES(?, ?);"

		stmt, err := DB.Prepare(query)
		if err != nil {
			log.Error("failed to prepare Insert Question query: ", err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(q.Question, faq.ID)
		if err != nil {
			log.Error("failed to execute Insert Question query: ", err)
			return err
		}
	}

	return nil
}

func DeleteAllSamples(id int) error {
	query := "DELETE FROM sample_questions WHERE intent_id=?;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete Question query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete Question query: ", err)
		return err
	}
	return nil
}

// TODO: Use transaction
func DeleteFAQ(id int) error {
	query := "DELETE FROM intent_answer WHERE id = ?;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete FAQ query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete FAQ query: ", err)
		return err
	}

	query = "DELETE FROM sample_questions WHERE intent_id=?;"

	stmt, err = DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete Question query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete Question query: ", err)
		return err
	}

	return nil
}

func GetAllFAQs() ([]FAQ, error) {
	var faqs []FAQ

	query := "SELECT * FROM intent_answer;"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Get All FAQs query: ", err)
		return faqs, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error("failed to execute Get All FAQs query: ", err)
		return faqs, err
	}
	defer rows.Close()

	for rows.Next() {
		var u FAQ
		err := rows.Scan(&u.ID, &u.IntentName, &u.Answer)
		if err != nil {
			log.Error("failed to scan FAQ: ", err)
			return faqs, err
		}

		u.Questions, err = GetAllQuestionsByIntentId(u.ID)
		if err != nil {
			log.Error("failed to Get All Questions by Intent ID: ", err)
			return faqs, err
		}
		faqs = append(faqs, u)
	}

	return faqs, nil
}

func GetAnswerFor(intent string) string {
	if intent == "" {
		return ""
	}

	faq, err := GetFAQByIntent(intent)
	if err != nil {
		return ""
	}

	return faq.Answer
}

func GetFAQByIntent(intent string) (FAQ, error) {
	var faq FAQ
	err := DB.QueryRow("SELECT * FROM `intent_answer` WHERE intent=?", intent).Scan(&faq.ID, &faq.IntentName, &faq.Answer)
	if err != nil {
		return faq, err
	}
	faq.Questions, err = GetAllQuestionsByIntentId(faq.ID)
	if err != nil {
		log.Error("failed to get all question by intent ID: ", err)
		return faq, err
	}

	return faq, nil
}

func GetFAQByID(id int) (FAQ, error) {
	var faq FAQ
	err := DB.QueryRow("SELECT * FROM `intent_answer` WHERE id=?", id).Scan(&faq.ID, &faq.IntentName, &faq.Answer)

	if err != nil {
		log.Error("failed to query faq by ID: ", err)
		return faq, err
	}
	faq.Questions, err = GetAllQuestionsByIntentId(faq.ID)
	if err != nil {
		log.Error("failed to get all question by intent ID: ", err)
		return faq, err
	}

	return faq, nil
}
