var mysql = require('mysql');

{
  host     : 'localhost',
  user     : 'root',
  password : 'exoticpi',
  database : 'exoticars',
}
);

connection.connect();

var queryString = 'SELECT * FROM Car';

connection.query(queryString, function(err, rows, fields) {
    if (err) throw err;

    for (var i in rows) {
        console.log('Post Titles: ', rows[i].idProducts);
    }
});

connection.end();

// Puts the data from the DB into the rigth place on page
function DisplayCarData(data){

	Product_name = document.getElementById('Name');
	Price = document.getElementById('Price');
	Description = document.getElementById('Description');
	UnitsInStock = document.getElementById('UnitsInStock');

  if(data.available == true){

  	if(document.getElementById("Name") != null){
      	Product_name.innerHTML = data.Product_name;
  	}
  	if(document.getElementById("Price") != null){
      	Price.innerHTML = data.Price;
  	}
  	if(document.getElementById("Description") != null){
      	Description.innerHTML = data.Description;
  	}
  	if(document.getElementById("UnitsInStock") != null){
      	UnitsInStock.innerHTML = data.UnitsInStock;
  	}
  }
}
