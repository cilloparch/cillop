package http

import (
	"fmt"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/middlewares/i18n"
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Config struct {
	Host          string
	Port          int
	Group         string
	AppName       string
	CreateHandler func(router fiber.Router) fiber.Router
	I18n          *i18np.I18n
	AcceptLangs   []string
	BodyLimit     int
}

func RunServer(cfg Config) {
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	RunServerOnAddr(addr, cfg)
}

func RunServerOnAddr(addr string, cfg Config) {
	if cfg.AppName == "" {
		cfg.AppName = "Cillop Server"
	}
	if cfg.BodyLimit == 0 {
		cfg.BodyLimit = 5 * 1024 * 1024
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: NewErrorHandler(ErrorHandlerConfig{
			DfMsgKey: "error_internal_server_error",
			I18n:     cfg.I18n,
		}),
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		CaseSensitive: true,
		AppName:       cfg.AppName,
		ServerHeader:  cfg.AppName,
		BodyLimit:     cfg.BodyLimit,
	})
	group := app.Group(fmt.Sprintf("/%v", cfg.Group))
	setGlobalMiddlewares(app, cfg)
	cfg.CreateHandler(group)
}

func setGlobalMiddlewares(router fiber.Router, cfg Config) {
	router.Use(recover.New(recover.ConfigDefault), compress.New(compress.Config{}), i18n.New(*cfg.I18n, cfg.AcceptLangs))
}