package main
import(
	"net/http"
	"github.com/gin-gonic/gin"
)
func main(){
	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	r.GET("/kin",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"kon",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Run()
}