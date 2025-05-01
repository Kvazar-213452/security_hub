package main_com

import (
	"encoding/json"
	"fmt"
	"head/main_com/module"
	"io"
	"net/http"
	"time"
)

func Post_install_model_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var request struct {
			Data string `json:"Data"`
		}

		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		err = module.Install_module(request.Data)
		var request_front int
		if err != nil {
			request_front = 0
		} else {
			request_front = 1
			module.MoveModuleToUninstall(request.Data)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(request_front)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Post_uinstall_model_app(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var request struct {
			Data string `json:"Data"`
		}

		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		err = module.UninstallModule(request.Data)
		var request_front int
		if err != nil {
			request_front = 0
		} else {
			request_front = 1
			module.MoveModuleToInstall(request.Data)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(request_front)
	} else {
		http.Error(w, "error", http.StatusMethodNotAllowed)
	}
}

func Rost_reload_model(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		module.KillAllModules()

		time.Sleep(1 * time.Second)

		err := module.RunModules(
			"../data/config_module.json",
			"result.json",
		)
		if err != nil {
			fmt.Printf("Помилка: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nil)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
