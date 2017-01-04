

// Puts the data from the DB into the rigth place on page
function DisplayCarData(data){

  Product_name = document.getElementById('Name');
  Price = document.getElementById('Price');
  Description = document.getElementById('Description');
  UnitsInStock = document.getElementById('UnitsInStock');

  if(data.ProductAvailable >= 1){

    if(document.getElementById("Name") != null){
        Product_name.innerHTML = data.ProductName;
    }
    if(document.getElementById("Price") != null){
        Price.innerHTML = "$"+data.Price;
    }
    if(document.getElementById("Description") != null){
        Description.innerHTML = data.ProductDescription;
    }
    if(document.getElementById("UnitsInStock") != null){
        UnitsInStock.innerHTML = data.UnitsInStock;
    }
  } else {

    if(document.getElementById("Name") != null){
        Product_name.innerHTML = data.ProductName;
    }
    if(document.getElementById("Price") != null){
        Price.innerHTML = "----";
        document.getElementById('BuyButton').href="/error"
    }
    if(document.getElementById("Description") != null){
        Description.innerHTML = data.ProductDescription;
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
        data = JSON.parse(xhr.response);
       // obj = JSON.parse(data)
        console.log(data)
        DisplayCarData(data)

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

window.onload = function() {
var pathArray = window.location.pathname.split( '/showroom_nologin/' );
var pathArray_login = window.location.pathname.split( '/showroom/' );
if (pathArray[1] == 'ferrari' || pathArray_login[1] == 'ferrari'){
  getCar('ferrari');
  getReview('ferrari');
}
if (pathArray[1] == 'camaro'  || pathArray_login[1] == 'camaro'){
  getCar("camaro");
  getReview('camaro');
}
if (pathArray[1] == 'mustang' || pathArray_login[1] == 'mustang'){
  getCar("mustang");
  getReview('mustang');
}
if (pathArray[1] == 'charger' || pathArray_login[1] == 'charger'){
  getCar("charger");
  getReview('charger');
}
}

function getReview(type) {

  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get',"/getReview/"+type, true);
  xhr.onreadystatechange = function() {
    var status;
    var data;
    var obj;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-readystate
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      if (status == 200) {
        data = JSON.parse(xhr.response);
       // obj = JSON.parse(data)
        console.log(data)
        DisplayReviewData(data)

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

function DisplayReviewData(data){

    for( var i=0, l=data.length; i<l; i++ ) {
          var pr = document.createElement("b");
          pr.id = "r"+i;

          document.getElementById('Product_review').appendChild(pr)
          document.getElementById('r'+i).innerHTML = "Rating: " + data[i].Rating + "/5 " + "\n" + "Review: " + data[i].Review;
          var mybr = document.createElement('br');
          document.getElementById('r'+i).appendChild(mybr);
    }
}

function check_info(){

  // get the users customerid

  var customerId = "3"
  document.getElementById('userId').value = customerId;

  window.location.replace("/addReview/");
  return false;

}
