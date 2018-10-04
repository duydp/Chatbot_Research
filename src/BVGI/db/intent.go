package db

import (
	log "github.com/Sirupsen/logrus"
)

func GetIntentId(intent string) (int, error) {
	var id int
	err := DB.QueryRow("SELECT id FROM `intent_answer` WHERE intent=?", intent).Scan(&id)
	if err != nil {
		log.Error("failed to count intents in database: ", err)
		return 0, err
	}

	return id, nil
}

func IsIntentExist(intent string) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM `intent_answer` WHERE intent=?", intent).Scan(&count)
	if err != nil {
		log.Error("failed to count intents in database: ", err)
		return false, err
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetAllIntents() ([]string, error) {
	var intents []string

	query := "SELECT intent FROM intent_answer;"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Get All Intents query: ", err)
		return intents, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error("failed to query Get All Intnets query: ", err)
		return intents, err
	}
	defer rows.Close()

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Error("failed to scan intent: ", err)
			return intents, err
		}
		intents = append(intents, s)
	}

	return intents, nil
}