package main

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/golang-migrate/migrate/v4"
	"github.com/train-formula/graphcms"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/october"
	"go.uber.org/zap"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	octoberServer := october.MustInitServiceFromEnv()

	configurator := october.NewEnvConfigurator()

	config := Config{}

	configurator.MustDecodeEnv(&config, "")

	databaseApplicationName := strings.TrimSpace(config.PGApplication)

	if databaseApplicationName == "" {
		databaseApplicationName = "graphcms"
	}

	db := pg.Connect(&pg.Options{
		Addr:            fmt.Sprintf("%s:%s", config.PGHost, config.PGPort),
		User:            config.PGUsername,
		Password:        config.PGPassword,
		Database:        config.PGDatabase,
		ApplicationName: databaseApplicationName,
		MinIdleConns:    5,
		PoolSize:        10,
	})
	defer db.Close()

	migrator, err := migrate.New("file://./schema/postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/graphcms?sslmode=disable&x-migrations-table=graphcms_migrations", config.PGUsername, config.PGPassword, config.PGHost, config.PGPort))
	if err != nil {
		zap.L().Fatal("Migrator creation error ", zap.Error(err))
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		zap.L().Fatal("Migrate error ", zap.Error(err))
	}

	ginServer := octoberServer.MustGenerateGQLGenServerServerFromEnv()

	ginServer.WithExecutableSchema(generated.NewExecutableSchema(generated.Config{Resolvers: &graphcms.Resolver{}}))

	err = ginServer.Start()

	if err != nil {
		panic(err)
	}

}
