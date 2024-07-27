import { getCookie } from "./cookieManager.js";

document.addEventListener("DOMContentLoaded", function () {
  const profileButton = document.getElementById("profile-button");
  if (profileButton != null) {
    profileButton.addEventListener("click", function (e) {
      e.preventDefault();
      profile(getCookie("logged_user_id"));
    });
  }

  const exitButton = document.getElementById("exit-button");
  if (exitButton != null) {
    exitButton.addEventListener("click", function (e) {
      e.preventDefault();
      exit();
    });
  }

  const booksButton = document.getElementById("books-button");
  if (booksButton != null) {
    booksButton.addEventListener("click", function (e) {
      e.preventDefault();
      books();
    });
  }
});

const profile = (id) => {
  if (id == undefined || id == null) {
    window.location.href = `/auth`;
    return;
  }

  window.location.href = `/users/${id}/profile`;
};

const exit = () => {
  document.cookie.split(";").forEach((cookie) => {
    const eqPos = cookie.indexOf("=");
    const name = eqPos > -1 ? cookie.substring(0, eqPos) : cookie;
    document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
  });

  window.location.href = `/auth`;
};

const books = () => {
  window.location.href = `/`;
};

export { profile };
