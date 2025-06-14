package OverHead

import (
	"database/sql"
	"fmt"
)

func DeletePallet(db *sql.DB, id int) error {
	query := `DELETE FROM pallets WHERE id = ?`
	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No pallet found with ID %d", id)
	}

	return nil
}
