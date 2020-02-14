package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/train-formula/graphcms/generated"
	"github.com/train-formula/graphcms/resolver"
	"github.com/train-formula/october"
	"github.com/willtrking/pgxload"
	"go.uber.org/zap"

	"github.com/golang-migrate/migrate/v4"
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

	pgxConf, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/graphcms?sslmode=disable",
		config.PGUsername, config.PGPassword, config.PGHost, config.PGPort))
	if err != nil {
		zap.L().Fatal("Failure to parse PGX config ", zap.Error(err))
	}

	pgxConf.MaxConns = 10
	pgxConf.MinConns = 5

	dbConn, err := pgxpool.ConnectConfig(context.Background(), pgxConf)
	if err != nil {
		zap.L().Fatal("Failure to connect to postgres ", zap.Error(err))
	}

	dbLoader, err := pgxload.NewPgxLoader(dbConn)
	if err != nil {
		zap.L().Fatal("Failure to create pgx loader ", zap.Error(err))
	}

	migrator, err := migrate.New("file://../schema/postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/graphcms?sslmode=disable&x-migrations-table=graphcms_migrations",
			config.PGUsername, config.PGPassword, config.PGHost, config.PGPort))
	if err != nil {
		zap.L().Fatal("Migrator creation error ", zap.Error(err))
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		zap.L().Fatal("Migrate error ", zap.Error(err))
	}

	ginServer := octoberServer.MustGenerateGQLGenServerServerFromEnv()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOriginFunc = func(_ string) bool {
		return true
	}

	ginServer.WithGinMiddleware(
		RegisterLoaders(dbLoader),
		cors.New(corsConfig),
	)

	ginServer.WithExecutableSchema(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(dbLoader, zap.L())}))

	octoberServer.Start(ginServer, october.DefaultGracefulShutdownSignals...)

}
