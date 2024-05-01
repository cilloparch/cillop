package http

import (
	"fmt"

	"github.com/cilloparch/cillop/v2/i18np"
	"github.com/cilloparch/cillop/v2/log"
	"github.com/cilloparch/cillop/v2/middlewares/i18n"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Config struct {
	Host           string
	Port           int
	Group          string
	AppName        string
	CreateHandler  func(router fiber.Router) fiber.Router
	I18n           *i18np.I18n
	AcceptLangs    []string
	BodyLimit      int
	ReadBufferSize int
	Debug          bool
	Logger         log.Service
}

// RunServer runs a server with the given configuration.
// It returns an error if the server fails to start.
// If the logger is nil, it will use the default logger.
func RunServer(cfg Config) error {
	if cfg.Logger == nil || cfg.Logger == log.Service(nil) {
		cfg.Logger = log.Default(log.Config{Debug: cfg.Debug})
	}
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	return RunServerOnAddr(addr, cfg)
}

func RunServerOnAddr(addr string, cfg Config) error {
	if cfg.AppName == "" {
		cfg.AppName = "Cillop Arch"
	}
	if cfg.BodyLimit == 0 {
		cfg.BodyLimit = 10 * 1024 * 1024
	}
	if cfg.ReadBufferSize == 0 {
		cfg.ReadBufferSize = 10 * 1024 * 1024
	}
	app := fiber.New(fiber.Config{
		ErrorHandler:   NewErrorHandler(cfg.Logger, cfg.I18n),
		JSONEncoder:    json.Marshal,
		JSONDecoder:    json.Unmarshal,
		CaseSensitive:  true,
		AppName:        cfg.AppName,
		ServerHeader:   cfg.AppName,
		BodyLimit:      cfg.BodyLimit,
		ReadBufferSize: cfg.ReadBufferSize,
	})
	app.Use(i18n.New(*cfg.I18n, cfg.AcceptLangs))
	group := app.Group(fmt.Sprintf("/%v", cfg.Group))
	setGlobalMiddlewares(app, cfg)
	cfg.CreateHandler(group)
	return app.Listen(addr)
}

func setGlobalMiddlewares(router fiber.Router, cfg Config) {
	router.Use(recover.New(recover.ConfigDefault), compress.New(compress.Config{}), i18n.New(*cfg.I18n, cfg.AcceptLangs))
}
