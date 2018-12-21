package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Subject string
}

type con struct {
	Name    string
	Content string
}

var (
	posts = make([]post, 0, 500)
	cons  = make([]con, 0, 1000)
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Printf("Failed to parse template, error: %v", err)
	}
	if r.Method == http.MethodPost {
		title := r.FormValue("titleName")
		posts = append(posts, post{Subject: title})
		if err := t.Execute(w, posts); err != nil {
			log.Print(err)
		}
	} else {
		if err := t.Execute(w, nil); err != nil {
			log.Printf("Failed to parsed template to the specified data object.")
		}
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	t, err := template.ParseFiles("./static" + r.URL.Path)
	if err != nil {
		log.Printf("Failed to parse template, error: %v", err)
	}
	if r.Method == http.MethodPost {
		name := r.FormValue("Username")
		comment := r.FormValue("userComment")
		tt := con{
			Name:    name,
			Content: comment,
		}
		cons = append(cons, tt)
		if err := t.Execute(w, cons); err != nil {
			log.Printf("Failed to parsed template to the specified data object.")
		}
	} else {
		if err := t.Execute(w, nil); err != nil {
			log.Print(err)
		}
	}
}

// func initAPIUser(r *mux.Router) {
// 	s := r.PathPrefix("/user").Subrouter()
// 	s.HandleFunc("/list", test1)
// 	s.HandleFunc("/add", test2)
// }

// func test1(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello this is /list router."))
// }

// func test2(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello this is /add router."))
// }
// func newAPIMux() *mux.Router {

// }

func main() {
	rtr := mux.NewRouter()
	// r := rtr.PathPrefix("/api").Subrouter()
	//分发路由
	// initAPIUser(r)
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/link/article_1.html", article)
	rtr.HandleFunc("/link/article_2.html", article)
	//配置静态路由
	fs := http.FileServer(http.Dir("/Users/hapi666/html/BBS/static"))
	rtr.PathPrefix("/").Handler(fs)

	http.Handle("/", rtr)
	// http.ListenAndServe(":8080", newAPIMux())
	http.ListenAndServe(":8080", nil)
}
