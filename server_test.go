package sunday

import (
	c "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"testing"
	"log"
)

func TestRun(t *testing.T) {

	c.Convey("The run function reacts gracefully to nil routes", t, func() {
		var r Routes
		_, err := Run(r)
		c.So(err, c.ShouldNotBeNil)
	})
	
	c.Convey("The run function creates an application", t, func() {
		r := NewRoutes()
		r.Get("/", func(r Request, res Response)(re Response, e error){ return res, nil })
		app, err := Run(r)
		c.So(err, c.ShouldBeNil)
		c.So(app, c.ShouldNotBeNil)
	})
	
//	c.Convey("It serves static files", t, func() {
//		var r Routes
//
//		opt := func(a Application) error {
//			a.SetStaticFilesDir(".")
//			return nil
//		}
//
//		_, err := Run(r, opt)
//		page, err := get("http://localhost:4000/README.md")
//
//		c.So(err, c.ShouldBeNil)
//		c.So(page, c.ShouldContainSubstring, "Sunday")
//	})
}

func get(url string) (page string, err error) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	
	return string(data), err
}