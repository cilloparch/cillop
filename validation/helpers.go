package validation

import (
	"regexp"
)

func CheckEmail(email string) bool {
	matched, _ := regexp.MatchString(emailRegexp, email)
	return matched
}

func CheckPhone(phone string) bool {
	matched, _ := regexp.MatchString(PhoneWithCountryCodeRegexp, phone)
	return matched
}

func CheckUserName(userName string) (bool, string) {
	matched, _ := regexp.MatchString(userNameRegexp, userName)
	if !matched {
		return false, ""
	}
	return true, userName
}
