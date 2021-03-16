package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func PrintBody(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(body))
	}
}
