package database

type Configs struct {
	DbPath          string `cenv:"db_path"`
	DbConfigs       string `cenv:"db_configs"`
	DbMigrationPath string `cenv:"db_migration_path"`
}
