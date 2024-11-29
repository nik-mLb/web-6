package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

var counter int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%d", counter)))
	case http.MethodPost:
		err := r.ParseForm()
		if err == nil {
			countStr := r.FormValue("count")
			var count int
			_, err2 := fmt.Sscanf(countStr, "%d", &count)
			if err2 != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("это не число"))
			} else {
				w.WriteHeader(http.StatusOK)
				counter += count
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
