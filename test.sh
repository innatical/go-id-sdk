cd ./test

set -a
[ -f .env ] && . ./.env

go run test.go