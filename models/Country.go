package models

type Country struct {
	CountryID   int    `db:"country_id"`
	CountryName string `db:"country_name"`
}
