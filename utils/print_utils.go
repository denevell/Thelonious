package utils

import "fmt"
import "log"
import "net/http"

func PrintlnAndFlush(w http.ResponseWriter, s string) {
	PrintAndFlush(w, s+"\n")
}

func PrintAndFlush(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, s)	
	if f, ok := w.(http.Flusher); ok {
    	f.Flush()
	} else {
		log.Println("Damn, no flush");
  	}
}