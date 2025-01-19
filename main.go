package main

import (
	"net/http"

	"github.com/jeypc/go_form_mahasiswa_uts/controllers/mahasiswacontroller"
)

func main() {
	http.HandleFunc("/", mahasiswacontroller.Index)
	http.HandleFunc("/mahasiswa/get_form", mahasiswacontroller.GetForm)
	http.HandleFunc("/mahasiswa/store", mahasiswacontroller.Store)
	http.HandleFunc("/mahasiswa/delete", mahasiswacontroller.Delete)
	http.HandleFunc("/mahasiswa/search", mahasiswacontroller.Search)


	http.ListenAndServe(":8000", nil)
}