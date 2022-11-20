package handler

import (
	"net/http"

	"github.com/Dias1c/forum/internal/lg"
)

//? debugLogHandler -
func debugLogHandler(fName string, r *http.Request) {
	lg.Info.Printf("%-30v | %-7v | %-30v \n", r.URL, r.Method, fName)
}
