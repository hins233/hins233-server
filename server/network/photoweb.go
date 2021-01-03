package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

const (
	ListDir      = 0x0001
	TEMPLATE_DIR = "resources/templates"
	UPLOAD_DIR   = "resources/uploads"
)

var templates = make(map[string]*template.Template)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
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

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "index", nil)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		checkErr(err)
		fileName := h.Filename
		defer f.Close()
		t, err := ioutil.TempFile(UPLOAD_DIR, fileName)
		checkErr(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		checkErr(err)
		http.Redirect(w, r, "/view?id="+fileName, http.StatusFound)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exits := isExists(imagePath); !exits {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	checkErr(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Printf("warn: panic in %v -%v/n", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, statisDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := statisDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exits := isExists(file); !exits {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./resources", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe failed:", err.Error())
	}
}
