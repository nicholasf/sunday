package sunday

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestApplication(t *testing.T) {
	c.Convey("Env", t, func() {
		c.Convey("defaults to 'development'", func() {
			var r Routes
			app, _ := Run(r)
			c.So(app.Env(), c.ShouldEqual, "development")
		})
		
	})
}