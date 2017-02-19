Postgres 起動
$ createdb chitchat
$ psql -f data/setup.sql -d chitchat
$ go build
$ ./chitchat
