
//function getAllData() {
window.onload = function() {
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
    console.log("Row "+i);

    var orderlist = document.createElement("b");
    orderlist.id = 'o'+i;
    var btn = document.createElement("BUTTON");

    var str = data[i].PaymentType;
    var res = str.split("/");
    var orderid = res[1];
    console.log(res[1]);

    IdOrders = document.getElementById('Orderlist').appendChild(orderlist);

    if (data[i].Sent == '1') {
      console.log("Order "+ i +" sent");
      if (IdOrders != null) {
        if (data[i].Paid == 1){
          var isPaid = "Paid";
        } else {
          var isPaid = "Not Paid";
        }
        console.log("Orderlist Not null");
        IdOrders.innerHTML = "OrderID: " + res[1]+ " | " + isPaid +  " | " + "PaymentType: " + res[0] + " | " + "Sent" + " ";

      }
    } else {
      console.log("Order "+ i +" not sent");
      if (data[i].Paid == 1){
        var isPaid = "Paid";
      } else {
        var isPaid = "Not Paid";
      }
      if (IdOrders != null) {
        console.log("Orderlist Not null");
        var t = document.createTextNode("Click to update order to sent");
        btn.appendChild(t);
        btn.id = "send"+i;
        IdOrders.innerHTML = "OrderId: " + res[1] + " | " + isPaid + " | PaymentType: " + res[0] + " | "+ " Order not sent " + " ";
        var addbutton = document.getElementById('Orderlist');
        if (addbutton != null) {
          console.log("Orderlist Not null");
          document.getElementById('o'+i).appendChild(btn);
        }
        console.log(orderid);
        document.getElementById("send"+i).onclick = function() { updateOrder(orderid); };
      }
    }
    var mybr = document.createElement('br');
    var addmybr = document.getElementById('o'+i);
      if (addmybr != null) {
        console.log("addmybr not null");
        addmybr.appendChild(mybr);
      }
    console.log("Done");
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
