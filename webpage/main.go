package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	// Third party packages
//	"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/navbar_login.html"

	pageTemplate := "static/templates/index.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func loggedinHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/navbar_logout.html"

	pageTemplate := "static/templates/index.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

	func loginHandler(w http.ResponseWriter, r *http.Request)  {

	// you access the cached templates with the defined name, not the filename

	pagePath := "static/templates/login.html"

	if t, err := template.ParseFiles(pagePath); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func adminLoginHandler(w http.ResponseWriter, r *http.Request)  {

// you access the cached templates with the defined name, not the filename

pagePath := "static/templates/adminlogin.html"

if t, err := template.ParseFiles(pagePath); err != nil {
	// Something gnarly happened.
	http.Error(w, err.Error(), http.StatusInternalServerError)
} else {
	// return to client via t.Execute
	t.Execute(w, nil)
}}

func adminPageHandler(w http.ResponseWriter, r *http.Request)  {

// you access the cached templates with the defined name, not the filename

pagePath := "static/templates/adminpage.html"

if t, err := template.ParseFiles(pagePath); err != nil {
	// Something gnarly happened.
	http.Error(w, err.Error(), http.StatusInternalServerError)
} else {
	// return to client via t.Execute
	t.Execute(w, nil)
}}

func showroomHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/navbar_login.html"

	//if p.ByName("name") == "ferrari" {
	if r.URL.Path == "/showroom_nologin/ferrari.html" {
		pageTemplate := "/static/templates/ferrari.html"

		if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
			// Something gnarly happened.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			// return to client via t.Execute
			t.Execute(w, nil)
		}
	}
}

/*	if p.ByName("name") == "charger" {
		pageTemplate := "static/templates/charger.html"

		if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
			// Something gnarly happened.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			// return to client via t.Execute
			t.Execute(w, nil)
		}
	}

	if p.ByName("name") == "camaro" {
		pageTemplate := "static/templates/camaro.html"

		if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
			// Something gnarly happened.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			// return to client via t.Execute
			t.Execute(w, nil)
		}
	}

	if p.ByName("name") == "mustang" {
		pageTemplate := "static/templates/mustang.html"

		if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
			// Something gnarly happened.
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			// return to client via t.Execute
			t.Execute(w, nil)
		}
	}
}

*/

	func showroom_nologinHandler(w http.ResponseWriter, r *http.Request) {

	  // you access the cached templates with the defined name, not the filename

	  pagePath := "static/templates/navbar_logout.html"

		//if p.ByName("name") == "ferrari" {
	  fmt.Println("hej")
		if r.URL.Path == "/showroom_nologin/ferrari" {	
			fmt.Println("Inside url path")
			pageTemplate := "static/templates/ferrari.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}
}

/*if p.ByName("name") == "charger" {
			pageTemplate := "static/templates/charger.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}

		if p.ByName("name") == "camaro" {
			pageTemplate := "static/templates/camaro.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}

		if p.ByName("name") == "mustang" {
			pageTemplate := "static/templates/mustang.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}
}
*/

func main() {
	// Instantiate a new router

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    

	//Real address for server, change back before pushing to git
	//bindAddr := "130.240.170.56:8080"

	//Address for testing server on LAN
	bindAddr := "127.0.0.1:8000"

//	r := httprouter.New()

	//Handlers for differnt pages
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/startpage", loggedinHandler)
    http.HandleFunc("/showroom/:name", showroomHandler)
   	http.HandleFunc("/showroom_nologin/:name", showroom_nologinHandler)
   	http.HandleFunc("/login", loginHandler)
   	http.HandleFunc("/admin_login", adminLoginHandler)
   	http.HandleFunc("/adminpage", adminPageHandler)
/*	r.GET("/", indexHandler)
	r.GET("/startpage", loggedinHandler)

	r.GET("/showroom/:name", showroomHandler)
	r.GET("/showroom_nologin/:name", showroom_nologinHandler)

	r.GET("/login", loginHandler)

	r.GET("/admin_login", adminLoginHandler)
	r.GET("/adminpage", adminPageHandler)
*/
	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
