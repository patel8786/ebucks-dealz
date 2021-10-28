package web

import (
	"embed"
	"html/template"
	"io"
	"fmt"

	"github.com/patel8786/ebucks-dealz/pkg/scraper"
)

//go:embed templates
var templatesFs embed.FS

type BaseContext struct {
	PathPrefix string
}

type DealzContext struct {
	BaseContext
	Title    string
	Products []scraper.Product
}

func RenderDealz(w io.Writer, c DealzContext) error {
fmt.Println(" pkg web render.go RenderDealz1")

	t, err := template.ParseFS(templatesFs, "templates/dealz.html.tpl")
	if err != nil {
		return err
	}
fmt.Println(" pkg web render.go RenderDealz2")
	t, err = t.ParseFS(templatesFs, "templates/common/*")
	if err != nil {
		return err
	}

fmt.Println(" pkg web render.go RenderDealz3")
	err = t.Execute(w, c)
	if err != nil {
		return err
	}
fmt.Println(" pkg web render.go RenderDealz4")
	return nil
}

func RenderHome(w io.Writer, c BaseContext) error {
	t, err := template.ParseFS(templatesFs, "templates/index.html.tpl")
	if err != nil {
		return err
	}
	t, err = t.ParseFS(templatesFs, "templates/common/*")
	if err != nil {
		return err
	}

	err = t.Execute(w, c)
	if err != nil {
		return err
	}
	return nil
}
