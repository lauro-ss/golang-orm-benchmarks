// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID           int32
	Isbn         string
	Title        string
	Author       string
	Genre        string
	Quantity     int32
	PublicizedAt pgtype.Timestamp
}
