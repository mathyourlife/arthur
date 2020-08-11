package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	templates            map[string]*template.Template
	StaticAssetIntegrity map[string]string
)

type PageData struct {
	Header struct {
		Styles []*HTMLStyle
	}
	Body struct {
		Container interface{}
		Navbar    struct {
			ActivePage string
		}
		Scripts []*HTMLScript
	}
}

type HTMLStyle struct {
	HRef      string
	Integrity string
}

type HTMLScript struct {
	Src       string
	Integrity string
	Code      template.JS
}

func NewPageData() *PageData {
	d := &PageData{}
	d.Header.Styles = []*HTMLStyle{}
	d.Body.Scripts = []*HTMLScript{}
	return d
}

func (d *PageData) HasBootstrap() *PageData {
	d.Header.Styles = append(d.Header.Styles, &HTMLStyle{
		HRef:      "https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css",
		Integrity: "sha384-9aIt2nRpC12Uk9gS9baDl411NQApFmC26EwAOH8WgZl5MYYxFfc+NcPb1dKGj7Sk",
	})
	d.Body.Scripts = append(d.Body.Scripts, &HTMLScript{
		Src:       "https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/js/bootstrap.min.js",
		Integrity: "sha384-OgVRvuATP1z7JjHLkuOU7Xw704+h835Lr+6QL9UvYjZE3Ipu6Tp75j7Bh/kR0JKI",
	})
	return d
}

func (d *PageData) HasJQuery() *PageData {
	d.Body.Scripts = append(d.Body.Scripts, &HTMLScript{
		Src:       "https://code.jquery.com/jquery-3.5.1.min.js",
		Integrity: "sha256-9/aliU8dGd2tb6OSsuzixeV4y/faTqgFtohetphbbj0=",
	})
	return d
}

func (d *PageData) HasMath() *PageData {
	d.Header.Styles = append(d.Header.Styles, &HTMLStyle{
		HRef:      "https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.css",
		Integrity: "sha384-AfEj0r4/OFrOo5t7NnNe46zW/tFgW6x/bCJG8FqQCEo3+Aro6EYUG4+cU+KJWu/X",
	})
	d.Body.Scripts = append(d.Body.Scripts, &HTMLScript{
		Src:       "https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.js",
		Integrity: "sha384-g7c+Jr9ZivxKLnZTDUhnkOnsh30B4H0rpLUpJ4jAIKs4fnJI+sEnkvrMWph2EDg4",
	})
	d.Body.Scripts = append(d.Body.Scripts, &HTMLScript{
		Src:       "https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/contrib/auto-render.min.js",
		Integrity: "sha384-mll67QQFJfxn0IYznZYonOWZ644AWYC+Pt2cHqMaRhXVrursRwvLnLaebdGIlYNa",
	})
	d.Body.Scripts = append(d.Body.Scripts, &HTMLScript{Code: template.JS(`
document.addEventListener("DOMContentLoaded", function() {
    renderMathInElement(document.body, {
        delimiters:
        [{left: "$$", right: "$$", display: false},
         {left: "\\(", right: "\\)", display: false},
         {left: "\\[", right: "\\]", display: true},]
    });
});
`)})
	return d
}

func parseTemplates(htmlTmplDir string) {
	templates = make(map[string]*template.Template)

	components := []string{}
	err := filepath.Walk(path.Join(htmlTmplDir, "components"), func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			components = append(components, path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	if debug {
		log.Printf("components: %s", components)
	}
	views := []string{}
	err = filepath.Walk(path.Join(htmlTmplDir, "views"), func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			views = append(views, path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		log.Printf("views: %s", views)
	}
	for _, view := range views {
		files := append(components, view)
		templates[path.Base(view)] = template.Must(template.New(path.Base(view)).Funcs(template.FuncMap{}).ParseFiles(files...))
	}
}

func watchDirs(htmlTmplDir string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				log.Println(event, ok)
				if !ok {
					return
				}
				log.Println("event:", event)
				parseTemplates(htmlTmplDir)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = filepath.Walk(htmlTmplDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return nil
			}
			fmt.Println(path, info.Size())
			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return watcher
}
