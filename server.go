package sunday

import (
	"log"
//	"net/http"
//	"strconv"
)

func Run(routes Routes, opts ...func(Application) error) (app Application, err error) {
	app, err =  NewApplication()
	
	for _, opt := range opts {
		err = opt(app)
		
		if err != nil {
			log.Fatal("Error configuring application: ", err.Error())
			return
		}
	}

//	http.FileServer(http.Dir(app.StaticFilesDir()))

	if err != nil {
		log.Fatal("Couldn't start Sunday: ", err.Error())
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
		log.Fatal("Could not start HTTP.", err.Error())

	}

	log.Print("Sunday running on Port ", app.Port(), " (", app.Env(), ")")

	return
}
