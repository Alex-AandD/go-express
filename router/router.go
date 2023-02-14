package router

/*
	path
	stack 		middleware of functions
	methods		map{ "GET": func, "Post": other function}

*/
import (
	"fmt"
	"net/http"
	"regexp"
	"github.com/go-express/request"
	"github.com/go-express/errors"
)

type NextFunction func() errors.Error
type HandlerFunc func(w http.ResponseWriter, r *request.Request, next NextFunction) errors.Error

type Route struct {
	Path 		string
	Stack	 	[]HandlerFunc
	Method		string 	
}

type Router struct {
	BasePath	string
	Rts 		[]*Route
}

func (rtr *Router) Get(path string, handlers ...HandlerFunc) errors.Error {
	expPath := expandPath(path)

	for _, e := range rtr.Rts {
		if e.Path == expPath && e.Method == "GET" {
			return errors.NewRouterError(fmt.Sprintf("Router Error - route %s already covers GET method", path))
		}
	}
	// create route
	route := &Route{
		Path: expPath,
		Stack: handlers,
		Method: "GET",
	}
	// add route
	rtr.Rts = append(rtr.Rts, route)
	return nil
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create a new request object
	R := &request.Request{R: r}
	// get the path
	path := r.URL.Path
	
	// check if a path exists
	if route, params := rtr.findPath(path); route != nil {
		// attach the params to the request
		R.Params = params
		// check if the method is correct
		if route.Method != R.GetMethod() {
			//http.Error(w, "Method not allowed on this route", http.StatusBadRequest) 
			handleError(w, errors.NewStatusError("method not allowed on this route", http.StatusBadRequest))
		}
		if err := handleRoute(route, w, R); err != nil {
			handleError(w, &err)
		}
	}
}

func handleError(w http.ResponseWriter, err errors.Error) {
	switch e := err.(type) {
		case errors.StatusError:
			http.Error(w, e.Error(), e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func handleRoute(route *Route, w http.ResponseWriter, r *request.Request) errors.Error {
	// next function
	curr := 0
	var next func() errors.Error
	next = func () errors.Error {
		if (curr == len(route.Stack) - 1) {
			currHandler := route.Stack[curr]
			return currHandler(w, r, func () errors.Error { return nil })
		}

		if (curr >= len(route.Stack)) {
			return errors.NewRouterError("router error")
		}
		currHandler := route.Stack[curr]
		curr++
		return currHandler(w, r, next)
	}
	next()
	return nil
}

func (rtr *Router) findPath(path string) (*Route, map[string]string) {
	for _, r:= range rtr.Rts {
		re := regexp.MustCompile(r.Path)
		matches := re.FindStringSubmatch(path) 
		if matches != nil {
			// good route found, proceed to get the parameters
			params := make(map[string]string)
			for i, name := range re.SubexpNames() {
				if i > 0 {
					params[name] = matches[i]
				}
			}
			return r, params
		}
	}
	return nil, nil 
}

/* SOME HELPER FUNCTIONS */
func expandPath(path string) string {
	// first extract the name of the parameters
	re := regexp.MustCompile(":([a-zA-Z0-9]+)")
	expandedPath := re.ReplaceAllString(path, "(?P<$1>[a-zA-Z0-9]+)")
	expandedPath = "^" + expandedPath + "$"
	return expandedPath
}