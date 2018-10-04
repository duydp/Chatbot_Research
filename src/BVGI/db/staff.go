package db

import (
	log "github.com/Sirupsen/logrus"
)

type Staff struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	FbID 	 string `json:"fbid"`
}

func InsertStaff(s Staff) error {
	exist, err := isStaffExist(s.FbID)
	if err != nil {
		log.Errorf("failed to insert staff: %+v. Error: %s", s, err.Error())
		return err
	}

	if exist {
		return nil
	}

	query := "INSERT INTO staff(fullname, fbid) VALUES(?,?);"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Insert Staff query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Fullname, s.FbID)
	if err != nil {
		log.Error("failed to execute Insert Staff query: ", err)
		return err
	}

	return nil
}

func DeleteStaff(id int) error {
	query := "DELETE FROM staff WHERE id = ?;"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Delete Staff query: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		log.Error("failed to execute Delete Staff query: ", err)
		return err
	}

	return nil
}

func GetAllStaffs() ([]Staff, error) {
	var staffs []Staff

	query := "SELECT * FROM staff;"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Error("failed to prepare Get All Staffs query: ", err)
		return staffs, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Error("failed to execute Get All Staffs query: ", err)
		return staffs, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Staff
		if err := rows.Scan(&s.ID, &s.Fullname, &s.FbID); err != nil {
			log.Error("failed to scan staffs: ", err)
			return staffs, err
		}
		staffs = append(staffs, s)
	}

	return staffs, nil
}

func isStaffExist(fbid string) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM staff WHERE fbid=?", fbid).Scan(&count)
	if err != nil {
		log.Error("failed to count staff in database: ", err)
		return false, err
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}