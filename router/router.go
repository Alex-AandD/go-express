package router

import (
	"net/http"
	"regexp"
	"github.com/go-express/handler"
)

//type FuncHandler func(env *env.Env, w *http.ResponseWriter, r http.Request) error


type RouteEntry struct {
	Path string
	Method string
	Params map[string]any
	Handler handler.Handler
}


type Router struct {
	BasePath string
	Entries []RouteEntry
}

func expandPath(path string) string {
	// first extract the name of the parameters
	re := regexp.MustCompile(":([a-zA-Z0-9]+)")
	expandedPath := re.ReplaceAllString(path, "(?P<$1>[a-zA-Z0-9]+)")
	expandedPath = "^" + expandedPath + "$"
	return expandedPath
}

func (rtr *Router) Get(path string, handler handler.Handler) {
	// register the route inside of the entries
	expPath := expandPath(path)
	entry := RouteEntry{Path: expPath, Method: "GET", Handler: handler }
	rtr.Entries = append(rtr.Entries, entry)
}

func (rtr *Router) Post(path string, handler handler.Handler) {
	expPath := expandPath(path)
	entry := RouteEntry{Path: expPath, Method: "POST", Handler: handler }
	rtr.Entries = append(rtr.Entries, entry)
}

func (rtr *Router) Put(path string, handler handler.Handler) {
	expPath := expandPath(path)
	entry := RouteEntry{Path: expPath, Method: "PUT", Handler: handler }
	rtr.Entries = append(rtr.Entries, entry)
}

func (rtr *Router) Delete(path string, handler handler.Handler) {
	expPath := expandPath(path)
	entry := RouteEntry{Path: expPath, Method: "DELETE", Handler: handler }
	rtr.Entries = append(rtr.Entries, entry)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the path of the request
	path := r.URL.Path
	if entry := rtr.Match(path); entry != nil {
		if r.Header.Get("method") != entry.Method {
		}
		entry.Handler.H(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func (rtr *Router) Match(path string) *RouteEntry {
	for _, entry := range rtr.Entries {
		re := regexp.MustCompile(entry.Path)
		matches := re.FindStringSubmatch(path) 
		if matches != nil {
			// good route found, proceed to get the parameters
			params := make(map[string]any)
			for i, name := range re.SubexpNames() {
				if i > 0 {
					params[name] = matches[i]
				}
			}
			entry.Params = params
			return &entry
		}
	}
	return nil 
}