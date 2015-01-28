package sunday

type Routes interface {
	Mappings() map[string][]Controller
	Get(path string, c ...Controller) error
}

type routes struct {
	mappings map[string][]Controller
}

func NewRoutes() Routes {
	r := &routes{
		mappings:  make(map[string][]Controller),
		
	}
	return Routes(r)
}


func (r *routes) Mappings() map[string][]Controller {
	return r.mappings
}

func (r *routes) Get(path string, c ...Controller) (err error) {
	r.mappings[path] = c
	return
}
