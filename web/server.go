//go:generate statik -src=./static -dest=./ -f

/*
IMPORTANT: all static files are embedded inside the final binary.
*/

package web

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/audisto/data-downloader/web/statik" // compiled static files
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

var (
	// EmbeddedFS Keep a reference to the embedded FileSystem holding our static files and templates
	EmbeddedFS http.FileSystem
	// TemplatesFSPrefix the path to template files *inside* the embedded FileSystem
	TemplatesFSPrefix = "/templates/"
	// TemplateFiles the list of template files to look for inside the embedded FileSystem
	TemplateFiles = [...]string{
		"footer.html", "header.html", "home.html"}
)

func init() {
	var err error
	EmbeddedFS, err = fs.New()
	if err != nil {
		log.Fatal(err)
	}
}

// StartWebInterface -
func StartWebInterface(port uint, debug bool) {

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	webDownloader := NewWebDownloader()
	server.SetHTMLTemplate(getTemplates())
	server.Use(Logger())
	server.Use(gin.Recovery())
	server.StaticFS("/static", EmbeddedFS)
	server.GET("/", webDownloader.homeHandler)
	server.POST("/login", webDownloader.doLogin)
	server.GET("/logout", webDownloader.doLogout)
	server.POST("/download", webDownloader.downloadHandler)
	server.POST("/stop", webDownloader.stopHandler)
	server.GET("/progress", webDownloader.progressHandler)

	fmt.Printf(banner, port)
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	server.Run(addr)
}

func getTemplates() *template.Template {
	tmpl := template.New("")

	for _, tmpName := range TemplateFiles {

		templateFile, err := EmbeddedFS.Open(TemplatesFSPrefix + tmpName)
		if err != nil {
			log.Fatal(err)
		}
		fileBytes, err := ioutil.ReadAll(templateFile)
		if err != nil {
			log.Fatal(err)
		}
		tmpl, err = tmpl.New(tmpName).Parse(string(fileBytes))
		if err != nil {
			log.Fatal(err)
		}
	}

	return tmpl
}
