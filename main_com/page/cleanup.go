package page

import (
	"fmt"
	config_main "head/main_com/config"
	"head/main_com/func_all"
	"net/http"
	"syscall"
)

//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post
//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post//post

func Post_cleanup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		func_all.AppendToLog("cleanup")

		cleanup()

		w.Write(nil)
	} else {
		http.Error(w, "Непідтримуваний метод", http.StatusMethodNotAllowed)
	}
}

//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func
//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func
//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func
//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func//func

func cleanup() {
	cleanupDLL, err := syscall.LoadDLL(config_main.Cleanup_dll)
	if err != nil {
		fmt.Printf("Не вдалося завантажити DLL: %v\n", err)
		return
	}
	defer cleanupDLL.Release()

	cleanupProc, err := cleanupDLL.FindProc("cleanup")
	if err != nil {
		return
	}

	cleanupProc.Call()
}
