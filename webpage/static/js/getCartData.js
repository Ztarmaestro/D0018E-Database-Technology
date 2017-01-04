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
function DisplayCartData(data) {

  // get customerId from session and save to userId
  //var userId = session.customerId

  for( var i=0, l=data.length; i<l; i++ ) {

    var pid = document.createElement("b");
    pid.id = "p"+i;
    var pq = document.createElement("b");
    pq.id = "q"+i;
    var ptp = document.createElement("b");
    ptp.id = "tp"+i;
    var pttp = document.createElement("b");
    pttp.id = "ttp"+i;

    Product_Id = document.getElementById('Product_Id').appendChild(pid);
    Quantity = document.getElementById('Quantity').appendChild(pq);
    TotalPrice = document.getElementById('TotalPrice').appendChild(ptp);

    var btn = document.createElement("BUTTON");
    var t = document.createTextNode("Delete product");
    btn.appendChild(t);
    btn.href="/removeFromCart/"+data.ProductName; /* +"/"+userId */
    document.getElementById('tp'+i).appendChild(btn);

  	if(document.getElementById("p"+i) != null){
      	Product_Id.innerHTML = data[i].ProductName;
  	}
  	if(document.getElementById("q"+i) != null){
      	Quantity.innerHTML = data[i].Quantity + " x ";
  	}
    if(document.getElementById("tp"+i) != null){
        TotalPrice.innerHTML = "$ "+data[i].TotalPrice;
    }
    var mybr = document.createElement('br');
    document.getElementById('p'+i).appendChild(mybr);
    document.getElementById('q'+i).appendChild(mybr);
    document.getElementById('tp'+i).appendChild(mybr);
 }
}

function getUserCart() {

  //Check if session exist and take the customerId and send with it
  //Else alert and do nothing

  var customerId = "3"

  //window.location = "/error";

  alert("You are not logged in and should not be here! Please register or login to an account");

  //should only run if you are loggedin. Need session to check!
  getCart(customerId)

}

function sendOrder() {

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
