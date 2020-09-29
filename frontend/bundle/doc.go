// +build dev

package bundle

import "net/http"

var Assets = http.Dir("frontend/html/")
