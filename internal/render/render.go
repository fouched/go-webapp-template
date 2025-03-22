package render

import (
	"embed"
	"fmt"
	"github.com/fouched/go-webapp-template/internal/config"
	"html/template"
	"net/http"
)

var app *config.App

// NewRenderer sets the config for the template package
func NewRenderer(a *config.App) {
	app = a
}

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	BoolMap   map[string]bool
	Data      map[string]interface{}
	CSRFToken string
	Success   string
	Warning   string
	Error     string
	//AuthLevel  int
	//UserID     int
	//UserName   string
	CSSVersion string
	//Validator  *validator.Validator
}

// with the go embed directive below we can compile
// the templates with the application in a single binary
//
//go:embed static/* static/css/* static/img/* static/js/*
//go:embed templates/* templates/customer/*
var embedFS embed.FS

var functions = template.FuncMap{
	"inc":          inc,
	"unescapeHTML": unescapeHTML,
}

func EmbedFS() embed.FS {
	return embedFS
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func inc(i int) int {
	return i + 1
}

// Template TODO remove passing of path - let handlers just pass full paths, which simplifies logic below
func Template(w http.ResponseWriter, r *http.Request, path string, page string, td *TemplateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)
	_, templateInMap := app.TemplateCache[templateToRender]

	if templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = parseTemplate(path, page, partials)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
		app.TemplateCache[templateToRender] = t
	}

	if td == nil {
		td = &TemplateData{}
	}

	td = addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.ErrorLog.Println(err)
		return err
	}

	return nil
}

func parseTemplate(path, page string, partials []string) (*template.Template, error) {
	var t *template.Template
	var err error

	templates := []string{"templates/base.layout.gohtml", fmt.Sprintf("templates/%s/%s.page.gohtml", path, page)}

	// build templates
	if len(partials) > 0 {
		for _, p := range partials {
			templates = append(templates, fmt.Sprintf("templates/%s.partial.gohtml", p))
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).
			ParseFS(embedFS, templates...)
	} else {
		if path == "" || path == "/" {
			t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).
				ParseFS(embedFS, "templates/base.layout.gohtml", fmt.Sprintf("templates/%s.page.gohtml", page))
		} else {
			t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).
				ParseFS(embedFS, "templates/base.layout.gohtml", fmt.Sprintf("templates/*/%s.page.gohtml", page))
		}
	}

	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return t, nil
}

// Partial TODO remove passing of path - let handlers should just pass full paths, which simplifies logic below
func Partial(w http.ResponseWriter, r *http.Request, path string, partial string, td *TemplateData, additionalPartials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s/%s.partial.gohtml", path, partial)

	_, templateInMap := app.TemplateCache[templateToRender]

	if templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = parsePartial(path, partial, additionalPartials)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
		app.TemplateCache[templateToRender] = t
	}

	if td == nil {
		td = &TemplateData{}
	}

	td = addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.ErrorLog.Println(err)
		return err
	}

	return nil
}

func parsePartial(path, partial string, additionalPartials []string) (*template.Template, error) {
	var t *template.Template
	var err error

	templates := []string{fmt.Sprintf("templates/%s/%s.partial.gohtml", path, partial), "templates/toast.partial.gohtml", "templates/pagination.partial.gohtml"}

	// build additional templates
	if len(additionalPartials) > 0 {
		for _, p := range additionalPartials {
			templates = append(templates, fmt.Sprintf("templates/%s/%s.partial.gohtml", path, p))
		}
	}

	if len(additionalPartials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.partial.gohtml", partial)).Funcs(functions).
			ParseFS(embedFS, templates...)
	} else {
		t, err = template.New(fmt.Sprintf("%s.partial.gohtml", partial)).Funcs(functions).
			ParseFS(embedFS, "templates/toast.partial.gohtml", fmt.Sprintf("templates/*/%s.partial.gohtml", partial))
	}

	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return t, nil
}

func addDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Success = app.Session.PopString(r.Context(), "success")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	return td
}
