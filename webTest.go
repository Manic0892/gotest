package main

import (
    "net/http"
    "fmt"
)

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
type Links struct {
    String string
    Struct string
}

func (s Links) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprintf(w, s.String+"<br>"+s.Struct)
}

func (s String) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP(w http.ResponseWriter,r *http.Request) {
    fmt.Fprint(w, s.Greeting+s.Punct+s.Who)
}

func main() {
    // your http.Handle calls here
    http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
    http.Handle("/", &Links{"<a href=\"/string\">String</a>", "<a href=\"/struct\">Struct</a>"})
    http.ListenAndServe("localhost:4000", nil)
}