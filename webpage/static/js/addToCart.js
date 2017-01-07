
function addToCart(customerId, carmodel) {
  //Type is the users id that is saved in the session. carmodel is the car that is added to the cart
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
  var res = str.split("=");
  var customerId = res[0];

  if (customerId != null) {
    addToCart(customerId, carmodel)
  } else {
    alert("You are not logged in and can't buy this product. Please register or login to an account");
  }  
}
