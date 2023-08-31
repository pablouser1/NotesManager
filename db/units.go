package db

import "github.com/pablouser1/NotesManager/models"

func GetUnits(id int64) ([]models.Unit, error) {
	rows, err := conn.Query("SELECT id, num, name, subject_id FROM units WHERE subject_id=?", id)
	if err != nil {
		return nil, err
	}

	var units []models.Unit

	for rows.Next() {
		var unit models.Unit
		if err := rows.Scan(&unit.ID, &unit.Num, &unit.Name, &unit.SubjectId); err != nil {
			return units, err
		}
		units = append(units, unit)
	}
	if err = rows.Err(); err != nil {
		return units, err
	}
	return units, nil
}
