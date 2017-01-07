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

    var CarName = data[i].ProductName;
    var pid = document.createElement("b");
    pid.id = "p"+i;

    Product_List = document.getElementById('Product_List').appendChild(pid);

    //btn.href="/removeFromCart/"+data.ProductName; /* +"/"+userId */

    var btn = document.createElement("BUTTON");
    var t = document.createTextNode("Delete product");
    btn.id = "delete/"+data[i].ProductName;
    btn.appendChild(t);

  	if(document.getElementById("p"+i) != null){
      	Product_List.innerHTML = data[i].ProductName + ": " + data[i].Quantity + " x " + "$ " + data[i].TotalPrice + " ";
        document.getElementById('p'+i).appendChild(btn);
    }
    document.getElementById("delete/"+data[i].ProductName).addEventListener("click", function(){
      var str = this.id;
      var res = str.split("/");
      var CarName = res[1];

      deleteFromCart(CarName); });
    var mybr = document.createElement('br');
    document.getElementById('p'+i).appendChild(mybr);
 }
}

function deleteFromCart(car){
  var str = document.cookie;
  var res = str.split("=");
  var res2 = res[0].split("; ");
  var customerId = res2[1];
  console.log(customerId);
  //Type is the users id that is saved in the session. carmodel is the car that is added to the cart
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('post',"/removeFromCart/"+car+"/"+customerId, true);
  xhr.onreadystatechange = function() {
    var status;
    //var data;
    //var obj;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-readystate
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      if (status == 200) {
        //data = JSON.parse(xhr.response);
        // obj = JSON.parse(data)
        //console.log(data)
        //Sends user to cart
        window.location = "/checkout";

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

function getUserCart() {
  //Check if session exist and take the customerId and send with it
  //Else alert and do nothing

  var str = document.cookie;
  var res = str.split("=");
  var res2 = res[0].split("; ");
  var customerId = res2[1];
  console.log(customerId);

  getCart(customerId)

}

function order_check_info(){
  // get the users customerid

  var str = document.cookie;
  var res = str.split("=");
  var res2 = res[0].split("; ");
  var customerId = res2[1];
  console.log(customerId);

  document.getElementById('order_userId').value = customerId;

  return true

}
