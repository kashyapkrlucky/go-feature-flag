package models

type FeatureFlag struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Enabled   bool   `db:"enabled" json:"enabled"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
