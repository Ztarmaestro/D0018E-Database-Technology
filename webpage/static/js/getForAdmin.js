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
          console.log(xhr.response + " before parse");
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
console.log("Time to display data");
  for( var i=0, l=data.length; i<l; i++ ) {

    var orderlist = document.createElement("b");
    orderlist.id = 'o'+i;
    var btn = document.createElement("BUTTON");

    IdOrders = document.getElementById('Orderlist').appendChild(orderlist);

    if (data[i].Sent == '1') {
      if (IdOrders != null) {
        if (data[i].Paid == 1){
          var isPaid = "Paid";
        } else {
          var isPaid = "Not Paid";
        }
        IdOrders.innerHTML = "OrderID: " + data[i].IdOrders + " | " + isPaid +  " | " + "PaymentType: " + data[i].PaymentType + " | " + "Sent" + " ";

      }
    } else {
      if (data[i].Paid == 1){
        var isPaid = "Paid";
      } else {
        var isPaid = "Not Paid";
      }
      if (IdOrders != null) {
        var t = document.createTextNode("Click to update order to sent");
        btn.appendChild(t);
        btn.id = "send/"+data[i].IdOrders;
        IdOrders.innerHTML = "OrderId: " + data[i].IdOrders + " | " + isPaid + " | PaymentType: " + data[i].PaymentType + " | "+ " Order not sent " + " ";
        var addbutton = document.getElementById('Orderlist');
        if (addbutton != null) {
          document.getElementById('o'+i).appendChild(btn);
        }
        document.getElementById("send/"+data[i].IdOrders).addEventListener("click", function(){

          var str = this.id;
          var res = str.split("/");
          var orderid = res[1];

          updateOrder(orderid); });
      }
    }
    var mybr = document.createElement('br');
    var addmybr = document.getElementById('o'+i);
      if (addmybr != null) {
        addmybr.appendChild(mybr);
      }
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

window.onload = function() {

  var customerId = callCookie("idcustomer")
  console.log(customerId);

  if (customerId == 1){
    getAllData()
  } else {
    alert("You are not an Admin and should not be here! Bye, bye!!!");
    deleteCookie()
  }

}
