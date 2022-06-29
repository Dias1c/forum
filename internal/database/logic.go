package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
)

func execMigration(db *sql.DB, mgPath string) error {
	f, err := os.OpenFile(mgPath, os.O_RDONLY, 0755)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %w", err)
	}
	defer f.Close()

	migrationData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(migrationData))
	if err != nil {
		return fmt.Errorf("db.Exec: %w", err)
	}
	return nil
}
