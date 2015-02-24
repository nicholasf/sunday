package sunday

type Controller interface {
    Do(r Request) (Model, View, error)
}

type ControllerCreatorFunc func() (Controller, error)

type DefaultController struct {}

