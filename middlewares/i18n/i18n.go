package i18n

import (
	"strings"

	"github.com/cilloparch/cillop/v2/i18np"
	"github.com/gofiber/fiber/v2"
)

// AcceptedLanguages is the list of accepted languages.
var AcceptedLanguages = []string{}

// GetLanguagesInContext returns the language and the accept-language header from the context.
func GetLanguagesInContext(i i18np.I18n, c *fiber.Ctx) (string, string) {
	l := c.Query("lang")
	a := c.Get("Accept-Language", i.Fallback())
	if l == "" {
		l = a
	}
	list := strings.Split(l, ";")
	alternative := ""
	locales := findLocales(list)
	for _, v := range AcceptedLanguages {
		if locales[v] {
			return v, a
		}
	}
	if len(list) > 1 {
		alternative = list[1]
	}
	if alternative != "" {
		return alternative, a
	}
	return l, a
}

// New returns a new i18n middleware.
func New(i i18np.I18n, acceptLangs []string) fiber.Handler {
	AcceptedLanguages = acceptLangs
	return func(c *fiber.Ctx) error {
		l, a := GetLanguagesInContext(i, c)
		c.Locals("lang", l)
		c.Locals("accept-language", a)
		return c.Next()
	}
}

// ParseLocale parses the locale from the context.
func ParseLocale(ctx *fiber.Ctx) string {
	return ctx.Locals("lang").(string)
}

// ParseLocales parses the locales from the context.
func ParseLocales(ctx *fiber.Ctx) (string, string) {
	return ctx.Locals("lang").(string), ctx.Locals("accept-language").(string)
}

func findLocales(list []string) map[string]bool {
	locales := make(map[string]bool)
	for _, li := range list {
		lineItems := strings.Split(li, ",")
		for _, word := range lineItems {
			if word == "en" || word == "tr" {
				locales[word] = true
			}
			if len(word) == 2 && word[1] == '-' {
				locales[strings.ToLower(word)] = true
			}
			if len(word) == 5 && word[2] == '-' {
				double := strings.Split(word, "-")
				locales[double[0]] = true
			}
		}
	}
	return locales
}
