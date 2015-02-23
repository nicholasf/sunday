package sunday

type View interface {
	Render(Request) (Response, error)
}

type ViewFunc func(Request) (Response, error)

func (vf ViewFunc) Render(r Request) (res Response, e error) {
    return vf(r)
}