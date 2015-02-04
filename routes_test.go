package sunday

//import (
//	c "github.com/smartystreets/goconvey/convey"
//	"io/ioutil"
//	"net/http"
//	"testing"
//	"log"
//)

//func TestRouting(t *testing.T)
//c.Convey("The router executes the controller", t, func() {
//	flag := "red"
//	r := NewRoutes()
//	r.Get("/", func(r Request, res Response)(re Response, e error){
//		flag = "green"
//		return res, nil
//	})
//	app, err := Run(r)
//	c.So(err, c.ShouldBeNil)
//	c.So(app, c.ShouldNotBeNil)
//})
//
