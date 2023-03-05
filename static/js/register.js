const companyCheckbox = document.getElementById("company");
const companyDetails = document.getElementById("companyDetails");


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

const emailInput = document.getElementById("email");
const repeatEmailInput = document.getElementById("repeatEmail");

function validateEmail() {
  if (emailInput.value !== repeatEmailInput.value) {
    repeatEmailInput.setCustomValidity("Emails must match");
  } else {
    repeatEmailInput.setCustomValidity("");
  }
}

emailInput.addEventListener("change", validateEmail);
repeatEmailInput.addEventListener("change", validateEmail);