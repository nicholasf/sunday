package sunday

import(
	"errors"
	"github.com/gorilla/mux"
	"net/http"
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

	for path, controllerChain := range routes.Mappings() {
		ro.gorillaRouter.HandleFunc(path, delegate(controllerChain...))
	}

	http.Handle("/", ro.gorillaRouter)

	return Router(ro), nil
}

func (r *router) Routes() Routes {
	return r.routes
}

func delegate(chain ...Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := NewRequest(r)
		response := NewResponse()

		for _, c := range chain {
			response, err = c(request, response)
			
			if err != nil {
				Log.Error(err.Error())
				break
			}
		}
	}
}