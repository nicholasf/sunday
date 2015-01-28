package sunday

type Response interface {
	Headers() map[string]string
	Data() []byte
}

type response struct {
	headers map[string]string
	data []byte
}

func NewResponse() (r Response) {
	re := &response{}
	return Response(re)
}

func (r *response) Headers() map[string]string {
	return r.headers
}

func (r *response) Data() []byte {
	return r.data
}