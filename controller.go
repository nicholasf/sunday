package sunday

type Controller interface {
    Do(r Request) (View, error)
}

type ControllerCreatorFunc func() (Controller, error)

type DefaultController struct {}

