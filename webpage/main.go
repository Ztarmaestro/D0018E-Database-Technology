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
	idProducts					string `json=idProducts`
	ProductName 				string `json=ProductDescription`
	Price								string `json=Price`
	ProductDescription 	string `json=ProductDescription`
	UnitsInStock 				string `json=UnitsInStock`
	ProductAvailable 		string `json=ProductAvailable`
}

type Cart struct {
	idProducts					string `json=idProducts`
	idCustomers 				string `json=idCustomers`
	Quantity 						string `json=Quantity`
	TotalPrice					string `json=TotalPrice`
}

type Orders struct {
	idOrders						int `json=idOrders`
	Sent 								int `json=Sent`
}

type Review struct {
	Rating							int `json=Rating`
	Review							string `json=Review`
}

var db *sql.DB
var err error

func registerHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("registerHandler")

	Email := req.FormValue("Email")
	password := req.FormValue("password")

	var user string

	// Create an sql.DB and check for errors
    //db, err = sql.Open("mysql", "martin:persson@/mydb")
		db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
  }


    err := db.QueryRow("SELECT Email FROM Customers WHERE Email=?", Email).Scan(&user)
  	log.Printf("email", Email)
   	log.Printf("password", password)

    switch {
    case err == sql.ErrNoRows:
    	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err == nil {
			http.Error(res, "Server error, unable to create your account.", 500)
            return
        }
        _, err = db.Exec("INSERT INTO Customers(Email, password) VALUES(?, ?)", Email, password)
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
    defer db.Close()}

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
		db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
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
    err := db.QueryRow("SELECT Email, Password, Admin FROM Customers WHERE Email=?", Email).Scan(&databaseUsername, &databasePassword, &Admin)
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
   	}
   	// sql.DB should be long lived "defer" closes it once this function ends
    defer db.Close()

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

	pageTemplate := "static/templates/index.html"

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
	//substring[2] contains the car name
	substring := strings.Split(result,"/")

	   // Grab from the database
    var idProducts, ProductName, Price, ProductDescription, UnitsInStock, ProductAvailable string

    // Create an sql.DB and check for errors
		//db, err = sql.Open("mysql", "martin:persson@/mydb")
		db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
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
    err := db.QueryRow("SELECT idProducts, ProductName, Price, ProductDescription, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &ProductName, &Price, &ProductDescription, &UnitsInStock, &ProductAvailable)
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
	w.Write(cardetails)}

func getCart(w http.ResponseWriter, r *http.Request) {
	result := r.URL.RequestURI()
	//substring[2] contains the customerId
	substring := strings.Split(result,"/")

	  // Grab from the database
		var Cart_result []Cart // create an array of Cart
    var idProducts, Quantity, TotalPrice string

    // Create an sql.DB and check for errors
		//db, err = sql.Open("mysql", "martin:persson@/mydb")
		db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
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
		// Need to grab every row with this id and add it to the struct in some way
    //err := db.QueryRow("SELECT idProducts, Quantity, TotalPrice FROM Cart WHERE idCustomers=?", substring[2]).Scan(&idProducts, &Quantity, &TotalPrice)

		rows, err := db.Query("SELECT idProducts, Quantity, TotalPrice FROM Cart WHERE idCustomers=?", substring[2])

		for rows.Next() {
		    cart := &Cart{}
				err := rows.Scan(&idProducts, &Quantity, &TotalPrice)

				if err != nil {
					panic(err.Error())
				}

				cart.idProducts = idProducts
				cart.Quantity = Quantity
				cart.TotalPrice = TotalPrice

				Cart_result = append(Cart_result, *cart)
		}

	if err != nil {
		} else {

		}

	defer db.Close()

	cartdetails,_ := json.Marshal(Cart_result)
	w.Write(cartdetails)}

func removeFromCart(w http.ResponseWriter, r *http.Request) {
	result := r.URL.RequestURI()
	//substring[2] contains the CarName
	//substring[3] contains the idProducts
	substring := strings.Split(result,"/")
	log.Printf(substring[2])
	log.Printf(substring[3])

    // Create an sql.DB and check for errors
		//db, err = sql.Open("mysql", "martin:persson@/mydb")
		db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
    if err != nil {
        panic(err.Error())
    }

    // Test the connection to the database
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    // Search the database for the ProductName provided
    // Delete from cart

		_, err = db.Exec("DELETE FROM Cart WHERE idProducts=? AND ProductName=?", substring[3], substring[2])

	if err != nil {
		} else {

		}

	defer db.Close()}

