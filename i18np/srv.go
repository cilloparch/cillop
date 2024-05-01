package i18np

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func (i *I18n) translate(c *i18n.LocalizeConfig, languages ...string) string {
	localizer := i18n.NewLocalizer(i.b, languages...)
	res, err := localizer.Localize(c)
	if err != nil {
		i.logger.Error(err, messages.FailedGetMsg)
		msgId := c.MessageID
		c.MessageID = i.fallbackMsgKey
		res, err = localizer.Localize(c)
		if err != nil {
			i.logger.Error(err, messages.FailedGetFallbackMsg)
			c.MessageID = msgId
			return c.MessageID
		}
		return res
	}
	return res
}
