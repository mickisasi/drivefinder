package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func FormatPostalCode(postalCode string) string {
	s := strings.Split(postalCode, " ")
	if len(s) == 1 {
		return postalCode
	} else {
		return url.QueryEscape(fmt.Sprintf("%s %s", s[0], s[1]))
	}
}
