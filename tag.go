package html

import (
	"regexp"
)

func StripTags(s string) string {
	beginTagRe := regexp.MustCompile("<[^>]+>")
	endTagRe := regexp.MustCompile("</[^>]+>")

	str := beginTagRe.ReplaceAllString(s, "")
	str = endTagRe.ReplaceAllString(str, "")

	return str
}
