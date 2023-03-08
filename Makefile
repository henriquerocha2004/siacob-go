create-migration:
	migrate create -ext sql -dir internal/infra/database/mysql/migrations -seq $(name) && chmod 777 -Rf internal/infra/database/mysql/migrations
run-migration:
	migrate -path internal/infra/database/mysql/migrations -database "mysql://root:root@tcp(siacob_go_db:3306)/siacob" -verbose up
rollback-migration:
	migrate -path internal/infra/database/mysql/migrations -database "mysql://root:root@tcp(siacob_go_db:3306)/siacob" -verbose down	
