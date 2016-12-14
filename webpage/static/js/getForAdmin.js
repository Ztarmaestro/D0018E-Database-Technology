
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

  Product_name = document.getElementById('Name');
  Price = document.getElementById('Price');
  Description = document.getElementById('Description');
  UnitsInStock = document.getElementById('UnitsInStock');

  if(data.ProductAvailable == 1){

    if(document.getElementById("Name") != null){
        Product_name.innerHTML = data.ProductName;
    }
    if(document.getElementById("Price") != null){
        Price.innerHTML = "$"+data.Price;
    }
    if(document.getElementById("Description") != null){
        Description.innerHTML = data.ProductDescription;
    }
    if(document.getElementById("UnitsInStock") != null){
        UnitsInStock.innerHTML = data.UnitsInStock;
    }
  }
}

function updateData(type) {
    //Type is the users id that is saved in the session. carmodel is the car that is added to the cart
    var xhr = typeof XMLHttpRequest != 'undefined'
      ? new XMLHttpRequest()
      : new ActiveXObject('Microsoft.XMLHTTP');
    xhr.open('get',"/everything/"+type, true);
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
          window.location = "/checkout";

        } else {
          console.log("error")
        }
      }
    };
    xhr.send();
}

//Filter the admin update (delete, add, update)
function whatUpdate(type) {

  if (delete){
    updateData('delete');
  }
  if (add){
    updateData('add');
  }
  if (update){
    updateData('update');
  }

}
