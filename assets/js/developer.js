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