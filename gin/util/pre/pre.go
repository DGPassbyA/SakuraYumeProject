package pre

import "strings"

func HTMLSpecialChars(s string) string {
	return strings.Replace(s, "<", "&lt;", -1)
}
