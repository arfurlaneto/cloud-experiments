package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type AppState struct {
	StartedAt time.Time
	Version   string
}

func NewAppState() *AppState {
	appState := AppState{}

	appState.StartedAt = time.Now()

	appVersion := os.Getenv("APP_VERSION")
	if appVersion == "" {
		appVersion = "UNKNOW"
	}
	appState.Version = appVersion

	return &appState
}

func (appState *AppState) Healthz(w http.ResponseWriter, r *http.Request) {
	output := fmt.Sprintf("Running Version %s. Started at %s.", appState.Version, appState.StartedAt)
	w.WriteHeader(200)
	w.Write([]byte(output))
}

func (appState *AppState) GetValuesFromConfigMapEnv(w http.ResponseWriter, r *http.Request) {
	value1 := os.Getenv("VALUE_1")
	value2 := os.Getenv("VALUE_2")
	fmt.Fprintf(w, "Values set at 'configmap-env' are %s=%s, %s=%s", "VALUE_1", value1, "VALUE_2", value2)
}

func (appState *AppState) GetValuesFromConfigMapFile(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/go/files/text.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Fprintf(w, "Values set at 'configmap-file.yaml' are %s=%s", "/go/files/text.txt", data)
}

func (appState *AppState) GetValuesFromSecret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "Values set at 'secret.yaml' are %s=%s, %s=%s", "USER", user, "PASSWORD", password)
}

func main() {

	state := NewAppState()

	fmt.Printf("Starting Version %s...\n", state.Version)
	http.HandleFunc("/", state.Healthz)
	http.HandleFunc("/configmapenv", state.GetValuesFromConfigMapEnv)
	http.HandleFunc("/configmapfile", state.GetValuesFromConfigMapFile)
	http.HandleFunc("/secret", state.GetValuesFromSecret)
	http.ListenAndServe(":8000", nil)
}
