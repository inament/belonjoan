package main

import (
	"database/sql"
)

// Belonjoan -
type Belonjoan struct {
	ItemID int     `json:"item_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

// Belonjoans -
type Belonjoans []Belonjoan

// BelonjoanStore -
type BelonjoanStore struct {
	db *sql.DB
}

// NewBelonjoanStore -
func NewBelonjoanStore(db *sql.DB) *BelonjoanStore {
	return &BelonjoanStore{db}
}

// Add -
func (br BelonjoanStore) Add(belonjoan Belonjoan) error {
	stmt, err := br.db.Prepare("INSERT INTO belonjoan(name, price) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(belonjoan.Name, belonjoan.Price)
	if err != nil {
		if err != nil {
			return err
		}
	}

	return nil
}

// FindAll -
func (br BelonjoanStore) FindAll() (*Belonjoans, error) {
	rows, err := br.db.Query("SELECT * FROM belonjoan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	belonjoans := make(Belonjoans, 0)
	for rows.Next() {
		var rowBelonjoan Belonjoan
		rows.Scan(&rowBelonjoan.ItemID, &rowBelonjoan.Name, &rowBelonjoan.Price)
		belonjoans = append(belonjoans, rowBelonjoan)
	}

	return &belonjoans, nil
}

// FindOne -
func (br BelonjoanStore) FindOne(itemID int) (*Belonjoan, error) {
	stmt, err := br.db.Prepare("SELECT * FROM belonjoan WHERE item_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var belonjoan Belonjoan
	err = stmt.QueryRow(itemID).Scan(&belonjoan.ItemID, &belonjoan.Name, &belonjoan.Price)
	if err != nil {
		return nil, err
	}

	return &belonjoan, nil
}

// UpdateItem -
func (br BelonjoanStore) UpdateItem(itemID int, belonjoan Belonjoan) error {
	stmt, err := br.db.Prepare("UPDATE belonjoan SET name = ?, price = ? WHERE item_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(belonjoan.Name, belonjoan.Price, itemID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteItem -
func (br BelonjoanStore) DeleteItem(itemID int) error {
	stmt, err := br.db.Prepare("DELETE FROM belonjoan WHERE item_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(itemID)
	if err != nil {
		return err
	}

	return nil
}