func addToCart(w http.ResponseWriter, r *http.Request) {
		result := r.URL.RequestURI()
		//substring[3] contains the customerId
		//substring[2] contains the ProductName
		substring := strings.Split(result,"/")
		log.Printf(substring[3])
		log.Printf(substring[2])

		   // Grab from the database
	    var idProducts, Price, UnitsInStock, ProductAvailable string
			var Quantity int

	    // Create an sql.DB and check for errors
			//db, err = sql.Open("mysql", "martin:persson@/mydb")
			db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
	    if err != nil {
	        panic(err.Error())
	    }

	    // Test the connection to the database
	    err = db.Ping()
	    if err != nil {
	        panic(err.Error())
	    }
	    // Search the database for the ProductName provided

			err = db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts)
			err = db.QueryRow("SELECT Quantity FROM Cart WHERE idProducts=? AND idCustomers=?", idProducts, substring[3]).Scan(&Quantity)

			if Quantity > 0 {

				var newQuantity = Quantity + 1

				// Insert to cart
				err := db.QueryRow("SELECT idProducts, Price, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &Price, &UnitsInStock, &ProductAvailable)
				_, err = db.Exec("DELETE FROM Cart WHERE idCustomers=? AND idProducts=?", substring[3], idProducts)
				log.Printf("update cart")
				_, err = db.Exec("INSERT INTO Cart(idCustomers, idProducts, Quantity, TotalPrice) VALUES(?, ?, ?, ?)", substring[3], idProducts, newQuantity, Price)
				if err != nil {
						panic(err.Error())
				}

				} else {
					err = db.QueryRow("SELECT idProducts, Price, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &Price, &UnitsInStock, &ProductAvailable)
					_, err = db.Exec("INSERT INTO Cart(idCustomers, idProducts, Quantity, TotalPrice) VALUES(?, ?, ?, ?)", substring[3], idProducts, 1, Price)
					log.Printf("First time inserting")

				}

		defer db.Close()}

/*func sendOrder(w http.ResponseWriter, r *http.Request)  {
		log.Printf("sendHandler")
		result := r.URL.RequestURI()
		//substring[2] contains the customerId
		substring := strings.Split(result,"/")

		// Take the cart for this customer and create an order.
		//	 Also remove what was bought from the DB and clear the cart

	    // Grab the address, card, etc from the submitted post form
	    Email := r.FormValue("Email")

	    // Grab from the database
	    var databaseUsername  string

	    // Create an sql.DB and check for errors
			db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
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
	    err := db.QueryRow("SELECT Email, Password, Admin FROM Customers WHERE Email=?", Email).Scan(&databaseUsername, &databasePassword, &Admin)
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
	   	}
	   	// sql.DB should be long lived "defer" closes it once this function ends
	    defer db.Close()}
*/

/* func addReview(w http.ResponseWriter, r *http.Request) {

	result := r.URL.RequestURI()
	//substring[2] contains the Rating
	//substring[3] contains the Review
	//substring[4] contains the customerId
	//substring[5] contains the ProductName
	substring := strings.Split(result,"/")

				  // Grab everything from the database
					var idProducts string

			    // Create an sql.DB and check for errors
					//db, err = sql.Open("mysql", "martin:persson@/mydb")
					db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
			    if err != nil {
			        panic(err.Error())
			    }

			    // Test the connection to the database
			    err = db.Ping()
			    if err != nil {
			        panic(err.Error())
			    }
					err := db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", substring[5]).Scan(&idProducts)
					_, err = db.Exec("INSERT INTO Review(idCustomers, idProducts, Rating, Review) VALUES(?, ?, ?, ?)", substring[4], idProducts, substring[2], substring[3])

				if err != nil {
					} else {

					}

				defer db.Close()} */

func getReview(w http.ResponseWriter, r *http.Request) {

	result := r.URL.RequestURI()
	//substring[2] contains the car name
	substring := strings.Split(result,"/")
	log.Printf(substring[2])

				  // Grab everything from the database

					var Review_result []Review // create an array of Orders
			    var Ratings int
					var Reviews, idProducts string

			    // Create an sql.DB and check for errors
					//db, err = sql.Open("mysql", "martin:persson@/mydb")
					db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
			    if err != nil {
			        panic(err.Error())
			    }

			    // Test the connection to the database
			    err = db.Ping()
			    if err != nil {
			        panic(err.Error())
			    }

					err := db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts)

					rows, err := db.Query("SELECT Rating, Review FROM Review WHERE idProducts=?", idProducts)

					for rows.Next() {
					    reviewlist := &Review{}
							err := rows.Scan(&Ratings, &Reviews)

							if err != nil {
								panic(err.Error())
							}

							reviewlist.Rating = Ratings
							reviewlist.Review = Reviews

							Review_result = append(Review_result, *reviewlist)
					}
				log.Printf("1", Review_result)
				if err != nil {
					} else {

					}

				defer db.Close()
				reviewdetails,_ := json.Marshal(Review_result)
				log.Printf("2", reviewdetails)
				w.Write(reviewdetails)}

