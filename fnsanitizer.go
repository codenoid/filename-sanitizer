package fnsanitizer

import "strings"

// simple function :p
func Sanitize(input string) string {
    input = strings.Replace(input, "..", "", -1)
    input = strings.Replace(input, "/", "_", -1)
    input = strings.Replace(input, "\\", "_", -1)
    
    return input
}
