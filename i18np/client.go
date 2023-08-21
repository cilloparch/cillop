package i18np

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// I18n is base struct for i18n
// b is i18n bundle
// Fallback is default language
type I18n struct {
	// b is i18n bundle
	b *i18n.Bundle

	// Fallback is default language
	Fallback string
}

// New is constructor for I18n
// fallback is default language
// return I18n
func New(fallback string) *I18n {
	b := i18n.NewBundle(language.English)
	return &I18n{b: b, Fallback: fallback}
}

// Load is load i18n file
// ld is directory path
// languages is language list
// example: i18n.Load("./i18n", "en", "ja")
func (i *I18n) Load(ld string, languages ...string) {
	i.b.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	for _, lang := range languages {
		i.b.MustLoadMessageFile(ld + "/" + lang + ".toml")
	}
}

// AddMessages is add i18n message
// lang is language
// messages is i18n message
// example: i18n.AddMessages("en", &i18n.Message{ID: "hello", Other: "Hello!"})
func (i *I18n) AddMessages(lang string, messages ...*i18n.Message) error {
	return i.b.AddMessages(language.MustParse(lang), messages...)
}

// Translate is translate i18n message
// key is i18n key
// languages is language list
// example: i18n.Translate("hello", "en")
// example: i18n.Translate("hello", "en", "ja")
// example: i18n.Translate("hello", "ja", "en")
func (i *I18n) Translate(key string, languages ...string) string {
	return i.translate(&i18n.LocalizeConfig{
		MessageID: key,
	}, languages...)
}

// TranslateWithParams is translate i18n message with params
// key is i18n key
// params is i18n params
// languages is language list
// example: i18n.TranslateWithParams("hello", i18n.P{"Name": "John"}, "en")
func (i *I18n) TranslateWithParams(key string, params interface{}, languages ...string) string {
	return i.translate(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: params,
	}, languages...)
}

// TranslateFromError is translate i18n message from I18nError
// err is I18nError
//
//	languages is language list
//
// example: i18n.TranslateFromError(err, "en")
// example: i18n.TranslateFromError(err, "en", "ja")
func (i *I18n) TranslateFromError(err Error, languages ...string) string {
	return i.translate(&i18n.LocalizeConfig{
		MessageID:    err.Key,
		TemplateData: err.Params,
	}, languages...)
}

// TranslateFromErrorDetail is translate i18n message from I18nError
// err is I18nError
// languages is language list
// return string, interface{}
// example: i18n.TranslateFromErrorDetail(err, "en")
// example: i18n.TranslateFromErrorDetail(err, "en", "ja")
// example: i18n.TranslateFromErrorDetail(err, "ja", "en")
func (i *I18n) TranslateFromErrorDetail(err Error, languages ...string) (string, interface{}) {
	if !err.IsDetails() {
		return i.TranslateFromError(err, languages...), nil
	}
	return i.translate(&i18n.LocalizeConfig{
		MessageID:    err.Key,
		TemplateData: err.Params,
	}, languages...), err.Details
}