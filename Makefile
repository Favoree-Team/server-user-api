db_migrate:
	go run migration/migrate.go migrate_db
db_seed:
	go run migration/migrate.go migrate_seed
db_drop:
	go run migration/migrate.go drop_db