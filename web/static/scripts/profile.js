import { getCookie } from "./cookieManager.js";
import { profile } from "./index.js";

document.addEventListener("DOMContentLoaded", function () {
  if (getCookie("logged_user_id") == null) {
    window.location.href = "/auth";
    return;
  }

  const buttons = document.querySelectorAll(".list-group-item .donation");
  buttons.forEach(function (button) {
    button.addEventListener("click", function () {
      donate(this.dataset.bookId);
    });
  });

  const deleteButtons = document.querySelectorAll(".list-group-item .delete-book-button");
  deleteButtons.forEach(function (button) {
    button.addEventListener("click", function () {
      console.log("Delete button: ", button.dataset.bookId);
      deleteBook(this.dataset.bookId);
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

const deleteBook = (id) => {
  console.log(id);
  fetch(`/books/${id}`, {
    method: "DELETE",
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
