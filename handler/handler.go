package handler
import (
	"github.com/go-express/env"
	"net/http"
)

type Handler struct {
	Env 	*env.Env
	H http.HandlerFunc
}

func NewHandler(env *env.Env, h http.HandlerFunc) Handler {
	handler := Handler{Env: env, H: h}
	return handler
}