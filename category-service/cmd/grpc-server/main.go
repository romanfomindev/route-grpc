package main

import (
	"flag"

	config "github.com/romanfomindev/route-grpc/tree/master/category-service/internal/config"
	"github.com/romanfomindev/route-grpc/tree/master/category-service/internal/server"
	"github.com/romanfomindev/route-grpc/tree/master/category-service/internal/service/category"
	category_repository "github.com/romanfomindev/route-grpc/tree/master/category-service/internal/service/category/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	categoryRepo := category_repository.New()
	categoryService := category.New(categoryRepo)

	if err := server.NewGrpcServer(
		categoryService,
	).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
