package config

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/openidea-marketplace/auth"
	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/internal/delivery/http"
	"github.com/openidea-marketplace/internal/delivery/http/middleware"
	"github.com/openidea-marketplace/internal/delivery/http/route"
	"github.com/openidea-marketplace/internal/repository/mysql"
	"github.com/openidea-marketplace/pkg/utils/hashing"
	"github.com/openidea-marketplace/user"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB  *sql.DB
	App *fiber.App
	Log domain.Logger
}

func Bootstrap(config *BootstrapConfig, viper *viper.Viper) {
	var timeCost, saltLen, memory, keyLen uint32
	var threads uint8
	timeCost = 1
	saltLen = 8
	memory = 64 * 1024
	threads = 4
	keyLen = 32
	hashing := hashing.NewArgon2idHash(timeCost, saltLen, memory, threads, keyLen)

	// setup repositories
	userRepository := mysql.NewUserRepository(config.DB)

	// setup usecases
	timeout := 5 * time.Second
	authUsecase := auth.NewAuthUsecase([]byte(viper.GetString("database.host")), config.Log)
	userUseCase := user.NewUsecase(userRepository, timeout, authUsecase, hashing)

	// setup handlers
	userHandler := http.NewUserHandler(userUseCase)

	// setup middlewares
	authMiddleware := middleware.NewAuthMiddleware(authUsecase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Setup()
}
