package utils

import "fmt"
import "log"
import "net/http"
import "io"

func PrintlnAndFlush(w io.Writer, s string) {
	PrintAndFlush(w, s+"\n")
}

func PrintAndFlush(w io.Writer, s string) {
	fmt.Fprintf(w, s)	
	if f, ok := w.(http.Flusher); ok {
    	f.Flush()
	} else {
		log.Println("Damn, no flush");
  	}
}