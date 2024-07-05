import { getCookie } from "./cookieManager.js";
import { profile } from "./index.js";

document.addEventListener("DOMContentLoaded", function () {
  const buttons = document.querySelectorAll(".list-group-item button");
  buttons.forEach(function (button) {
    button.addEventListener("click", function () {
      donate(this.dataset.bookId);
    });
  });
});

const donate = (id) => {
  fetch(`/books/${id}/donate`, {
    method: "PUT",
    headers: {
      Authorization: "Bearer " + getCookie("jwt"),
    },
  })
    .then((response) => {
      profile(getCookie("logged_user_id"));
    })
    .catch((error) => {
      console.error("Fetch error:", error);
    });
};
