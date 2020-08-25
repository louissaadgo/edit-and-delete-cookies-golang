package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func serve() {
	http.HandleFunc("/", index)
	http.HandleFunc("/del", del)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func main() {
	serve()
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "visits",
			Value: "0",
			Path:  "/",
		}
	}
	num, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	num++
	c.Value = strconv.Itoa(num)
	http.SetCookie(w, c)
	io.WriteString(w, "Visits: "+c.Value)
	io.WriteString(w, "\nGo to localhost:8080/del to delete the cookie")
}

func del(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("visits")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
}
