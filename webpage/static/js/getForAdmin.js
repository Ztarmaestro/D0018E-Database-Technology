
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
          DisplayAllData(data)

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

    var orderid = data[i].IdOrders;
    var orderlist = document.createElement("b");
    orderlist.id = "o"+i;
    var btn = document.createElement("BUTTON");

    if (data[i].Sent == '1') {
      IdOrders = document.getElementById('orderlist').appendChild(orderlist);
      var t = document.createTextNode("Sent");
      btn.appendChild(t);
      IdOrders.innerHTML = data[i].IdOrders + " | PaymentType: " + data[i].PaymentType + " ";
      document.getElementById('o'+i).appendChild(btn);
    } else {
      if (data[i].Paid == 1){
        var isPaid = "Paid";
      } else {
        var isPaid = "Not Paid";
      }
      IdOrders = document.getElementById('orderlist').appendChild(orderlist);
      var t = document.createTextNode("Click to update order to sent");
      btn.appendChild(t);
      btn.id = "send"+i;
      IdOrders.innerHTML = "OrderId: " + data[i].IdOrders + " | " + paid + " | PaymentType: " + data[i].PaymentType + " ";
      document.getElementById('o'+i).appendChild(btn);
      document.getElementById("send"+i).addEventListener("click", function(){ updateOrder(orderid); });
    }
    var mybr = document.createElement('br');
    document.getElementById('o'+i).appendChild(mybr);
 }
}

function updateOrder(orderid){
  //Type is the users id that is saved in the session. carmodel is the car that is added to the cart
  var xhr = typeof XMLHttpRequest != 'undefined'
    ? new XMLHttpRequest()
    : new ActiveXObject('Microsoft.XMLHTTP');
  xhr.open('post',"/update/"+orderid, true);
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
        window.location = "/adminpage";

      } else {
        console.log("error")
      }
    }
  };
  xhr.send();
}
