package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dinata1312/TechGP-Project/config"
	"github.com/dinata1312/TechGP-Project/docs"
	receiptHandler "github.com/dinata1312/TechGP-Project/internal/domain/receipt/handler"
	receiptService "github.com/dinata1312/TechGP-Project/internal/domain/receipt/service"
	"github.com/dinata1312/TechGP-Project/internal/repositories/receipt_repo/receipt_pg"
	"github.com/dinata1312/TechGP-Project/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type server struct {
	cfg config.Config
	r   *gin.Engine
}

func NewServer(cfg config.Config) *server {
	return &server{
		cfg: cfg,
		r:   gin.Default(),
	}
}

func (s *server) Run() {
	db, err := postgres.NewDB(
		s.cfg.Postgres.Host,
		s.cfg.Postgres.Port,
		s.cfg.Postgres.User,
		s.cfg.Postgres.Password,
		s.cfg.Postgres.DBName,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = postgres.InitializeTable(db); err != nil {
		if err = db.Close(); err != nil {
			log.Printf("db graceful shutdown: %s\n", err.Error())
		} else {
			fmt.Printf("db graceful shutdown succeeded\n")
		}
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", s.cfg.Http.Host, s.cfg.Http.Port)

	// internalJWT := internal_jwt.NewInternalJwt()

	receiptRepo := receipt_pg.NewRepo(db)

	receiptService := receiptService.NewReceiptService(receiptRepo)

	// authMiddlerware := auth.NewAuthMiddleware(ctx, internalJWT, s.cfg, userService)

	receiptHandler := receiptHandler.NewReceiptHandler(s.r, ctx, receiptService)

	receiptHandler.MapRoutes()

	go func() {
		log.Printf("Listening on PORT: %s\n", s.cfg.Http.Port)
		if err := s.runGinServer(); err != nil {
			log.Printf("s.r.Run: %s\n", err.Error())
		}

	}()

	oscall := <-ch

	if err = db.Close(); err != nil {
		log.Printf("db graceful shutdown: %s\n", err.Error())
	} else {
		fmt.Printf("db graceful shutdown succeeded\n")
	}

	fmt.Printf("system call: %+v\n", oscall)
}
