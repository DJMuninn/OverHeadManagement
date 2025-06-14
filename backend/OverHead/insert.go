package OverHead

import (
    "database/sql"
    "encoding/json"
)

func InsertPallet(db *sql.DB, pallet Pallet) error {
    skuJSON, err := json.Marshal(pallet.SKUList)
    if err != nil {
        return err
    }

    query := `
        INSERT INTO pallets (aisle, bay, side, sku_list)
        VALUES (?, ?, ?, ?)
    `
    _, err = db.Exec(query, pallet.Aisle, pallet.Bay, pallet.Side, string(skuJSON))
    return err
}