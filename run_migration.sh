if [ ! -f .env ]
then
  export $(cat .env | xargs)
fi
migrate -path scripts/migration -database ${DB_URL} -verbose up