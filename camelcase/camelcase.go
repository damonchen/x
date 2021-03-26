package camelcase

import (
	"strings"
	"unicode"
)

func split(src string) []string {
	sp := Split(src)
	if len(sp) == 0 {
		return nil
	}

	r := make([]string, len(sp))
	j := 0
	for _, item := range sp {
		item = strings.TrimSpace(item)
		if len(item) == 0 || item == "_" {
			continue
		}
		r[j] = item
		j++
	}
	return r
}

func FirstLetterLower(s string) string {
	if len(s) == 0 {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

func FirstLetterUpper(s string) string {
	if len(s) == 0 {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func PascalCase(src string) string {
	sp := split(src)
	if len(sp) == 0 {
		return ""
	}

	for i := range sp {
		sp[i] = FirstLetterUpper(sp[i])

	}
	return strings.Join(sp, "")
}

func CamelCase(src string) string {
	sp := split(src)
	if len(sp) == 0 {
		return ""
	}

	for i := range sp {
		if i == 0 {
			sp[i] = FirstLetterLower(sp[i])
		} else {
			sp[i] = FirstLetterUpper(sp[i])
		}
	}
	return strings.Join(sp, "")
}

func SnakeCase(src string) string {
	sp := split(src)
	if len(sp) == 0 {
		return ""
	}

	for i := range sp {
		sp[i] = FirstLetterLower(sp[i])
	}
	return strings.Join(sp, "_")
}
