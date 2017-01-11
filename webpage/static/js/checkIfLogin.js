function checkUserLogin() {
  var customerId = callCookie("idcustomer");
  console.log(customerId);

  if (customerId != "") {
    console.log("User logged in");
  } else {
    alert("You are not logged in and and should not be here. Please register or login to an account");
    window.location = "/";
  }

}
