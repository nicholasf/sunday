package sunday

import (
	"net/http"
	"strconv"
)

func Run(routes Routes, opts ...func(Application) error) (app Application, err error) {
	Log = NewLogger()
	app, err = NewApplication(routes)

	if err != nil {
		Log.Fatal("Could not compiles routes: " + err.Error())
		return
	}
	
	app, err = configure(app, opts...)
	
	if err != nil {
		Log.Fatal("Couldn't start Sunday: " + err.Error())
		return
	}

    Log.Info("Sunday running on Port " + strconv.Itoa(app.Port()) + " (" + app.Env() + ")")

	err = http.ListenAndServe(":" + strconv.Itoa(app.Port()) , nil)

	if err != nil {
		Log.Fatal("Failure starting Sunday: " + err.Error())

	}

	return
}

func configure(app Application, opts ...func(Application) error) (a Application, err error) {
	for _, opt := range opts {
		err = opt(app)

		if err != nil {
			Log.Fatal("Error configuring application: " + err.Error())
			return
		}
	}
	
	a = app
	
	return
}