const hide = document.querySelector(".editter");
const topuser = document.querySelector(".top-user");
const userinfo = document.querySelector(".user-info");
const userscore = document.querySelector(".user-score");
const form = document.querySelector(".edit-form");


hide.addEventListener("click", () => {
    topuser.classList.toggle("inform");
    userinfo.classList.toggle("inform");
    userscore.classList.toggle("inform");
    form.classList.toggle("inform");
  });

  var password = document.getElementById("password")
  , confirm_password = document.getElementById("confirm_password");

function validatePassword(){
  if(password.value != confirm_password.value) {
    confirm_password.setCustomValidity("Sanırım Yanlış Bir Tuşa Bastın. Birdaha Dene 😉");
  } else {
    confirm_password.setCustomValidity('');
  }
}

password.onchange = validatePassword;
confirm_password.onkeyup = validatePassword;