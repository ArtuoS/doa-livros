import { getCookie } from "./cookieManager.js";

document.addEventListener("DOMContentLoaded", function () {
  const buttons = document.querySelectorAll(".book button");
  buttons.forEach(function (button) {
    button.addEventListener("click", function () {
      if (!this.disabled) {
        redeem({
          ToUserId: parseInt(getCookie("logged_user_id")),
          FromUserId: parseInt(this.closest(".book").dataset.owner),
          BookId: parseInt(this.closest(".book").id),
        });
      }
    });
  });
});

const redeem = (data) => {
  fetch("/books/redeem", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + getCookie("jwt"),
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.redirected) {
        window.location.href = response.url;
      } else {
        return response.json();
      }
    })
    .catch((error) => {
      console.error("Fetch error:", error);
    });
};
