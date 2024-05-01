package i18np

type msg struct {
	FailedGetFallbackMsg string
	FailedGetMsg         string
	FailedParseLocale    string
}

var messages = msg{
	FailedGetFallbackMsg: "failed to get fallback message",
	FailedGetMsg:         "failed to get message",
	FailedParseLocale:    "failed to parse fallback language",
}
