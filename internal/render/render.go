package render

import (
	"embed"
	"fmt"
	"github.com/fouched/go-webapp-template/internal/config"
	"html/template"
	"net/http"
	"strings"
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
//go:embed templates
var templateFS embed.FS

var functions = template.FuncMap{
	"inc":          inc,
	"unescapeHTML": unescapeHTML,
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func inc(i int) int {
	return i + 1
}

func Template(w http.ResponseWriter, r *http.Request, page string, td *TemplateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	_, templateInMap := app.TemplateCache[templateToRender]

	if templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
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

func parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).
			ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).
			ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}

	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	app.TemplateCache[templateToRender] = t

	return t, nil
}

func Partial(w http.ResponseWriter, r *http.Request, partial string, td *TemplateData, additionalPartials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.partial.gohtml", partial)

	_, templateInMap := app.TemplateCache[templateToRender]

	if templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = parsePartial(additionalPartials, partial, templateToRender)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
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

func parsePartial(additionalPartials []string, partial, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build additional partials
	if len(additionalPartials) > 0 {
		for i, x := range additionalPartials {
			additionalPartials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	if len(additionalPartials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.partial.gohtml", partial)).Funcs(functions).
			ParseFS(templateFS, "templates/toast.partial.gohtml", strings.Join(additionalPartials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.partial.gohtml", partial)).Funcs(functions).
			ParseFS(templateFS, "templates/toast.partial.gohtml", templateToRender)
	}

	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	app.TemplateCache[partial] = t

	return t, nil
}

func addDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Success = app.Session.PopString(r.Context(), "success")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	return td
}
