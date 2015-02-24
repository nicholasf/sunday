package sunday

type Response interface {
	Headers() map[string]string
	Data() []byte
	SetData(data []byte) (e error)
}

type response struct {
	headers map[string]string
	data []byte
	status string
}

func NewResponse() (r Response) {
	re := &response{ status: "200" }
	return Response(re)
}

func (r *response) Headers() map[string]string {
	return r.headers
}

func (r *response) Data() []byte {
	return r.data
}

func (r *response) SetData(d []byte) (e error) {
	r.data = d
	return
}