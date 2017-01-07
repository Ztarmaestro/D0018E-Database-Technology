
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

function checkIfLogin(carmodel) {

  var str = document.cookie;
  console.log(str);
  var res = str.split("=");
  console.log(res);
  var res2 = res[0].split("; ");
  console.log(res2);
  var customerId = res2[1];
  console.log(customerId);

  if (customerId != null) {
    addToCart(customerId, carmodel)
  } else {
    alert("You are not logged in and can't buy this product. Please register or login to an account");
  }
}
