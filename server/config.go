package main

type Config struct {
	PGHost string `october:"PG_HOST"`
	PGUsername string `october:"PG_USERNAME"`
	PGDatabase string `october:"PG_DATABASE"`
	PGPort string `october:"PG_PORT"`
	PGPassword string `october:"PG_PASSWORD"`
	PGApplication string `october:"PG_APPLICATION"`
}