package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http" // пакет для поддержки HTTP протокола
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	name := r.URL.Query().Get("name") // значение параметра
	w.Write([]byte("Hello," + name + "!"))
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
