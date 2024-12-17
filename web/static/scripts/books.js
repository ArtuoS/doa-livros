import { getCookie } from "./cookieManager.js";
import { showModal } from "./utils.js";

document.addEventListener("DOMContentLoaded", function () {
  if (getCookie("logged_user_id") == null) {
    window.location.href = "/auth";
    return;
  }

  const buttons = document.querySelectorAll(".book button");
  buttons.forEach(function (button) {
    if (parseInt(button.closest(".book").dataset.owner) == parseInt(getCookie("logged_user_id"))) {
      button.disabled = true;
      button.className = "btn btn-secondary";
    } 

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

  const createBookModal = document.getElementById("create-book-modal");
  const createBookButton = createBookModal.querySelector("#create-book-button");
  createBookButton.addEventListener("click", function () {
    create({
      Title: createBookModal.querySelector("#title").value,
      Author: createBookModal.querySelector("#author").value,
      UserId: parseInt(getCookie("logged_user_id"))
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

const create = (data) => {
  fetch("/books", {
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
