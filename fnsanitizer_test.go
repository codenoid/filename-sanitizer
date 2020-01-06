package fnsanitizer

import "testing"

func TestSanitize(t *testing.T) {
    got := Sanitize("../../../folder/file-name.3gp")
    if got != "___folder_file-name.3gp" {
        t.Errorf(`Sanitize("../../../folder/file-name.3gp") = %v; want ___folder_file-name.3gp`, got)
    }
}
