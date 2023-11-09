/*
Copyright © 2023 Ajith

*/
package server

import (
	"os"
	"log"
	"time"
	"bytes"
	"embed"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"html/template"

	"github.com/spf13/cobra"
	"github.com/gabriel-vasile/mimetype"
)

var rootDir string

//go:embed templates/*
var f embed.FS

func indexHTMLTemplateHandler(response http.ResponseWriter, request *http.Request) {
	relativePath := strings.TrimPrefix(request.URL.Path, "/")
	log.Println(filepath.Clean(relativePath + "/"))

	path := filepath.Join(rootDir, relativePath)
	
	fileInfo, err := os.Stat(path)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		if err != nil {
			log.Println(err)
		}
		tmpl := template.Must(template.ParseFS(f, "templates/404.html"))
		tmpl.Execute(response, nil)
	} else {

		if fileInfo.IsDir() {
			searchTerm := request.URL.Query().Get("s")
			content, err := GetContents(path, relativePath, searchTerm)

			if err != nil {
				log.Println(err)
				response.WriteHeader(http.StatusInternalServerError)
				tmpl := template.Must(template.ParseFS(f, "templates/500.html"))
				tmpl.Execute(response, nil)
			}
			
			tmpl := template.Must(template.ParseFS(f, "templates/index.html"))
			tmpl.Execute(response, content)
		
		} else {
		
			data, err := ioutil.ReadFile(path)
			if(err != nil){
				log.Println(err)
				response.WriteHeader(http.StatusInternalServerError)
				tmpl := template.Must(template.ParseFS(f, "templates/500.html"))
				tmpl.Execute(response, nil)
			}

			mtype, err := mimetype.DetectFile(path)
			if err != nil {

				response.WriteHeader(http.StatusInternalServerError)
				tmpl := template.Must(template.ParseFS(f, "templates/500.html"))
				tmpl.Execute(response, nil)

			}

			response.Header().Set("Content-Disposition", "attachment; filename=" + fileInfo.Name())
			response.Header().Set("Content-Type", mtype.String())
			response.WriteHeader(http.StatusOK)
			http.ServeContent(response, request, path, time.Now(), bytes.NewReader(data))
		}
	}
}

func Init(cmd *cobra.Command, args []string) {
	port, _ := cmd.Flags().GetString("port")
	directory, _ := cmd.Flags().GetString("directory")

	// remove trailing quotes from the directory path
	directory = (strings.Replace(filepath.Clean(directory), "\"", "", -1))

	http.HandleFunc("/", indexHTMLTemplateHandler)

	var err error
	if directory != "" {
		rootDir, err = filepath.Abs(directory)
	} else {
		rootDir, err = os.Getwd()
	}
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix("HLS: ")
	log.Printf("Serving %s on %s:%s\n", rootDir, GetOutboundIP(), port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
	log.Println("Server Shutdown.")

}
