info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go build -o ./GoCurrencyCourseBot ./main.go

prod:
	@echo "Running migrations..."
	/root/GoCurrencyCourseBot migrate
	@echo "Running main app..."
	/root/GoCurrencyCourseBot >> $(info_log) 2>> $(error_log)

