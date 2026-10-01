package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html/template"
	"os"
	"syscall"
	"slices"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	PID int32 `json:"pid"`
	Name string `json:"name"`
	CPU float64 `json:"cpu"`
	Memory float64 `json:"memory"`
}

type KillRequest struct {
	PID int `json:"pid"`
}

func home_page(w http.ResponseWriter, r *http.Request) {
	html_templ, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Html nicht gefunden", http.StatusInternalServerError)
		return
	}
	html_templ.Execute(w, nil)
}

func getProcessesHandler(w http.ResponseWriter, r *http.Request) {
	procs, err := process.Processes()
	if err != nil {
		http.Error(w, "Fehler beim Auslesen", http.StatusInternalServerError)
		return
	}

	var processList []ProcessInfo

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()

		processList = append(processList, ProcessInfo{
			PID: p.Pid,
			Name: name,
			CPU: cpu,
			Memory: float64(mem),
		})
	}

	slices.SortFunc(processList, func(a, b ProcessInfo) int {
		if a.CPU > b.CPU {
			return -1
		}
		if a.CPU < b.CPU {
			return 1
		}
		return 0
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(processList)

}

func killProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Nur POST erlaubt", http.StatusMethodNotAllowed)
		return
	}

	var req KillRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.PID <= 0 {
		http.Error(w, " Ungulitige PID", http.StatusBadRequest)
		return
	}

	proc, err := os.FindProcess(req.PID)
	if err != nil {
		http.Error(w, "Prozess nicht gefunden", http.StatusNotFound)
		return
	}

	err = proc.Signal(syscall.SIGTERM)
	if err != nil {
		http.Error(w, "Fehler beim Beenden des Prozesses", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func main() {

	http.HandleFunc("/", home_page)
	http.HandleFunc("/api/processes", getProcessesHandler)
	http.HandleFunc("/api/kill/", killProcessHandler)

	fmt.Println("Server run: http://localhost:8976")
	http.ListenAndServe(":9876", nil)
}