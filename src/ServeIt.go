package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	msgflag := make(chan string)
	if len(os.Args) < 2 {
		log.Println("Server will start on default port 80")
		go servefiles("80", msgflag)
		<-msgflag

	} else {
		go servefiles(os.Args[1], msgflag)
		<-msgflag
	}
}

func servefiles(port string, msgflag chan string) {
	httpfs := http.FileServer(http.Dir("./"))
	http.Handle("/", httpfs)

	log.Println("Listening on : ", port)
	portTolisten := ":" + port
	err := http.ListenAndServe(portTolisten, nil)
	if err != nil {
		log.Fatal(err)
	}
	msgflag <- "File server started"
}

func startindefaultbrowser(urlval string) {

	if runtime.GOOS == "linux" {
		err := exec.Command("xdg-open", urlval).Start()
		log.Fatal(err)
	} else if runtime.GOOS == "windows" {
		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", urlval).Start()
		log.Fatal(err)
	}
}
