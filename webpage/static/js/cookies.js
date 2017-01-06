function setCookie(cname, cvalue, exdays) {

    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ d.toUTCString();
    var cname = document.getElementById("Email")	
    var cvalue = "hehe"
    document.cookie = cname// + "=" + cvalue + ";" + expires + ";path=/";
}



function getCookie(type) {
	var cname = document.getElementById("Email")
	var cpassword = document.getElementById("password")
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('get',"/Customers/"+cname+"/"+password, true);


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
        setCookie(data, "", 1)

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}

