package settings


import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


// загрузка переменных окружения
var _ error = godotenv.Load(".env")

// токен бота
var BotToken string = os.Getenv("BOT_TOKEN")

// JSON с доступными валютами
var PathToAvailableCurrencies string = "./settings/available_currency_list.json"
// нижний порог доступных дат для получения исторического курса
var LowDate string = "02/02/1996"

// настройки redis
var redisHost string = os.Getenv("REDIS_HOST")
var redisPort string = os.Getenv("REDIS_PORT")

var RedisAddr string = redisHost + ":" + redisPort
var RedisPassword string = os.Getenv("REDIS_PASSWORD")

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
var fatalLog *log.Logger = log.New(os.Stderr, "[FATAL]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		fatalLog.Panic(err)
	}
}


