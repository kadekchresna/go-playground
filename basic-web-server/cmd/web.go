package cmd

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"basic.web/config"
	driver_db "basic.web/driver/db"
	"basic.web/helper/logger"
	echopprof "basic.web/helper/pprof"
	orderRepo "basic.web/internal/repository/orders"
	settlementsRepo "basic.web/internal/repository/settlements"
	settlementsUc "basic.web/internal/usecase/settlements"
	v1 "basic.web/internal/web/v1"
	settlementsHandler "basic.web/internal/web/v1/settlements"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

const (
	STAGING     = `stg`
	PRODUCTIOON = `prd`
)

func init() {
	if os.Getenv("APP_ENV") != PRODUCTIOON {

		// init invoke env before everything
		cobra.OnInitialize(initConfig)

	}

	// adding command invokable
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "web",
	Short: "Running Web Service",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func initConfig() {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("error load ENV, %s", err.Error()))
	}
}

func run() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := config.InitConfig()

	db := driver_db.InitDB(config)

	//repo
	settlementsRepos := settlementsRepo.NewSettlementRepo(db)
	orderRepos := orderRepo.NewOrderRepo(db)

	//usecase
	settlementsUsecase := settlementsUc.NewSettlementsUsecase(settlementsRepos, orderRepos)

	//handler
	settlementsHandler := settlementsHandler.NewSettlementsHandler(settlementsUsecase)

	v1.InitAPI(e, settlementsHandler)

	e.Any("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Register all the standard library debug endpoints.
	// curl --output pprofile "http://127.0.0.1:50218/debug/pprof/profile?seconds=100"
	// go tool pprof -http localhost:5432 pprofile
	echopprof.RegisterPprofRoutes(e)
	s := http.Server{
		Addr:    fmt.Sprintf(":%d", config.AppPort),
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	logger.Log().Info(fmt.Sprintf("%s service started...", config.AppName))
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	logger.Log().Info(fmt.Sprintf("%s service finished", config.AppName))
}
