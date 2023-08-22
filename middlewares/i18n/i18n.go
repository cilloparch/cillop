package i18n

import (
	"strings"

	"github.com/cilloparch/cillop/i18np"
	"github.com/gofiber/fiber/v2"
)

var AcceptedLanguages = []string{}

func GetLanguagesInContext(i i18np.I18n, c *fiber.Ctx) (string, string) {
	l := c.Query("lang")
	a := c.Get("Accept-Language", i.Fallback)
	if l == "" {
		l = a
	}
	list := strings.Split(l, ",")
	alternative := ""

	for _, la := range list {
		for _, v := range AcceptedLanguages {
			if strings.Contains(la, v) {
				return v, a
			}
			if strings.Contains(la, v[:2]) {
				alternative = v
			}
		}
	}

	if alternative != "" {
		return alternative, a
	}
	return l, a
}

func New(i i18np.I18n, acceptLangs []string) fiber.Handler {
	AcceptedLanguages = acceptLangs
	return func(c *fiber.Ctx) error {
		l, a := GetLanguagesInContext(i, c)
		c.Locals("lang", l)
		c.Locals("accept-language", a)
		return c.Next()
	}
}
