package headers

import (
	"database/sql"
)

func InitializeTables(db *sql.DB) error {
	// Create pallets table
	palletTable := `
    CREATE TABLE IF NOT EXISTS pallets (
        id INT AUTO_INCREMENT PRIMARY KEY,
        aisle INT NOT NULL,
        bay INT NOT NULL,
        side VARCHAR(255),
        sku_list JSON
    );`

	_, err := db.Exec(palletTable)
	if err != nil {
		return err
	}

	// Create bays table
	bayTable := `
    CREATE TABLE IF NOT EXISTS bays (
        id INT AUTO_INCREMENT PRIMARY KEY,
        aisle INT NOT NULL,
        bay_number INT NOT NULL,
        shelf_num INT,
        shelf_sku JSON,
        pallet_ids JSON
    );`

	_, err = db.Exec(bayTable)
	if err != nil {
		return err
	}

	return nil
}
