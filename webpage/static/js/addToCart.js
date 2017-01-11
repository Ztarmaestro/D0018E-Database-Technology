
function addToCart(customerId, carmodel) {
  //Type is the users id that is saved in the session. carmodel is the car that is added to the cart

  console.log("customerId: " + customerId);
  console.log("carmodel: " + carmodel);

  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('post',"/addToCart/"+carmodel+"/"+customerId, true);
  xhr.onreadystatechange = function() {
    var status;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-readystate
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      if (status == 200) {

        window.location = "/checkout";

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

function checkIfLogin(carmodel) {

  var customerId = callCookie("idcustomer")
  console.log(customerId);

  if (customerId != "") {
    addToCart(customerId, carmodel)
  } else {
    alert("You are not logged in and can't buy this product. Please register or login to an account");
  }
}
