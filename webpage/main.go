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
	"strconv"

	// Third party packages not using
	//"github.com/julienschmidt/httprouter"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	//	"strconv"
	//	"golang.org/x/crypto/bcrypt"

)

type Car struct {
	IdProducts					string `json=IdProducts`
	ProductName 				string `json=ProductDescription`
	Price								string `json=Price`
	ProductDescription 	string `json=ProductDescription`
	UnitsInStock 				string `json=UnitsInStock`
	ProductAvailable 		string `json=ProductAvailable`
}

type Cart struct {
	IdProducts					string `json=IdProducts`
	Quantity 						string `json=Quantity`
	TotalPrice					string `json=TotalPrice`
	ProductName					string `json=ProductName`
}

type Orders struct {
	IdOrders						int `json=IdOrders`
	Sent 								int `json=Sent`
	Paid 								int `json=Paid`
	PaymentType					string `json=PaymentType`
	Details							string `json=Details`
}

type Review struct {
	Rating							int `json=Rating`
	Review							string `json=Review`
}

type User struct {
	IdCustomers						int64 `json=IdCustomers`
}

var db *sql.DB
var err error

func registerHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Register new User")
	result := req.URL.RequestURI()
	//substring[2] contains the username
	//substring[3] contains the password
	substring := strings.Split(result,"/")
	Email := substring[2]
	password := substring[3]

	/* Email := req.FormValue("registerEmail")
	password := req.FormValue("registerpassword") */

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

		if err == nil {
			http.Error(res, "Server error, unable to create your account.", 500)
						http.Redirect(res, req, "/login", 301)
        } else {

        r, err := db.Exec("INSERT INTO Customers(Email, password) VALUES(?, ?)", Email, password)
        if err != nil {
						http.Error(res, "Server error, unable to create your account.", 500)
            return
        }
				id, err := r.LastInsertId()
				if err != nil {
				            println("Error:", err.Error())
				        } else {
				            println("LastInsertId:", id)
										log.Printf("User added to DB")
										return
				        }
				}

    defer db.Close()}

func authHandler(res http.ResponseWriter, req *http.Request)  {
	log.Printf("Authenticating that the User exist in DB")
	result := req.URL.RequestURI()
	//substring[2] contains the username
	//substring[3] contains the password
	substring := strings.Split(result,"/")
	Email := substring[2]
	password := substring[3]

    // Grab from the database
    var databaseUsername  string
    var databasePassword  string
    var Admin string
    var idCustomers int


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
    err := db.QueryRow("SELECT Email, idCustomers, Password, Admin FROM Customers WHERE Email=?", Email).Scan(&databaseUsername, &idCustomers, &databasePassword, &Admin)
	fmt.Println("hello", Email)
	if err == nil {
    		if (Email == databaseUsername && password == databasePassword){
					err = db.QueryRow("SELECT idCustomers FROM Customers WHERE Email=?", Email).Scan(&idCustomers)
					if err != nil {
							http.Error(res, "Server error, unable to create your account.", 500)
							return
					}
    			if (Admin == "1"){
						log.Printf("User exist, Superadmin!")
						user := &User{}
						i64 := int64(idCustomers)
						user.IdCustomers = i64
						userdetails,_ := json.Marshal(user)
						res.Write(userdetails)
    			} else {
						fmt.Println("ID ", idCustomers)
        		user := &User{}
						i64 := int64(idCustomers)
						user.IdCustomers = i64
						userdetails,_ := json.Marshal(user)

						log.Printf("User exist in DB, sent back userdetails and set cookie")
        		res.Write(userdetails)
        		//	http.Redirect(res, req, "/startpage", 301)

        		}
        	} else{
        			http.Redirect(res,req,"/login",301)
        	}
    } else{
        		http.Redirect(res,req,"/login",301)
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

	pageTemplate := "static/templates/index_loggedin.html"

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
	car.IdProducts = idProducts
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
    var idProducts, Quantity, TotalPrice, ProductName string

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

				err = db.QueryRow("SELECT ProductName FROM Products WHERE idProducts=?", idProducts).Scan(&ProductName)

				cart.IdProducts = idProducts
				cart.Quantity = Quantity
				cart.TotalPrice = TotalPrice
				cart.ProductName = ProductName

				Cart_result = append(Cart_result, *cart)
		}
		log.Printf("1", Cart_result)

	if err != nil {
		} else {

		}

	defer db.Close()

	cartdetails,_ := json.Marshal(Cart_result)
	w.Write(cartdetails)}

