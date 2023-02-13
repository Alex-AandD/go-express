package handler
import (
	"github.com/go-express/env"
	"net/http"
)

type FuncHandler func (w http.ResponseWriter, r *http.Request, next FuncHandler) error

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code 	int
	Err  	error
}

func (se *StatusError) Error() string {
	return se.Err.Error()
}

func (se *StatusError) Status() int {
	return se.Code
}


type Handler struct {
	Env 	*env.Env
	H 		FuncHandler
}

func NewHandler(env *env.Env, h FuncHandler) Handler {
	handler := Handler{Env: env, H: h}
	return handler
}