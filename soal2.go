package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type respon struct {
	Input  string
	Output string
}

func sorter(inputan string) string {
	arr := strings.Split(inputan, "")
	vocal := "aAiIuUeEoO"
	hurufHidup := []string{}
	hurufMati := []string{}
	huruf := []string{}
	for i := 0; i < len(arr); i++ {
		if strings.Contains(vocal, arr[i]) {
			hurufHidup = append(hurufHidup, arr[i])
		} else {
			hurufMati = append(hurufMati, arr[i])
		}
	}
	sort.Strings(hurufHidup)
	sort.Strings(hurufMati)
	huruf = append(hurufHidup, hurufMati...)
	return strings.Join(huruf, "")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inputan string
	if len(r.FormValue("input")) == 0 {
		http.Error(w, "please fill the input", http.StatusInternalServerError)
		return
	} else if len(r.URL.Query()["input"][0]) != 0 {
		inputan = r.URL.Query()["input"][0]
	} else {
		inputan = r.FormValue("input")
	}
	if r.Method == "POST" {
		var err error
		data := []respon{
			respon{inputan, sorter(inputan)},
		}
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "GET" {
		data := []respon{
			respon{inputan, sorter(inputan)},
		}
		result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "request method is false", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
