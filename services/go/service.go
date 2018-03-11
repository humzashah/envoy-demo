package main

import (
  "fmt"
  "io"
  "net/http"
  "net/http/httptrace"
  "os"
  "strconv"
  "time"
)

func service_name () string {
  service_name, _ := os.LookupEnv("THIS_SERVICE_NAME")
  return service_name
}

func formatted_time () string {
  return time.Now().Format(time.RubyDate)
}

func health (w http.ResponseWriter, r *http.Request) {
  status_code_str, exists := os.LookupEnv("DUMMY_STATUS_CODE")
  if exists {
    status_code_int, _ := strconv.Atoi(status_code_str)
    http.Error(w, status_code_str, status_code_int)
  } else {
    io.WriteString(w, "OK /health\n")
  }
  fmt.Println(formatted_time() + " " + service_name() + " served /health")
}

func info (w http.ResponseWriter, r *http.Request) {
  status_code_str, exists := os.LookupEnv("DUMMY_STATUS_CODE")
  if exists {
    status_code_int, _ := strconv.Atoi(status_code_str)
    http.Error(w, status_code_str, status_code_int)
  } else {
    io.WriteString(w, "OK /info\n")
  }
  fmt.Println(formatted_time() + " " + service_name() + " served /info")
}

func jump (w http.ResponseWriter, r *http.Request) {
  status_code_str, exists := os.LookupEnv("DUMMY_STATUS_CODE")
  if exists {
    status_code_int, _ := strconv.Atoi(status_code_str)
    http.Error(w, status_code_str, status_code_int)
  } else {
    target, _ := os.LookupEnv("JUMP_TARGET_URL")
    req, _ := http.NewRequest("GET", target, nil)
    trace := &httptrace.ClientTrace{
      DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
        fmt.Printf("DNS Info: %+v\n", dnsInfo)
      },
      GotConn: func(connInfo httptrace.GotConnInfo) {
        fmt.Printf("Got Conn: %+v\n", connInfo)
      },
    }
    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
    http.DefaultTransport.RoundTrip(req)
    io.WriteString(w, "OK /jump\n")
  }
  fmt.Println(formatted_time() + " " + service_name() + " served /jump")
}

func main () {
  fmt.Println("Running " + service_name())
  http.HandleFunc("/health", health)
  http.HandleFunc("/info", info)
  http.HandleFunc("/jump", jump)
  http.ListenAndServe(os.Getenv("THIS_SERVICE_ADDRESS"), nil)
}
