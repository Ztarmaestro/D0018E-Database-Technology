function setCookie(cname, cvalue, exdays) {
    console.log("setcookie");
    var cname = "idcustomer";
    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ d.toUTCString();
    console.log("Cookie name " + cname);
    console.log("Cookie value " + cvalue);
    console.log("Cookie expires " + expires);

    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    window.location="/startpage";

}

function deleteCookie(){
  //document.cookie = 'myCookie=; expires='+new Date(0).toUTCString() +'; path=/myPath/';
  //document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
  var customerId = callCookie("idcustomer");
  console.log(customerId);
  document.cookie = "idcustomer" + "=" + customerId + ";" + "expires=Thu, 01 Jan 1970 00:00:00 UTC" + ";path=/";
  //document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/http://79.136.28.205:8080/";
  window.location="/";

}

function newCustomerCookie() {
	var cname = document.getElementById("registerEmail").value;
	var cpassword = document.getElementById("registerpassword").value;
  console.log("username and password");
  console.log(cname);
  console.log(cpassword);

  //Check if first letter in cname i capital
  var nameok = checkNewUser(cname);

if (nameok.length == 0){
  window.location="/login";
} else {
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');

  xhr.open('GET', '/register/'+ cname + '/' + cpassword, true);

  xhr.onreadystatechange = function() {
    var status;
    var data;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-
    console.log(xhr.readyState)
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      console.log("status")
      console.log(status);
      if (status == 200) {
        data = JSON.parse(xhr.response);
       // obj = JSON.parse(data)
        console.log(data)
        setCookie("", data.IdCustomers, 1);
      } else {
        data = JSON.parse(xhr.response);
        console.log(data)
        console.log("error")
      }
    }
  };
  xhr.send();
 }
}

function getCookie(type) {
  console.log("getcookie");
	var cname = document.getElementById("loginEmail").value;
	var cpassword = document.getElementById("loginpassword").value;
  console.log("username and password");
  console.log(cname);
  console.log(cpassword);

  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');

  xhr.open('GET', '/auth/'+ cname + '/' + cpassword, true);

  xhr.onreadystatechange = function() {
    var status;
    var data;
    // https://xhr.spec.whatwg.org/#dom-xmlhttprequest-
    console.log(xhr.readyState)
    if (xhr.readyState == 4) { // `DONE`
      status = xhr.status;
      console.log("status")
      console.log(status);
      if (status == 200) {
        data = JSON.parse(xhr.response);
       // obj = JSON.parse(data)
        console.log(data)
        setCookie("", data.IdCustomers, 1);
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

function callCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function checkNewUser(username) {

  if(username[0].toUpperCase() == username[0])
  {
     console.log("First letter is Uppercase " + username[0]);
     return "ok";
  } else {
    console.log("First letter is not Uppercase " + username[0]);
    alert("First letter needs to be Uppercase! Try again.");
    return "";

  }
}
