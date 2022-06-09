package mai

import (
	"github.com/gin-gonic/gin"
	"github.com/unidoc/unioffice/common/license"
	"github.com/workexercise2/api/text"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	err := license.SetMeteredKey(`2da620c9244fc36c9a9723eb70f05ade3b8492ce1e861b5d23b41c4f7f031aaa`)
	if err != nil {
		panic(err)
	}
}

func mai() {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/newtext", text.New)
	r.Run(":8000")
}
