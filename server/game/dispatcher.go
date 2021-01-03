package game

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path"
)

const (
	TEMPLATE_DIR = "resources/templates"
)

// jps 算法

var templates = make(map[string]*template.Template)

func RenderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	if tmpl == "" {
		tmpl = "index"
	} else {
		tmpl = "g" + tmpl
	}
	err := templates[fixAppendStr(TEMPLATE_DIR, tmpl)].Execute(w, locals)
	checkErr(err)
}

func fixAppendStr(pre, next string) string {
	next += ".html"
	if pre[len(pre)-1] == '/' && next[0] == '/' {
		return pre + next[1:]
	}
	if pre[len(pre)-1] != '/' && next[0] != '/' {
		return pre + "/" + next
	}
	return pre + next
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	checkErr(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}
