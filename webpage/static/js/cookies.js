function setCookie(cname, cvalue, exdays) {
    console.log("setcookie");
    var d = new Date();
    var cvalue = "idcustomer";
    var expires = "expires="+ d.toUTCString();

    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    console.log(cname);
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    window.location="/startpage"
}

function deleteCookie(cname, cvalue, exdays){
  var str = document.cookie;
  console.log(str);
  var cname = str.charAt(0);

  location.reload();
    
    window.location="/"

 // console.log("setcookie");
    //var date = new Date("12/15/1990");
    //d.setDate(30);
   // var expires = "expires="+ d.toUTCString();
    //console.log(cname);
  //document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
  //document.cookie = cname + cvalue + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT" +";path=/";
  //document.cookie = cname + "=" + cvalue + ";" + "-1" + ";path=/";
   // document.cookie = cname + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';



}


function newCustomerCookie() {
	var cname = document.getElementById("registerEmail").value;
	var cpassword = document.getElementById("registerpassword").value;
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
        setCookie(data.IdCustomers, "", 1)
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();

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
        setCookie(data.IdCustomers, "", 1)
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}
