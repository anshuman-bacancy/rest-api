
console.log("HI")

document.getElementById("postButton").addEventListener("click", function() {
  var empname = document.getElementById("name")
  var email = document.getElementById("email")
  var position = document.getElementById("position")

  console.log(empname, email)
});