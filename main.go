package main

import (
	"fmt"
	"net/http"
	"text/template"

	asciiart "main.go/printascii"
)

type asciiART struct {
	Art string
}

type PageVariables struct {
	Code int
}

func main() {
	fmt.Println("Ctrl + Click1 >> http://localhost:8080/")
	http.Handle("/index/", http.StripPrefix("/index/", http.FileServer(http.Dir("./index"))))
	http.HandleFunc("/", MainPageFunc)
	http.HandleFunc("/ascii-art", Asciiart)
	http.ListenAndServe(":8080", nil)
}

func MainPageFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index/index.html")
	if err != nil {
		errorCode := http.StatusInternalServerError
		errorText := "Internal Server Error"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	if r.URL.Path != "/" {
		errorCode := http.StatusNotFound
		errorText := "Status not Found"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	if r.Method != http.MethodGet {
		errorCode := http.StatusMethodNotAllowed
		errorText := "Method Not Allowed"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}

	if t == nil {
		errorCode := http.StatusInternalServerError
		errorText := "Internal Server Error"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		errorCode := http.StatusInternalServerError
		errorText := "Internal Server Error"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
}

func Asciiart(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		errorCode := http.StatusNotFound
		errorText := "Status not Found"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	if r.Method != http.MethodPost {
		errorCode := http.StatusMethodNotAllowed
		errorText := "Method Not Allowed"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	t, err := template.ParseFiles("index/index.html")
	if err != nil {
		errorCode := http.StatusInternalServerError
		errorText := "Internal Server Error"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	text := r.FormValue("asciiInput")
	style := r.FormValue("style-selector")

	res := string(text)
	for _, ch := range res {
		if !(ch >= 32 && ch <= 126) && ch != 10 && ch != 13 {
			errorCode := http.StatusBadRequest
			errorText := "Bad Request, you have to input only ASCII symbols"
			ErrorHandler(w, r, errorCode, errorText)
			return
		}
	}
	art, err := asciiart.GetT(text, style)
	if err != nil {
		errorCode := http.StatusBadRequest
		errorText := "Bad Request"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
	data := asciiART{
		Art: art,
	}

	err = t.Execute(w, data)
	if err != nil {
		errorCode := http.StatusInternalServerError
		errorText := "Internal Server Error"
		ErrorHandler(w, r, errorCode, errorText)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, errorCode int, errorText string) {
	data := PageVariables{
		Code: errorCode,
	}

	tmpl, err := template.ParseFiles("index/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(errorCode)

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
