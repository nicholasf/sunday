package sunday

type View interface {
	Render(Model) (Response, error)
}

type ViewFunc func(Model) (Response, error)

func (vf ViewFunc) Render(m Model) (res Response, e error) {
    return vf(m)
}