func getAll(w http.ResponseWriter, r *http.Request) {

								  // Grab everything from the database

									var Orders_result []Orders // create an array of Orders
							    var idOrders, Sent int

							    // Create an sql.DB and check for errors
									//db, err = sql.Open("mysql", "martin:persson@/mydb")
									db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
							    if err != nil {
							        panic(err.Error())
							    }

							    // Test the connection to the database
							    err = db.Ping()
							    if err != nil {
							        panic(err.Error())
							    }

									rows, err := db.Query("SELECT idOrders, Sent FROM Orders")

									for rows.Next() {
									    Orders := &Orders{}
											err := rows.Scan(&idOrders, &Sent)

											if err != nil {
												panic(err.Error())
											}

											Orders.idOrders = idOrders
											Orders.Sent = Sent

											Orders_result = append(Orders_result, *Orders)
									}

								if err != nil {
									} else {

									}

								defer db.Close()

								orderdetails,_ := json.Marshal(Orders_result)
								w.Write(orderdetails)}

func updateDB(w http.ResponseWriter, r *http.Request) {

	result := r.URL.RequestURI()
	//substring[2] contains the idOrders
	substring := strings.Split(result,"/")
	log.Printf(substring[2])

	// Grab from the database
	//var idOrders, Sent int

	// Create an sql.DB and check for errors
	//db, err = sql.Open("mysql", "martin:persson@/mydb")
	db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
	if err != nil {
		 panic(err.Error())
	}

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
			panic(err.Error())
	}

	// update to sent to 1=sent

	// update
  stmt, err := db.Prepare("update Orders set Sent=? where idOrders=?")

	if err != nil {
			panic(err.Error())
	}

	_, err = stmt.Exec(1, substring[2])

	if err != nil {
			panic(err.Error())
	}

	/*err := db.Exec("UPDATE Orders SET Sent=? WHERE idOrders=?", substring[2])

	if err != nil {
			panic(err.Error())
	}*/

	defer db.Close()}

func showroomHandler(w http.ResponseWriter, r *http.Request) {

  // you access the cached templates with the defined name, not the filename

  pagePath := "static/templates/navbar_logout.html"

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

func errorHandler(w http.ResponseWriter, r *http.Request)  {

	// you access the cached templates with the defined name, not the filename

	pagePath := "static/templates/error.html"

	if t, err := template.ParseFiles(pagePath); err != nil {
		// Something gnarly happened.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// return to client via t.Execute
		t.Execute(w, nil)}}

func main() {

	// Instantiate a new router

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Real address for server, change back before pushing to git
	bindAddr := "192.168.1.242:8080"

	//Address for testing server on LAN
	//bindAddr := "127.0.0.1:8000"

  //Mox Address
	//bindAddr := "130.240.110.93:8000"

	//Handlers for differnt pages
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/startpage", loggedinHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/admin_login", adminLoginHandler)
  http.HandleFunc("/adminpage", adminPageHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/error", errorHandler)

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
	http.HandleFunc("/cart/", getCart)
	http.HandleFunc("/addToCart/", addToCart)
	http.HandleFunc("/removeFromCart/", removeFromCart)

	/* sendOrder, clean up everything
	http.HandleFunc("/done/", sendOrder)*/

	/* For Admin */
	http.HandleFunc("/everything", getAll)
	http.HandleFunc("/update/", updateDB)

	/* For review */
	http.HandleFunc("/getReview/", getReview)
	/* http.HandleFunc("/addReview", addReview) */

	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, nil))}
