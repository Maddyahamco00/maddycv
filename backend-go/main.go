package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/register", registerHandler)
    mux.HandleFunc("/login", loginHandler)
    mux.HandleFunc("/", profileHandler)

    port := getEnv("PORT", "8080")
    addr := ":" + port
    log.Printf("Starting MaddyCV server on %s", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"message":"register endpoint pending"}`))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotImplemented)
    w.Write([]byte(`{"message":"login endpoint pending"}`))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"message":"public profile endpoint pending"}`))
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
