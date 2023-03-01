const divPostive = document.getElementById("positiveAlert");
const divNegative = document.getElementById("negativeAlert");

if (divNegative.textContent.trim().length > 0) {
  divNegative.classList.add("alert");
  divNegative.classList.add("alert-danger");
}
if (divPostive.textContent.trim().length > 0) {
    divPostive.classList.add("alert");
    divPostive.classList.add("alert-success");
}