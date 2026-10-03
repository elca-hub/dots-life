.PHONY: migrate-create migrate-up migrate-down migrate-force gorm-generate

# 環境変数の読み込み
include .env.example
export

# マイグレーションコマンド
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir ./infra/database/migrations -seq $$name; \
	echo "Migration created: $$name"

migrate-up:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./infra/database/migrations up
	

migrate-down:
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./infra/database/migrations down

migrate-force:
	@read -p "Enter version number: " version; \
	migrate -database "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(localhost:$(MYSQL_PORT))/$(MYSQL_DATABASE)" -path ./infra/database/migrations force $$version

gorm-generate/build:
	cd ./cmd/gorm-generate && go build .

gorm-generate/run:
	cd ./cmd/gorm-generate && ./gorm-generate
