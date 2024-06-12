info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go build -o ./go_app ./main.go

prod:
	@echo "Running migrations..."
	/root/go_app migrate
	@echo "Running main app..."
	/root/go_app >> $(info_log) 2>> $(error_log)

