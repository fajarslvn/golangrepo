package go_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
	if err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// input/post a file from a html form
	// r.ParseMultiPlatform(32 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	// Create the file and decide the destination folder (./resources/)
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// Save file to destination folder
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	// Build success template by getting file name
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(
				http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8181",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/small.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Test Picture")
	file, _ := writer.CreateFormFile("file", "ExUploadFile.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	req := httptest.NewRequest("POST", "http://localhost:8181", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)

	res := rec.Result()
	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyRes))
}
