package main

import (
	"log"
	"net/http"

	"github.com/mathyourlife/arthur/tib"
)

var (
	debug = false
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	pd := NewPageData().HasJQuery().HasBootstrap().HasMath()

	pd.Header.Styles = append(pd.Header.Styles, &HTMLStyle{
		HRef:      "/static/css/styles.css",
		Integrity: "sha384-Qvp4fqp9WtSvyDRfyJHCJ+8WHsrzGFi/TImxiqRt9VhWGTUvSVvkb8bwSWR062NG",
	})
	pd.Body.Scripts = append(pd.Body.Scripts, &HTMLScript{
		Src:       "/static/js/arthur.js",
		Integrity: "sha384-Gt8OHrojQG0oL1bFYZzZR4jtxcslsKaTaOS4mAUQcr1Mlhf7bueUPSktS2HYfyqL",
	})

	t := templates["index.html"]
	if t == nil {
		log.Println("[ERROR] index.html template does not exist")
		return
	}
	err := t.ExecuteTemplate(w, "index.html", pd)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}
}

func main() {
	log.Println("starting server")
	debug = true

	parseTemplates("html")
	watcher := watchDirs("html")
	defer watcher.Close()

	mux := http.NewServeMux()
	tib.SetupMux("api/v1/", mux)

	mux.HandleFunc("/", handleMain)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("http server failed: %s", err)
	}

	http.ListenAndServe(":8080", nil)
}
