package sunday

import(
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

type Router interface {
	Routes() Routes
}

type router struct {
	routes Routes
	gorillaRouter *mux.Router
}

func NewRouter(routes Routes) (r Router, e error) {
	if routes == nil {
		e = errors.New("Routes is nil.")
		return
	}

	ro := &router{
		routes: routes,
		gorillaRouter:  mux.NewRouter(),
	}

	for path, controller := range routes.Mappings() {
		ro.gorillaRouter.HandleFunc(path, route(controller))
	}

	http.Handle("/", ro.gorillaRouter)

	return Router(ro), nil
}

func (r *router) Routes() Routes {
	return r.routes
}

func route(c Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Log.Info(r.Method + ": " + r.RequestURI)
		request, err := NewRequest(r)
		model, view, err := c.Do(request)
		response, err := view.Render(model)

		if err != nil {
			Log.Error(err.Error())
		}

		fmt.Fprintf(w, string(response.Data()))
	}
}