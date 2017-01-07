function setCookie(cname, cvalue, exdays) {
    console.log("setcookie");
    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ d.toUTCString();
    var cname = document.getElementById("registerEmail").value;
    console.log(cname);
    var cvalue = "HEJ ERIK";
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
    window.location="/startpage"
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
        setCookie(data, "", 1)
      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

