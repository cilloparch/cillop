package i18np

import "github.com/nicksnyder/go-i18n/v2/i18n"

func (i *I18n) translate(c *i18n.LocalizeConfig, languages ...string) string {
	localizer := i18n.NewLocalizer(i.b, languages...)
	return localizer.MustLocalize(c)
}
