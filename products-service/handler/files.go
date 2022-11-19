package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/subhankardas/go-microservices/products-service/filestore"
)

type Files struct {
	storage *filestore.LocalFileStorage
	log     *log.Logger
}

func NewFilesHandler(logr *log.Logger) *Files {
	// Create new local storage with base path and max. file size as 5 MB
	fstore, err := filestore.NewLocalFileStorage("./uploads", 1024*1000*5, logr)
	if err != nil {
		logr.Printf("[ERROR] Unable to create local file storage with error %v", err)
	}
	return &Files{storage: fstore, log: logr}
}

func (files *Files) UploadFile(response http.ResponseWriter, request *http.Request) {
	files.log.Println("Handle POST request for file upload.")

	// Parse multipart form data from request
	err := request.ParseMultipartForm(1024 * 128)
	if err != nil {
		files.log.Printf("[ERROR] Bad request with error %v.", err)
		http.Error(response, "Invalid form data.", http.StatusBadRequest)
		return
	}

	// Read and validate product ID in form data
	id, err := strconv.Atoi(request.FormValue("id"))
	if err != nil {
		files.log.Printf("[ERROR] Bad request with invalid product ID = %v.", id)
		http.Error(response, "Product ID not found in form data.", http.StatusBadRequest)
		return
	}

	// Read and validate multipart file in form data
	file, fheader, err := request.FormFile("file")
	if err != nil {
		files.log.Println("File not found.")
		return
	}

	// Create filepath i.e. id/filename and save into storage
	fpath := filepath.Join(request.FormValue("id"), fheader.Filename)
	err = files.storage.Save(fpath, file)
	if err != nil {
		files.log.Printf("[ERROR] Failed uploading file with error %v.", err)
		http.Error(response, "Failed uploading file.", http.StatusInternalServerError)
	}
}
