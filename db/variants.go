package db

import (
	"github.com/pablouser1/NotesManager/models"
)

func GetVariants() ([]models.Variant, error) {
	rows, err := conn.Query("SELECT id, name, slug FROM variants ORDER BY slug ASC")
	if err != nil {
		return nil, err
	}

	var variants []models.Variant

	// Default option
	variants = append(variants, models.Variant{
		ID:   -1,
		Name: "Default",
		Slug: "default",
	})

	for rows.Next() {
		var subject models.Subject
		if err := rows.Scan(&subject.ID, &subject.Name, &subject.Slug); err != nil {
			return variants, err
		}
		variants = append(variants, subject)
	}
	if err = rows.Err(); err != nil {
		return variants, err
	}

	return variants, nil
}

func GetVariantByName(name string) (models.Variant, error) {
	stmt, err := conn.Prepare("SELECT id, name, slug FROM variants WHERE name=?")
	if err != nil {
		return models.Variant{}, err
	}

	var variant models.Variant

	err = stmt.QueryRow(name).Scan(&variant.ID, &variant.Name, &variant.Slug)

	if err != nil {
		return variant, err
	}

	return variant, nil
}
