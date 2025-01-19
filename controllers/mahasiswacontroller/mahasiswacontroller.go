package mahasiswacontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/go_form_mahasiswa_uts/entities"
	"github.com/jeypc/go_form_mahasiswa_uts/models/mahasiswamodel"
)

var mahasiswaModel = mahasiswamodel.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/mahasiswa/index.php")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}

	temp, _ := template.New("data.php").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/mahasiswa/data.php")

	var mahasiswa []entities.Mahasiswa
	err := mahasiswaModel.FindAll(&mahasiswa)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}

	temp.ExecuteTemplate(buffer, "data.php", data)

	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	npm, err := strconv.ParseInt(queryString.Get("npm"), 10, 64)

	var data map[string]interface{}
	var mahasiswa entities.Mahasiswa

	if err != nil {
		data = map[string]interface{}{
			"title": "Tambah Data Mahasiswa",
			"mahasiswa": mahasiswa,
		}
	} else {
		
		err := mahasiswaModel.Find(npm, &mahasiswa)
		if err != nil {
			panic(err)
		}

		data = map[string]interface{}{
			"title":     "Edit Data Mahasiswa",
			"mahasiswa": mahasiswa,
		}
	}

	temp, _ := template.ParseFiles("views/mahasiswa/form.php")
	temp.Execute(w, data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		r.ParseForm()
		var mahasiswa entities.Mahasiswa

		mahasiswa.Nama = r.Form.Get("nama")
		mahasiswa.Prodi = r.Form.Get("prodi")
		mahasiswa.Kelas = r.Form.Get("kelas")

		npm, err := strconv.ParseInt(r.Form.Get("npm"), 10, 64)

		var data map[string]interface{}

		if err != nil {
			err := mahasiswaModel.Create(&mahasiswa)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil disimpan",
				"data":    template.HTML(GetData()),
			}
		} else {
			mahasiswa.NPM = npm
			err := mahasiswaModel.Update(mahasiswa)
			if err != nil {
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"message": "Data berhasil diubah",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}


func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	npm, err := strconv.ParseInt(r.Form.Get("npm"), 10, 64)
	if err != nil {
		panic(err)
	}
	err = mahasiswaModel.Delete(npm)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"message": "Data berhasil dihapus",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func Search(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        searchTerm := r.URL.Query().Get("search")

        var mahasiswa []entities.Mahasiswa
        err := mahasiswaModel.SearchByName(searchTerm, &mahasiswa)
        if err != nil {
            ResponseError(w, http.StatusInternalServerError, err.Error())
            return
        }

        data := map[string]interface{}{
            "mahasiswa": mahasiswa,
        }

        temp, _ := template.New("data.php").Funcs(template.FuncMap{
            "increment": func(a, b int) int {
                return a + b
            },
        }).ParseFiles("views/mahasiswa/data.php")

        buffer := &bytes.Buffer{}
        temp.ExecuteTemplate(buffer, "data.php", data)

        ResponseJson(w, http.StatusOK, buffer.String())
    }
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	ResponseJson(w, code, map[string]string{"error": message})
}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}