package sunday

import(
	"net/http"
)

type Request interface {
//	Session Session
//	Cookie Cookie
	Logger() Logger
	Path() string
	Params() map[string]string
	HttpRequest() *http.Request
}

type request struct {
	logger Logger
	path string
	params map[string]string
	httpRequest *http.Request
}

func NewRequest(req *http.Request) (r Request, e error) {
	re := &request{
		httpRequest: req,
	}
	
	return Request(re), nil
}

func (r *request) Logger() Logger {
	return r.logger
}

func (r *request) Path() string {
	return r.path
}

func (r *request) Params() map[string]string {
	return r.params
}

func (r *request) HttpRequest() *http.Request {
	return r.httpRequest
}