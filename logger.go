package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/sudo-adduser-jordan/gcolor"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (record *statusRecorder) WriteHeader(code int) {
	record.status = code
	record.ResponseWriter.WriteHeader(code)
}

func Logger(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Initialize the status to 500 in case WriteHeader is not called
		record := statusRecorder{w, 500}
		defer func() {
			log.Printf(
				"| %s |	%s|	%s %s %s",
				statusColor(record.status),
				time.Since(start),
				gcolor.GreenText(r.Host),
				methodColor(r.Method),
				r.URL.Path,
			)

		}()
		f.ServeHTTP(&record, r)
	})
}

func statusColor(status int) lipgloss.Style {
	s := strconv.Itoa(status)
	s = fmt.Sprintf(" " + s + " ")

	switch status {
	case 200:
		return gcolor.GreenText(s)
	case 300:
		return gcolor.YellowText(s)
	case 400:
		return gcolor.BlueText(s)
	case 500:
		return gcolor.RedText(s)
	default:
		return gcolor.RedText(s)
	}
}

func methodColor(method string) lipgloss.Style {
	s := fmt.Sprintf(" " + method + " ")
	switch method {
	case "GET":
		return gcolor.BlueLabel(s)
	case "POST":
		return gcolor.GreenLabel(s)
	case "PUT":
		return gcolor.PurpleLabel(s)
	case "DELETE":
		return gcolor.RedLabel(s)
	default:
		return gcolor.RedText(s)
	}
}
