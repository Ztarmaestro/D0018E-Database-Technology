

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
  xhr.open('get',"/car/"+type, true);
  xhr.onreadystatechange = function() {
    var status;
    var data;
    var obj;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-readystate
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      if (status == 200) {
        console.log(xhr.response)
        data = JSON.parse(xhr.response);
        obj = JSON.parse(data)
        console.log(obj)
            
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

var pathArray = window.location.pathname.split( '/showroom_nologin/' );
var pathArray_login = window.location.pathname.split( '/showroom/' );
if (pathArray[1] == 'ferrari' || pathArray_login[1] == 'ferrari'){
  getCar('ferrari');  
}
if (pathArray[1] == 'camaro'  || pathArray_login[1] == 'camaro'){
  getCar("camaro");  
}
if (pathArray[1] == 'mustang' || pathArray_login[1] == 'mustang'){
  getCar("mustang");  
}
if (pathArray[1] == 'charger' || pathArray_login[1] == 'charger'){
  getCar("charger");  
}

