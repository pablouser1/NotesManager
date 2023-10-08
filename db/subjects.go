package db

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/pablouser1/NotesManager/models"
)

func AddSubject(name string) (models.Subject, error) {
	slugName := slug.Make(name)
	res, err := conn.Exec("INSERT INTO subjects (name, slug) VALUES (?, ?)", name, slugName)
	if err != nil {
		return models.Subject{}, fmt.Errorf("AddSubject: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Subject{}, err
	}

	subject := models.Subject{
		ID:   id,
		Name: name,
		Slug: slugName,
	}

	return subject, nil
}

func GetSubjects() ([]models.Subject, error) {
	rows, err := conn.Query("SELECT id, name, slug FROM subjects ORDER BY slug ASC")
	if err != nil {
		return nil, err
	}

	var subjects []models.Subject

	for rows.Next() {
		var subject models.Subject
		if err := rows.Scan(&subject.ID, &subject.Name, &subject.Slug); err != nil {
			return subjects, err
		}
		subjects = append(subjects, subject)
	}
	if err = rows.Err(); err != nil {
		return subjects, err
	}
	return subjects, nil
}
