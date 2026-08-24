# Makefile untuk mempersingkat pengetikan command sql migration
include .env
export

DSN=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

migrate-create:
	docker compose run --rm migrate create -ext sql -dir /migrations -seq $(name)
# example command : sudo make migrate-create name=create_users
# $(var) = menerima parameter dari luar

migrate-up:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" up

migrate-down:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" down 1

migrate-force:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" force $(v)

migrate-version:
	docker compose run --rm migrate -path=/migrations -database="$(DSN)" version