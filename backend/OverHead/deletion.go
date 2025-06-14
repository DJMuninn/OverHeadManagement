package OverHead

import (
	"database/sql"
)

// DeletePalletByID deletes a pallet by its primary ID.
func DeletePallet(db *sql.DB, id int) error {
	query := `DELETE FROM pallets WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
