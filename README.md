## Запуск проекта:

Создайте файл .env в корневом каталоге и добавьте следующие значения:

### `

APP_ENV="dev"

MONGO_URI="mongodb://127.0.0.1:27017"
MONGO_USER="<youusername>"
MONGO_PASSWORD="<yourpassword>"

REDIS_PASSWORD=""

HTTP_HOST="http://localhost"

JWT_KEY="<secretkey>"

BCRYPT_MIN_COST=6
BCRYPT_DEFAULT_COST=14
BCRYPT_MAX_COST=30
`

Запуск докер командой `docker-compose up`
запуск dev версии `air -d`