func removeFromCart(w http.ResponseWriter, r *http.Request) {
	result := r.URL.RequestURI()
	//substring[2] contains the CarName
	//substring[3] contains the idCustomer
	substring := strings.Split(result,"/")
	log.Printf("Time to remove car", substring[2])
	log.Printf("From cart", substring[3])

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
    // Search the database for the ProductId provided
    // Delete from cart

		err := db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts)
		if err != nil {
				panic(err.Error())
		}

		_, err = db.Exec("DELETE FROM Cart WHERE idProducts=? AND idCustomers=?", idProducts, substring[3])
	if err != nil {
		} else {

		}

	defer db.Close()}

func addToCart(w http.ResponseWriter, r *http.Request) {
		result := r.URL.RequestURI()
		//substring[3] contains the customerId
		//substring[2] contains the ProductName
		substring := strings.Split(result,"/")

		   // Grab from the database
	    var idProducts, Price, ProductAvailable string
			var UnitsInStock, Quantity int
			var newQuantity = 0
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

			err = db.QueryRow("SELECT idProducts, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &UnitsInStock, &ProductAvailable)
			err = db.QueryRow("SELECT Quantity FROM Cart WHERE idProducts=? AND idCustomers=?", idProducts, substring[3]).Scan(&Quantity)

			var canIget1More = Quantity + 1
			var newUnitInStock = UnitsInStock - canIget1More
			log.Printf("new unit in stock ", newUnitInStock)

			if UnitsInStock == 1 {
				if Quantity >= 1 {
					log.Printf("Can't buy one more")
					http.Redirect(w,r,"/showroom/"+substring[2],301)
				} else {
				newQuantity = 1
				err := db.QueryRow("SELECT idProducts, Price, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &Price, &UnitsInStock, &ProductAvailable)
				log.Printf("First time inserting ")
				_, err = db.Exec("INSERT INTO Cart(idCustomers, idProducts, Quantity, TotalPrice) VALUES(?, ?, ?, ?)", substring[3], idProducts, newQuantity, Price)

				if err != nil {
						panic(err.Error())
				}
				}
			} else if newUnitInStock >= 0 {

				newQuantity = Quantity + 1

				// Insert to cart
				err := db.QueryRow("SELECT idProducts, Price, UnitsInStock, ProductAvailable FROM Products WHERE ProductName=?", substring[2]).Scan(&idProducts, &Price, &UnitsInStock, &ProductAvailable)
				_, err = db.Exec("DELETE FROM Cart WHERE idCustomers=? AND idProducts=?", substring[3], idProducts)
				log.Printf("update cart")
				_, err = db.Exec("INSERT INTO Cart(idCustomers, idProducts, Quantity, TotalPrice) VALUES(?, ?, ?, ?)", substring[3], idProducts, newQuantity, Price)
				if err != nil {
						panic(err.Error())
				}

				} else {
					http.Redirect(w,r,"/showroom/"+substring[2],301)

				}

		defer db.Close()}

func sendOrder(w http.ResponseWriter, r *http.Request) {
		log.Printf("Placing order")

		userId := r.FormValue("order_userId")

		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("id_email")
		name := r.FormValue("id_name")
		address := r.FormValue("id_address_line")
		city := r.FormValue("id_city")
		postalcode := r.FormValue("id_postalcode")
		phone := r.FormValue("id_phone")

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

	    var idProducts int
	    var Quantity int
	    var TotalPrice int
			var UnitsInStock int
			var ProductAvailable int
	    var ProductName string
			var PaymentType = "Invoice"
			var newIdPayment int
			var databaseUsername  string
			var databasePassword  string
			var NewestOrderID int

			err := db.QueryRow("SELECT Email, Password FROM Customers WHERE idCustomers=?", userId).Scan(&databaseUsername, &databasePassword)

			if err == nil {
					if (username == databaseUsername && password == databasePassword){
						log.Printf("Confimed info correct")

						err = db.QueryRow("SELECT idOrders FROM Orders WHERE idOrders = (SELECT MAX(idOrders) FROM Orders)").Scan(&NewestOrderID)

						if err != nil {
							NewestOrderID = 0
						}

						if NewestOrderID != 0 {
							newIdPayment = NewestOrderID + 1
						} else {
							newIdPayment = NewestOrderID + 1
						}

						_, err = db.Exec("INSERT INTO Orders(idOrders, idPayment, idCustomers, Email, Fullname, Address, City, Postalcode, Phone) VALUES(?,?,?,?,?,?,?,?,?)", newIdPayment, newIdPayment, userId, email, name, address, city, postalcode, phone)
						_, err = db.Exec("INSERT INTO Payment(idPayment, PaymentType) VALUES(?,?)", newIdPayment, PaymentType)

						rows, err := db.Query("SELECT idProducts, Quantity, TotalPrice FROM Cart WHERE idCustomers=?", userId)

						if err != nil {
							panic(err.Error())
						}

						for rows.Next() {
							  log.Printf("Insert products in cart to orderdetail for user ", userId)
								err := rows.Scan(&idProducts, &Quantity, &TotalPrice)

								if err != nil {
									panic(err.Error())
								}
								err = db.QueryRow("SELECT ProductName, UnitsInStock, ProductAvailable FROM Products WHERE idProducts=?", idProducts).Scan(&ProductName, &UnitsInStock, &ProductAvailable)
								if ProductAvailable == 0 {
									_, err = db.Exec("DELETE FROM Orders WHERE idOrders=?", newIdPayment)
									_, err = db.Exec("DELETE FROM Cart WHERE idCustomers=? AND idProducts=?", userId, idProducts)
									http.Redirect(w,r,"/checkout",301)

								} else {
								_, err = db.Exec("INSERT INTO OrderDetails(idOrders, idProducts, ProductName, Quantity, Price) VALUES(?,?,?,?,?)", newIdPayment, idProducts, ProductName, Quantity, TotalPrice)

								if err != nil {
									panic(err.Error())
								}
							}
						}

						rows2, err := db.Query("SELECT idProducts, Quantity, TotalPrice FROM Cart WHERE idCustomers=?", userId)

						if err != nil {
							panic(err.Error())
						}

						for rows2.Next() {
								err := rows2.Scan(&idProducts, &Quantity, &TotalPrice)

								if err != nil {
									panic(err.Error())
								}
								err = db.QueryRow("SELECT ProductName, UnitsInStock, ProductAvailable FROM Products WHERE idProducts=?", idProducts).Scan(&ProductName, &UnitsInStock, &ProductAvailable)

								var updatedQuantity = UnitsInStock - Quantity
								log.Printf("Update UnitsInStock")
								_, err = db.Exec("update Products set UnitsInStock=? where idProducts=?", updatedQuantity, idProducts)

								if updatedQuantity == 0 {
									log.Printf("No more of that cartype left. Set ProductAvailable to 0")
									_, err = db.Exec("update Products set ProductAvailable=? where idProducts=?", 0, idProducts)
								}

								if err != nil {
									panic(err.Error())
								}
						}

						log.Printf("Order Added. Empty cart")

						_, err = db.Exec("DELETE FROM Cart WHERE idCustomers=?", userId)

					}
				}

				http.Redirect(w,r,"/startpage",301)

			defer db.Close()}

func addReview(w http.ResponseWriter, req *http.Request) {

	Rating := req.FormValue("Rating")
	Review := req.FormValue("Review")
	carmodel := req.FormValue("cartype")
	userId := req.FormValue("userId")

	var idProducts int
	var idcustomerexists int

	db, err = sql.Open("mysql", "pi:exoticpi@/mydb")
	if err != nil {
			panic(err.Error())
	}

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
			panic(err.Error())
	}

	err = db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", carmodel).Scan(&idProducts)
	err = db.QueryRow("SELECT idCustomers FROM Review WHERE idProducts=?", idProducts).Scan(&idcustomerexists)

	ConvertedIdCustomers, err := strconv.Atoi(userId)

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	if (idcustomerexists != ConvertedIdCustomers) && (userId != "") {

					err := db.QueryRow("SELECT idProducts FROM Products WHERE ProductName=?", carmodel).Scan(&idProducts)
					_, err = db.Exec("INSERT INTO Review(idCustomers, idProducts, Rating, Review) VALUES(?, ?, ?, ?)", userId, idProducts, Rating, Review)
						http.Redirect(w,req,"/showroom/"+carmodel,301)
				if err != nil {
					} else {

					}

				defer db.Close()
	} else if (idcustomerexists == ConvertedIdCustomers) && (userId != "") {
		http.Redirect(w,req,"/showroom/"+carmodel,301)
	} else{
		http.Redirect(w,req,"/login",301)
	}}

func getReview(w http.ResponseWriter, r *http.Request) {

	result := r.URL.RequestURI()
	//substring[2] contains the car name
	substring := strings.Split(result,"/")

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

				if err != nil {
					} else {

					}

				defer db.Close()
				reviewdetails,_ := json.Marshal(Review_result)
				w.Write(reviewdetails)}

func getAll(w http.ResponseWriter, r *http.Request) {

								  // Grab everything from the database

									var Orders_result []Orders // create an array of Orders
							    var idOrders, Sent, Paid, Quantity, Price int
									var PaymentType, ProductName string
									var OrderInfo = ""

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

									rows, err := db.Query("SELECT idOrders, Sent, Paid FROM Orders")

									for rows.Next() {

									    Orders := &Orders{}
											err := rows.Scan(&idOrders, &Sent, &Paid)

											rows2, err := db.Query("SELECT ProductName, Quantity, Price FROM OrderDetails WHERE idOrders=?", idOrders)

											if err != nil {
												panic(err.Error())
											}
											
											for rows2.Next() {
												err := rows.Scan(&ProductName, &Quantity, &Price)
												sQuantity := strconv.Itoa(Quantity)
												sPrice := strconv.Itoa(Price)
												OrderInfo += OrderInfo + " ProductName: " + ProductName + ", Quantity: " + sQuantity + ", Price: " + sPrice + "; "
											}

											log.Printf("OrderInfo", OrderInfo)

											Orders.IdOrders = idOrders
											Orders.Details = OrderInfo
											Orders.Sent = Sent
											Orders.Paid = Paid

											if err != nil {
												panic(err.Error())
											}

											err = db.QueryRow("SELECT PaymentType FROM Payment WHERE idPayment=?", idOrders).Scan(&PaymentType)

											if err != nil {
												panic(err.Error())
											}

											Orders.PaymentType = PaymentType

											Orders_result = append(Orders_result, *Orders)
									}
								log.Printf("1", Orders_result)
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

	// update
  stmt, err := db.Prepare("update Orders set Sent=? where idOrders=?")

	if err != nil {
			panic(err.Error())
	}

	_, err = stmt.Exec(1, substring[2])

	if err != nil {
			panic(err.Error())
	}

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

func main() {

	// Instantiate a new router

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	/* Real LAN address for server */
	bindAddr := "192.168.1.242:8080"

  //Mox Address
	//bindAddr := "130.240.110.93:8000"

	/* Handlers for differnt pages and call functions */
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/startpage", loggedinHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/adminpage", adminPageHandler)
	http.HandleFunc("/checkout", checkoutHandler)
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/register/", registerHandler)

	/* Handler for the Cars. If logged in or not */
	http.HandleFunc("/showroom/ferrari", showroomHandler)
  http.HandleFunc("/showroom_nologin/ferrari", showroom_nologinHandler)

	http.HandleFunc("/showroom/mustang", showroomHandler)
	http.HandleFunc("/showroom_nologin/mustang", showroom_nologinHandler)

	http.HandleFunc("/showroom/charger", showroomHandler)
	http.HandleFunc("/showroom_nologin/charger", showroom_nologinHandler)

	http.HandleFunc("/showroom/camaro", showroomHandler)
	http.HandleFunc("/showroom_nologin/camaro", showroom_nologinHandler)

	/* GET FUNCTIONS */
	http.HandleFunc("/car/", getCar)
	http.HandleFunc("/cart/", getCart)
	http.HandleFunc("/addToCart/", addToCart)
	http.HandleFunc("/removeFromCart/", removeFromCart)

	/* Place Order */
	http.HandleFunc("/done/", sendOrder)

	/* For Admin */
	http.HandleFunc("/everything", getAll)
	http.HandleFunc("/update/", updateDB)

	/* For Review */
	http.HandleFunc("/getReview/", getReview)
	http.HandleFunc("/addReview/", addReview)

	fmt.Println("Server running on", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, nil))}
