package action

import "strings"

func stringFormat(format string, args ...string) string {
    r := strings.NewReplacer(args...)
    return r.Replace(format)
}
