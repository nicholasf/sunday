package sunday

import (
	"net/http"
	"strconv"
)

var l Logger

func Run(routes Routes, opts ...func(Application) error) (app Application, err error) {
	Log = NewLogger()
	app, err = NewApplication(routes)

	if err != nil {
		Log.Fatal("Could not compiles routes: " + err.Error())
		return
	}
	
	for _, opt := range opts {
		err = opt(app)
		
		if err != nil {
			Log.Fatal("Error configuring application: " + err.Error())
			return
		}
	}

	if err != nil {
		Log.Fatal("Couldn't start Sunday: " + err.Error())
	}

	err = http.ListenAndServe(":" + strconv.Itoa(app.Port()) , nil)

	if err != nil {
		Log.Fatal("Could not start HTTP." + err.Error())

	}

	Log.Info("Sunday running on Port " + strconv.Itoa(app.Port()) + " (" + app.Env() + ")")

	return
}
