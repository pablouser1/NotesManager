package db

import (
	"fmt"

	"github.com/pablouser1/NotesManager/models"
)

func AddUnit(num int64, name string, subjectId int64) (models.Unit, error) {
	res, err := conn.Exec("INSERT INTO units (num, name, subject_id) VALUES (?, ?, ?)", num, name, subjectId)
	if err != nil {
		return models.Unit{}, fmt.Errorf("AddUnit: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Unit{}, err
	}

	unit := models.Unit{
		ID:        id,
		Num:       num,
		Name:      name,
		SubjectId: subjectId,
	}

	return unit, nil
}

func GetUnits(id int64) ([]models.Unit, error) {
	rows, err := conn.Query("SELECT id, num, name, subject_id FROM units WHERE subject_id=? ORDER BY num ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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
