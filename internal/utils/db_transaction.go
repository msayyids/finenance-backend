package utils

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func CommitOrRollback(tx *sqlx.Tx, err *error) {
	if r := recover(); r != nil {
		// Jika panic terjadi rollback
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("rollback failed after panic: %v\n", rollbackErr)
		}
		log.Printf("recovered from panic: %v\n", r)
		// boleh dihapus kalau tidak mau propagate panic
	} else if *err != nil {
		// Jika ada error biasa  rollback
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("rollback failed after error: %v\n", rollbackErr)
		}
	} else {
		// Kalau aman →commit
		if commitErr := tx.Commit(); commitErr != nil {
			log.Printf("commit failed: %v\n", commitErr)
		}
	}
}
