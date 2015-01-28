package sunday

import(
	"os"
	
)

type Application interface {
	Routes() Routes
	SetRoutes(r Routes) error
	Port() int
	SetPort(port int) error
	SetStaticFilesDir(path string) error
	StaticFilesDir() string
	Logger() Logger
	SetLogger(l Logger) error
	Env() string
	SetEnv(env string) error
}

type app struct {
	env string
	opts []func(Application) error
	port int
	routes Routes
	staticFilesDir string
}

//NewApp creates an App and returns it as an Application
func NewApplication(r Routes) (ap Application, err error) {
	a := &app{
		port: 7000,
		staticFilesDir: "./public",
		routes: r,
	}
	
	Log = NewLogger()
	
	ap = Application(a)
	return ap, nil
}

//Routes struct for the Application
func (a *app) Routes() Routes {
	return a.routes
}

//Set Routes
func (a *app) SetRoutes(r Routes) (e error) {
	a.routes = r
	return
}

//Returns the env
func (a *app) Env() string {
	if len(a.env) > 0 {
		return a.env
		
	}
	
	a.env = os.Getenv("SUNDAY_ENV")

	if len(a.env) == 0 {
		a.env = "development"
	}
	
	return a.env
}

func (a *app) SetEnv(env string) (e error) {
	a.env = env
	return
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

//Set the dir static files are served from 
func (a *app) SetStaticFilesDir(dirname string) (e error) {
	a.staticFilesDir = dirname
	return
}

//Returns the Application Logger
func (a *app) Logger() Logger {
	return Log
}

func (a *app) SetLogger(l Logger) (e error) {
	Log = l
	return
}