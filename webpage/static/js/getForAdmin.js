
function getAllData() {
    //Type is the users id that is saved in the session. carmodel is the car that is added to the cart
    var xhr = typeof XMLHttpRequest != 'undefined'
      ? new XMLHttpRequest()
      : new ActiveXObject('Microsoft.XMLHTTP');
    xhr.open('get',"/everything", true);
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
          //Sends user to cart
          window.onload = DisplayAllData(data)

        } else {
          console.log("error")
        }
      }
    };
    xhr.send();
}

// Put all data on the admin page Orders produkt etc.
function DisplayAllData(data){

  for( var i=0, l=data.length; i<l; i++ ) {

    var b = document.createElement("b");
    b.id = "o"+i
    var btn = document.createElement("BUTTON");
    if (data.Sent == 1) {
      IdOrders = document.getElementById('orderlist').appendChild(b);

      var t = document.createTextNode("Sent");
      btn.appendChild(t);
      document.getElementById('o'+i).appendChild(btn);

    } else {
      IdOrders = document.getElementById('orderlist').appendChild(b);

      var t = document.createTextNode("Update order to sent");
      btn.appendChild(t);
      btn.href="/update/"+data.IdOrders
      document.getElementById('o'+i).appendChild(btn);

    }

  	if(document.getElementById("o"+i) != null){
      	IdOrders.innerHTML = data.IdOrders;
  	}
    document.write("\n");
 }
}
