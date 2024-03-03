package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/sudo-adduser-jordan/go/styles"
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
				styles.GreenText(r.Host),
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
		return styles.GreenText(s)
	case 300:
		return styles.YellowText(s)
	case 400:
		return styles.BlueText(s)
	case 500:
		return styles.RedText(s)
	default:
		return styles.RedText(s)
	}
}

func methodColor(method string) lipgloss.Style {
	s := fmt.Sprintf(" " + method + " ")
	switch method {
	case "GET":
		return styles.BlueLabel(s)
	case "POST":
		return styles.GreenLabel(s)
	case "PUT":
		return styles.PurpleLabel(s)
	case "DELETE":
		return styles.RedLabel(s)
	default:
		return styles.RedText(s)
	}
}
