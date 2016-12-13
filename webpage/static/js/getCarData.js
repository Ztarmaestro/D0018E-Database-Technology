function getData(carname){

getJSON('https://79.136.28.205:8000/carname', function(data) {alert('Your public IP address is: ' + data.ip); }, function(status) {  alert('Something went wrong.');});

console.log("hej")
console.log(carnumber);
// var mysql = require('mysql'); 
 var connection = mysql({

  host     : 'localhost',
  user     : 'martin',
  password : 'persson',
  database : 'mydb'
}
);

connection.connect();

var queryString = 'SELECT * FROM Products WHERE ProductName = carname';

connection.query(queryString, function(err, rows, fields) {
        console.log('Post Titles: ', rows[0].idProducts);
        console.log('Post Titles: ', rows[0].ProductName);
        console.log('Post Titles: ', rows[0].Price);
        console.log('Post Titles: ', rows[0].ProductDescription);
        console.log('Post Titles: ', rows[0].UnitsInStock);
});

connection.end();
DisplayCarData(rows)

}

// Puts the data from the DB into the rigth place on page
function DisplayCarData(data){

  Product_name = document.getElementById('Name');
  Price = document.getElementById('Price');
  Description = document.getElementById('Description');
  UnitsInStock = document.getElementById('UnitsInStock');

  if(data[0].ProductAvailable == 1){

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
  } else {
    
    if(document.getElementById("Name") != null){
        Product_name.innerHTML = data.Product_name;
    }
    if(document.getElementById("Price") != null){
        Price.innerHTML = "----";
    }
    if(document.getElementById("Description") != null){
        Description.innerHTML = data.Description;
    }
    if(document.getElementById("UnitsInStock") != null){
        UnitsInStock.innerHTML = "0";
    }
  }
}

function getCar(type) {
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get','localhost:8000/car/'+type, true);
  xhr.onreadystatechange = function() {
    var status;
    var data;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-readystate
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      if (status == 200) {
        console.log(xhr.response)
        data = JSON.parse(xhr.responseText);
        
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}