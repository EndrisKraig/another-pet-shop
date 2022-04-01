if [ .env ]
then
  export $(cat .env | xargs)
fi
migrate -path scripts/migration -database ${DB_URL} -verbose up

#migrate create -ext sql -dir scripts/migration -seq create_users_table