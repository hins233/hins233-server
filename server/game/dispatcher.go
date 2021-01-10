package game

import (
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"path"
	"server/server/game/service"
)

// jps 算法

// 分发到service中
const (
	DispatcherKeyGame = "gameId"
	DispatcherKeyMsg  = "msgId"
)

func Dispatcher(param map[string]interface{}, conn net.Conn) {
	gameId, ok := param[DispatcherKeyGame].(float64)
	if !ok {
		return
	}
	msgId, ok := param[DispatcherKeyMsg].(float64)
	if !ok {
		return
	}
	data := param["data"].(map[string]interface{})
	service.Action(int(gameId), int(msgId), data, conn)
}

/**
==== 下面是 html 模板 =====
*/
const (
	TemplateDir = "resources/templates"
)

var templates = make(map[string]*template.Template)

func RenderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	if tmpl == "" {
		tmpl = "index"
	} else {
		tmpl = "g" + tmpl
	}
	err := templates[getTemplateName(TemplateDir, tmpl)].Execute(w, locals)
	checkErr(err)
}

func getTemplateName(pre, next string) string {
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
	fileInfoArr, err := ioutil.ReadDir(TemplateDir)
	checkErr(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TemplateDir + "/" + templateName
		t := template.Must(template.ParseFiles(templatePath))
		templates[templatePath] = t
	}
}
