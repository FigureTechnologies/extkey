package commands

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func CmdServe() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server()
		},
	}
}

func newRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/generate", generateKeys)
	return router
}

func mkError(msg string) interface{} {
	return struct {
		Msg string `json:"message"`
	}{
		Msg: msg,
	}
}

func generateKeys(c *gin.Context) {
	hrp := c.Query("hrp")
	if hrp == "" {
		c.JSON(400, mkError("missing hrp"))
		return
	}
	hdPath := c.Query("hdPath")
	fmt.Printf("generating keys for hrp:%s hdPath:%s\n", hrp, hdPath)

	key, err := GenerateExtKey(hrp, hdPath)
	if err != nil {
		c.AbortWithError(505, err)
		return
	}

	c.JSON(200, key)
}

func server() error {
	router := newRouter()
	err := router.Run(":9000")
	if err != nil {
		return err
	}
	return nil
}
