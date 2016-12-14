function getCart(type) {
  //Type is the users id that is saved in the session

  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get',"/cart/"+type, true);
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
        DisplayCartData(data)

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

// Puts the data from the DB into the rigth place on page
function DisplayCartData(data){

  var b = document.createElement("b");

	Product_Id = document.getElementById('Product_Id').appendChild(b);
	TotalPrice = document.getElementById('TotalPrice').appendChild(b);
	Quantity = document.getElementById('Quantity').appendChild(b);

  //This is the total cost of all product. Need to calulate it somehow!!
  TotatlTotalPrice = document.getElementById('Totaltotalprice');

	if(document.getElementById("Product_Id") != null){
    	Product_Id.innerHTML = data.idProducts;
	}
	if(document.getElementById("Quantity") != null){
    	Quantity.innerHTML = data.Quantity;
	}
  if(document.getElementById("TotalPrice") != null){
      TotalPrice.innerHTML = "$"+data.TotalPrice;
  }
  if(document.getElementById("TotaltotalPrice") != null){
      TotalPrice.innerHTML = "$"+data.TotalPrice;
  }
}

getUserCart(){

  //Check if session exist and take the customerId and send with it
  //Else alert and do nothing

  var customerId = "3"

  alert("You are not logged in and should not be here! Please register or login to an account");

  //should only run if you are loggedin. Need session to check!
  getCart(customerId)

}

sendOrder(){

  //Check if session exist and take the customerId and send with it
  //Else alert and do nothing
  alert("Order Sent!");

  var customerId = "3"

  alert("You are not logged in and should not be here! Please register or login to an account");

  //should only run if you are loggedin. Need session to check!
  sendCart(customerId)

}

function sendCart(type) {
  //Type is the users id that is saved in the session

  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get',"/done/"+type, true);
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
        window.location = "/startpage";

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}