package utility

import (
	"net/http"
	"regexp"
	"strings"
)

func GetRealIP(r *http.Request) string {
	varRealIP := r.Header.Get("X-Real-Ip")

	if len(varRealIP) > 0 {
		return varRealIP
	}

	valForwardedIP := r.Header.Get("X-Forwarded-For")

	if len(valForwardedIP) > 0 {
		strIPs := strings.Split(valForwardedIP, ",")

		if len(strIPs) > 0 {
			retReg, err := regexp.MatchString("((?:(?:25[0-5]|2[0-4]\\d|[01]?\\d?\\d)\\.){3}(?:25[0-5]|2[0-4]\\d|[01]?\\d?\\d))", strIPs[0])

			if err == nil && retReg == true {
				return strIPs[0]
			}
		}
	}

	return "127.0.0.1"
}
