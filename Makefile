DB_DSN := "postgres://postgres:21814046098@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/app/main.go

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number

commit:
	git add .
	powershell -NoProfile -Command "& { [Console]::OutputEncoding = [System.Text.Encoding]::UTF8; [Console]::InputEncoding = [System.Text.Encoding]::UTF8; $$m = Read-Host \"Введите сообщение коммита\"; git commit -m $$m; git push }"
