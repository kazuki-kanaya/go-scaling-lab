package main

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	var ready = true

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("not ready\n"))
			return
		}
	})

	mux.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		ms := queryDurationMS(r, 100)
		burnCPU(ms)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	mux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		ms := queryDurationMS(r, 100)
		time.Sleep(time.Duration(ms) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
func queryDurationMS(r *http.Request, fallback int) int {
	ms, err := strconv.Atoi(r.URL.Query().Get("ms"))
	if err != nil || ms <= 0 {
		return fallback
	}
	return ms
}

var cpuSink float64

func burnCPU(ms int) {
	deadline := time.Now().Add(time.Duration(ms) * time.Millisecond)
	value := 0.0
	for time.Now().Before(deadline) {
		value += math.Sqrt(value + 1.2345)
		if value > 1e6 {
			value = 0
		}
	}
	cpuSink = value
}
