package main

import "fmt"
import "html/template"
import "net/http"

func myhtml(w http.ResponseWriter,r * http.Request){

var data = map[string]string{

	"nama":"Agus Taufik",

}

var t, err = template.ParseFiles("test.html")
if err != nil {
fmt.Println(err.Error())

}
t.Execute(w, data)
}
func main(){
http.HandleFunc("/", myhtml)
fmt.Println("starting web server at http://localhost:1908/")
http.ListenAndServe(":1908", nil)

}