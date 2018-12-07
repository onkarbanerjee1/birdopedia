package birds

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// All returns all birds
func All(db *sql.DB) ([]*Bird, error) {
	rows, err := db.Query(`SELECT common_name as commonName, 
	scientific_name as scientificName, 
	pic_url as pictureURL, 
	habitat,
	endangered, 
	postedby 
	FROM birds;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	birds := []*Bird{}
	for rows.Next() {
		bird := Bird{}
		if err = rows.Scan(&bird.CommonName,
			&bird.ScientificName,
			&bird.PictureURL,
			pq.Array(&bird.Habitat),
			&bird.Endangered,
			&bird.PostedBy); err != nil {
			return nil, err
		}
		birds = append(birds, &bird)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return birds, nil
}

// ByName returns all birds with the common name
func ByName(db *sql.DB, name string) (*Bird, error) {
	rows := db.QueryRow(`SELECT common_name as commonName, 
	scientific_name as scientificName, 
	pic_url as pictureURL, 
	habitat,
	endangered, 
	postedby 
	FROM birds where UPPER(common_name) = UPPER($1);`, name)

	bird := Bird{}

	if err := rows.Scan(&bird.CommonName,
		&bird.ScientificName,
		&bird.PictureURL,
		pq.Array(&bird.Habitat),
		&bird.Endangered,
		&bird.PostedBy); err != nil {
		return nil, err
	}

	return &bird, nil
}

// ByID returns a bird for the given ID
func ByID(db *sql.DB, id int) (*Bird, error) {
	row := db.QueryRow(`SELECT common_name as commonName, 
	scientific_name as scientificName, 
	pic_url as pictureURL, 
	habitat,
	endangered, 
	postedby 
	FROM birds where id = $1;`, id)

	bird := Bird{}
	err := row.Scan(&bird.CommonName, &bird.ScientificName, &bird.PictureURL, pq.Array(&bird.Habitat),
		&bird.Endangered, &bird.PostedBy)
	if err != nil {
		return nil, err
	}

	return &bird, nil
}

func getID(db *sql.DB, name string) (int, error) {
	row := db.QueryRow(`SELECT "ID" FROM birds where common_name = $1;`, name)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Insert adds a new Bird to the db
func Insert(db *sql.DB, bird *Bird) (int64, error) {
	if db == nil {
		return 0, fmt.Errorf("db not passed")
	}

	res, err := db.Exec(`INSERT INTO birds(
		common_name, scientific_name, pic_url, habitat, endangered, postedby) 
		VALUES ($1, $2, $3, $4, $5 ,$6);`,
		bird.CommonName,
		bird.ScientificName,
		bird.PictureURL,
		pq.Array(bird.Habitat),
		bird.Endangered,
		bird.PostedBy)

	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// Update updates a record
func Update(db *sql.DB, id int, bird *Bird) error {
	if db == nil {
		return fmt.Errorf("db not passed")
	}

	sql := `UPDATE birds
	SET common_name=$1, scientific_name=$2, pic_url=$3, habitat=$4,
		endangered=$5, postedby=$6  WHERE "ID"=$7;`
	res, err := db.Exec(sql, bird.CommonName, bird.ScientificName, bird.PictureURL, pq.Array(bird.Habitat), bird.Endangered, bird.PostedBy, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

// Delete deletes a record
func Delete(db *sql.DB, id int) error {
	if db == nil {
		return fmt.Errorf("db not passed")
	}

	sql := `DELETE FROM birds WHERE "ID"=$1;`
	res, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}
