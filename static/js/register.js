//const personCheckbox = document.getElementById("person");
const companyCheckbox = document.getElementById("company");
const companyDetails = document.getElementById("companyDetails");

// personCheckbox.addEventListener("change", function() {
//   if (personCheckbox.checked) {
//     companyCheckbox.checked = false;
//     companyDetails.style.display = "none";
//   }
// });

companyCheckbox.addEventListener("change", function() {
  if (companyCheckbox.checked) {
    //personCheckbox.checked = false;
    companyDetails.style.display = "block";
  } else {
    companyDetails.style.display = "none";
  }
});

const passwordInput = document.getElementById("password");
const repeatPasswordInput = document.getElementById("repeatPassword");

function validatePassword() {
  if (passwordInput.value !== repeatPasswordInput.value) {
    repeatPasswordInput.setCustomValidity("Passwords must match");
  } else {
    repeatPasswordInput.setCustomValidity("");
  }
}

passwordInput.addEventListener("change", validatePassword);
repeatPasswordInput.addEventListener("change", validatePassword);