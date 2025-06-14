package OverHead

import (
	"database/sql"
	"encoding/json"
)

// UpdatePallet updates an existing pallet by ID.
func UpdatePallet(db *sql.DB, pallet Pallet) error {
	skuJSON, err := json.Marshal(pallet.SKUList)
	if err != nil {
		return err
	}

	query := `
        UPDATE pallets
        SET aisle = ?, bay = ?, side = ?, sku_list = ?
        WHERE id = ?
    `
	_, err = db.Exec(query, pallet.Aisle, pallet.Bay, pallet.Side, string(skuJSON), pallet.ID)
	return err
}
