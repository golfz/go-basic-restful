package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/now", now)
	http.HandleFunc("/env", env)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func now(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Now string `json:"now"`
	}{
		Now: time.Now().Format(time.RFC3339),
	}

	b, _ := json.Marshal(data)
	fmt.Fprintln(w, string(b))
}

func env(w http.ResponseWriter, r *http.Request) {
	envA := os.Getenv("ENV_A")
	envB := os.Getenv("ENV_B")

	data := struct {
		EnvA string `json:"env_a"`
		EnvB string `json:"env_b"`
	}{
		EnvA: envA,
		EnvB: envB,
	}

	b, _ := json.Marshal(data)
	fmt.Fprintln(w, string(b))
}
