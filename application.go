package sunday

type Application interface {
	Routes() Routes
	Port() int
	SetStaticDir(path string)
	StaticDir() string
}

type app struct {
	routes Routes
	opts []func(Application) error
	port int
	staticFilesDir string
}

//NewApp creates an App and returns it as an Application
func NewApplication() (ap Application, err error) {
	a := &app{
		port: 4000,
		staticFilesDir: "./public",
	}
	ap = Application(a)
	return ap, nil
}

//Routes strct for the Application
func (a *app) Routes() Routes {
	return a.routes
}

//Returns the port of the Application
 func (a *app) Port() int {
	return a.port
}

//Sets the port of the Application
func (a *app) SetPort(port int) (e error) {
	a.port = port
	return
}

//Returns the dir of static files
func (a *app) StaticFilesDir() string {
	return a.staticFilesDir
}

func (a *app) SetStaticFilesDir(dirname string) (e error) {
	a.staticFilesDir = dirname
	return
}
