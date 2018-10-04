package db

import (
	log "github.com/Sirupsen/logrus"
)

type History struct {
	ID       	int    `json:"id"`
	Question 	string `json:"question"`
	Answer 	 	string `json:"answer"`
	AskTime 	string `json:"asktime"`
}

func InsertHistory(h History) error {
	query := "INSERT INTO history SET question=?, answer=?, ask_time=NOW();"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Insert History query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(h.Question, h.Answer)
	if err != nil {
		log.Error("failed to execute Insert History query: ", err)
		return err
	}

	return nil
}

func UpdateHistoryAnswer(question string, answer string) error {
	query := "UPDATE history SET answer=? WHERE question=? ;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare UpdateHistoryAnswer: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(answer, question)
	if err != nil {
		log.Error("failed to execute UpdateHistoryAnswer query: ", err)
		return err
	}

	return nil
}

func DeleteHistory(id int) error {
	query := "DELETE FROM history WHERE id = ?;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete History query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete History query: ", err)
		return err
	}

	return nil
}

func GetAllHistories() ([]History, error) {
	var histories []History

	query := "SELECT * FROM history ORDER BY ask_time DESC;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Get All History query: ", err)
		return histories, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error("failed to execute Get All History query: ", err)
		return histories, err
	}
	defer rows.Close()

	for rows.Next() {
		var h History
		if err := rows.Scan(&h.ID, &h.Question, &h.Answer, &h.AskTime); err != nil {
			log.Error("failed to scan history row: ", err)
			return histories, err
		}
		histories = append(histories, h)
	}

	return histories, nil
}

func GetHistoryRows() (int, error) {
	query := "SELECT COUNT(*) FROM history;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare GetHistoryRows query: ", err)
		return 0, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error("failed to execute GetHistoryRows query: ", err)
		return 0, err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Error("failed to scan GetHistoryRows row: ", err)
			return 0, err
		}
	}

	return count, nil
}

func GetHistoryForPagination(limit, offset int) ([]History, error) {
	var histories []History

	query := "SELECT * FROM history ORDER BY ask_time DESC LIMIT ? OFFSET ?;"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare GetHistoryForPagination query: ", err)
		return histories, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		log.Error("failed to execute GetHistoryForPagination query: ", err)
		return histories, err
	}
	defer rows.Close()

	for rows.Next() {
		var h History
		if err := rows.Scan(&h.ID, &h.Question, &h.Answer, &h.AskTime); err != nil {
			log.Error("failed to scan GetHistoryForPagination row: ", err)
			return histories, err
		}
		histories = append(histories, h)
	}

	return histories, nil
}