package main

import "github.com/gin-gonic/gin"
import "fmt"
type res1 struct {
 Caller  string `json:"caller" binding:"required"`
}

func main() {
	r := gin.Default()
	r.POST("/offer", func(c *gin.Context) {
	var req res1
    	if err := c.BindJSON(&req); err != nil {
        return
    	}
	fmt.Println(req.Caller)
	if req.Caller == "cc3750c95c13d248e96fd595e9b4655bbc14e810054707a31b28777f8de94225" {

		c.JSON(200, gin.H{"message": "ggwp. FLAG is <HERE>"})

	} else {
			c.JSON(200, gin.H{"message": "NOPE."})

	
	}

})
r.POST("/app-check", func(c *gin.Context) {
	var req res1
    	if err := c.BindJSON(&req); err != nil {
        return
    	}
	fmt.Println(req.Caller)
	if req.Caller == "com.ctf.android.premiumplugin.findingneemo" {

		c.JSON(200, gin.H{"message": "ggwp. FLAG is <HERE>"})

	} else {
			c.JSON(200, gin.H{"message": "NOPE."})

	
	}

})
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
