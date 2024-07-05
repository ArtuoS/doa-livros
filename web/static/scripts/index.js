import { getCookie } from "./cookieManager.js";

document.addEventListener("DOMContentLoaded", function () {
  const button = document.getElementById("profile-button");
  if (button == null) {
    return;
  }
  
  button.addEventListener("click", function (e) {
    e.preventDefault();
    profile(getCookie("logged_user_id"));
  });
});

const profile = (id) => {
    window.location.href = `/users/${id}/profile`;
};

export { profile }