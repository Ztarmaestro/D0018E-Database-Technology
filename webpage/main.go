package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
//	"golang.org/x/crypto/bcrypt"
	// Third party packages
	//"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
)

type Car struct {
	idProducts	string `json=idProducts`
	ProductName string `json=ProductDescription`
	Price		string `json=Price`
	ProductDescription string `json=ProductDescription`
	UnitsInStock string `json=UnitsInStock`
	ProductAvailable string `json=ProductAvailable`
}

var db *sql.DB
var err error


func registerHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("registerHandler")

	Email := req.FormValue("Email")
	password := req.FormValue("password")

	var user string

	// Create an sql.DB and check for errors
    db, err = sql.Open("mysql", "martin:persson@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }


    err := db.QueryRow("SELECT Email FROM customers WHERE Email=?", Email).Scan(&user)
  	log.Printf("email", Email)
   	log.Printf("password", password)

    switch {
    case err == sql.ErrNoRows:
    	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err == nil {
			http.Error(res, "Server error, unable to create your account.", 500)
            return
        }
        _, err = db.Exec("INSERT INTO customers(Email, password) VALUES(?, ?)", Email, password)
        if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
            return
        }
       http.Redirect(res, req, "/startpage", 301)
       return
    case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
        return
    default:
    	http.Redirect(res, req, "/login", 301)
    }
    defer db.Close()


}


func authHandler(w http.ResponseWriter, r *http.Request)  {
	log.Printf("authHandler")
    // Grab the username/password from the submitted post form
    Email := r.FormValue("Email")
    password := r.FormValue("password")



    // Grab from the database
    var databaseUsername  string
    var databasePassword  string
    var Admin string


    // Create an sql.DB and check for errors
    db, err = sql.Open("mysql", "martin:persson@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    // Search the database for the username provided
    // If it exists grab the password for validation
    err := db.QueryRow("SELECT Email, Password, Admin FROM customers WHERE Email=?", Email).Scan(&databaseUsername, &databasePassword, &Admin)
	if err == nil {
    		if (Email == databaseUsername && password == databasePassword){
    			if (Admin == "1"){
    				http.Redirect(w, r, "/adminpage", 301)
    			} else {
        		http.Redirect(w, r, "/startpage", 301)
        		}
        	} else{
        			http.Redirect(w,r,"/login",301)
        	}
    } else{
        		http.Redirect(w,r,"/login",301)

    //	}
   	}
   	    // sql.DB should be long lived "defer" closes it once this function ends
    defer db.Close()

    
  /*  // Validate the password if hased
    err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
    // If wrong password redirect to the login
    if err != nil {
        http.Redirect(w, r, "/login", 301)
        return
    }

    // If the login succeeded
    w.Write([]byte("Hello " + databaseUsername))
*/

 }

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

	pageTemplate := "static/templates/login_index.html"

	if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func loginHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/login.html"

	if t, err := template.ParseFiles(pagePath); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}


func adminLoginHandler(w http.ResponseWriter, r *http.Request) {

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

func checkoutHandler(w http.ResponseWriter, r *http.Request)  {

	// you access the cached templates with the defined name, not the filename

	pagePath := "static/templates/checkout.html"

	if t, err := template.ParseFiles(pagePath); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)
	}}

func getCar(w http.ResponseWriter, r *http.Request) {
	result := r.URL.RequestURI()
	substring := strings.Split(result,"/")

	   // Grab from the database
    var idProducts, ProductName, Price, ProductDescription, UnitsInStock, ProductAvailable string

    // Create an sql.DB and check for errors
    db, err = sql.Open("mysql", "martin:persson@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    // Search the database for the username provided
    // If it exists grab the password for validation
    err := db.QueryRow("SELECT idProducts, ProductName, Price, ProductDescription, UnitsInStock, ProductAvailable FROM products WHERE ProductName=?", substring[2]).Scan(&idProducts, &ProductName, &Price, &ProductDescription, &UnitsInStock, &ProductAvailable)
	if err != nil {
		} else {

		}

	defer db.Close()
	
	car := &Car{}
	car.idProducts = idProducts
	car.ProductName = ProductName
	car.Price = Price
	car.ProductDescription = ProductDescription
	car.UnitsInStock = UnitsInStock
	car.ProductAvailable = ProductAvailable
	cardetails,_ := json.Marshal(car)
	w.Write(cardetails)
}

func showroomHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/navbar_logout.html"

	data := r.URL.Path

	// Split on slash.
    result := strings.Split(data, "/")

		if result[2] == "ferrari" {











		//TEST FÖR HÄMTA DESCRIPTION I DATABAS
				// Create an sql.DB and check for errors
/*   db, err = sql.Open("mysql", "martin:persson@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

	var ProductDescription string
	err := db.QueryRow("SELECT ProductDescription FROM products WHERE idProducts=47").Scan(&ProductDescription)

	log.Printf("Description:", ProductDescription)
	if err == nil {
		} else {
	}
	defer db.Close()
*/








			

			pageTemplate := "static/templates/ferrari.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}
		if result[2] == "mustang" {
			pageTemplate := "static/templates/mustang.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}
		if result[2] == "charger" {
			pageTemplate := "static/templates/charger.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}
		if result[2] == "camaro" {
			pageTemplate := "static/templates/camaro.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
		}}

func showroom_nologinHandler(w http.ResponseWriter, r *http.Request) {

	// you access the cached templates with the defined name, not the filename

	pagePath := "static/templates/navbar_login.html"

	data := r.URL.Path

	// Split on slash.
		result := strings.Split(data, "/")

	if result[2] == "ferrari" {
		pageTemplate := "static/templates/ferrari.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
	}
	if result[2] == "mustang" {
			pageTemplate := "static/templates/mustang.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
	}
	if result[2] == "charger" {
			pageTemplate := "static/templates/charger.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
	}
	if result[2] == "camaro" {
			pageTemplate := "static/templates/camaro.html"

			if t, err := template.ParseFiles(pagePath, pageTemplate); err != nil {
				// Something gnarly happened.
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				// return to client via t.Execute
				t.Execute(w, nil)
			}
	}}

func main() {

	// Instantiate a new router

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Real address for server, change back before pushing to git
	//bindAddr := "130.240.170.56:8080"

	//Address for testing server on LAN
	//bindAddr := "127.0.0.1:8000"
	bindAddr := "130.240.110.93:8000"

	//Handlers for differnt pages
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/startpage", loggedinHandler)
   	http.HandleFunc("/login", loginHandler)
   	http.HandleFunc("/admin_login", adminLoginHandler)
   	http.HandleFunc("/adminpage", adminPageHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/register", registerHandler)

	http.HandleFunc("/showroom/ferrari", showroomHandler)
   	http.HandleFunc("/showroom_nologin/ferrari", showroom_nologinHandler)

	http.HandleFunc("/showroom/mustang", showroomHandler)
	http.HandleFunc("/showroom_nologin/mustang", showroom_nologinHandler)

	http.HandleFunc("/showroom/charger", showroomHandler)
	http.HandleFunc("/showroom_nologin/charger", showroom_nologinHandler)

	http.HandleFunc("/showroom/camaro", showroomHandler)
	http.HandleFunc("/showroom_nologin/camaro", showroom_nologinHandler)


	// GET FUNCTIONS
	http.HandleFunc("/car/", getCar)	
	
	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, nil))}
