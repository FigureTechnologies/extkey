package commands

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var CmdServe = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		laddr, err := cmd.PersistentFlags().GetString("laddr")
		if err != nil {
			return err
		}
		return server(laddr)
	},
}

func init() {
	addFlags(CmdServe, flagLAddr)
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

	var seedBz []byte
	var err error
	seed := c.Query("seed")
	if seed != "" {
		seedBz, err = base64.URLEncoding.DecodeString(seed)
		if err != nil {
			c.JSON(400, mkError("unable to decode seed"))
			return
		}
	}

	hdPaths := c.Request.URL.Query()["hdPath"]
	key, err := GenerateExtKey(hrp, hdPaths, seedBz)
	if err != nil {
		_ = c.AbortWithError(505, err)
		return
	}

	c.JSON(200, key)
}

func server(laddr string) error {
	router := newRouter()
	err := router.Run(laddr)
	if err != nil {
		return err
	}
	return nil
}
