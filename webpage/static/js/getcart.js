var mysql = require('mysql');

{
  host     : 'localhost',
  user     : 'root',
  password : 'exoticpi',
  database : 'exoticars',
}
);

connection.connect();

var queryString = 'SELECT * FROM Cart';

connection.query(queryString, function(err, rows, fields) {
    if (err) throw err;

    for (var i in rows) {
        console.log('Post Titles: ', rows[i].idProducts);
    }
});

connection.end();

// Puts the data from the DB into the rigth place on page
function DisplayCartData(data){

	Product_Id = document.getElementById('Product_Id');
	TotalPrice = document.getElementById('TotalPrice');
	Quantity = document.getElementById('Quantity');

	if(document.getElementById("Product_Id") != null){
    	Product_Id.innerHTML = data.idProducts;
	}
	if(document.getElementById("TotalPrice") != null){
    	TotalPrice.innerHTML = data.TotalPrice;
	}
	if(document.getElementById("Quantity") != null){
    	Quantity.innerHTML = data.Quantity;
	}
}
