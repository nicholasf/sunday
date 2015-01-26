package sunday

import (
	c "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
	"log"
)

func TestRun(t *testing.T) {
	
	c.Convey("The run function creates an application", t, func() {
		var r Routes 
		app, err := Run(r)
		c.So(err, c.ShouldBeNil)
		c.So(app, c.ShouldNotBeNil)
	})
	
	c.Convey("It serves static files", t, func() {
		var r Routes
		
		opt := func(a *Application) error {
			a.SetStaticDir(".")
		}
		
		app, err := Run(r, opt)
		page, err := get("http://localhost:4000/README.md")
		c.So(page, c.ShouldContainSubstring, "Sunday")
	})
}


func get(url string) (page text, err error) {
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