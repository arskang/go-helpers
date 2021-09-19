package goHelpers

import "regexp"

func IsRFC(text string) bool {
	r, _ := regexp.Compile(`^[a-zñA-ZÑ\x26]{3,6}[0-9]{2}(0[1-9]|1[0-2])(0[1-9]|[12][0-9]|3[01])[a-zA-Z0-9]{3}$`)
	return r.Match([]byte(text))
}

func IsEmail(text string) bool {
	r, _ := regexp.Compile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	return r.Match([]byte(text))
}

func IsMultiEmail(text string) bool {
	r, _ := regexp.Compile(`^((([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,})),?)+$`)
	return r.Match([]byte(text))
}

func IsPhone(text string) bool {
	r, _ := regexp.Compile(`^[0-9]{10,15}$`)
	return r.Match([]byte(text))
}

func IsZipCode(text string) bool {
	r, _ := regexp.Compile(`^(?:\d{5}(?:[-]\d{4})?|\d{9})$`)
	return r.Match(([]byte(text)))
}

func IsUrl(text string) bool {
	r, _ := regexp.Compile(`^[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
	return r.Match(([]byte(text)))
}
