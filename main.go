package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	addr = flag.String("addr", ":8091", "web server listening address")
	filepath = flag.String("fpath", "", "filepath to zip bomb")
	filename = flag.String("fname", "sys-memory-dump.zip", "filename used to serve zip bomb")
	delay = flag.Duration("delay", 0, "delay for the response")
	webServer = flag.String("websrv", "Apache/1.8", "web server header value")
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	flag.Parse()

	fileData, err := ioutil.ReadFile(*filepath)
	if err != nil {
		log.Fatalf("failed to read file, %s\n", err.Error())
	}

	engine := gin.Default()
	engine.Use(fileHandler(fileData))

	if err := engine.Run(*addr); err != nil {
		log.Fatal(err)
	}
}


func fileHandler(data []byte) gin.HandlerFunc {
	delayVal := *delay
	filenameVal := *filename
	webServerVal := *webServer

	return func (c *gin.Context) {
		c.Header("Server", webServerVal)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filenameVal)

		if delayVal > 0 {
			time.Sleep(delayVal)
		}

		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}