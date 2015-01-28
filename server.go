package sunday

import (
//	"net/http"
	"strconv"
)

var l Logger

func Run(routes Routes, opts ...func(Application) error) (app Application, err error) {
	app, err = NewApplication(routes)
	l = app.Logger()
	
	for _, opt := range opts {
		err = opt(app)
		
		if err != nil {
			l.Fatal("Error configuring application: " + err.Error())
			return
		}
	}

	if err != nil {
		l.Fatal("Couldn't start Sunday: " + err.Error())
	}

//	router, err := routes.Router()
//
//	if err != nil {
//		log.Fatal("Could not build routes: ", err.Error())
//
//	}
//
//	http.Handle("/", router)
//	log.Print("3 - ", strconv.Itoa(app.Port()))

//	err =  http.ListenAndServe(":" + strconv.Itoa(app.Port()) , nil)

	if err != nil {
		l.Fatal("Could not start HTTP." + err.Error())

	}

	l.Info("Sunday running on Port " + strconv.Itoa(app.Port()) + " (" + app.Env() + ")")

	return
}
