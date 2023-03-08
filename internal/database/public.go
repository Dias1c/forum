package database

import (
	"database/sql"
	"fmt"
)

// InitDatabase - Initing database by configs.
// Sets configs, make migrations and any things. Prepare and returs database
func InitDatabase(configs *Configs) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", configs.DbPath+configs.DbConfigs)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	err = execMigration(db, configs.DbMigrationPath)
	if err != nil {
		return nil, fmt.Errorf("ExecMigration: %w", err)
	}

	return db, err
}
