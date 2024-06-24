document.addEventListener("DOMContentLoaded", function () {
  const buttons = document.querySelectorAll(".book button");
  buttons.forEach(function (button) {
    button.addEventListener("click", function () {
      if (!this.disabled) {
        redeem({
          ToUserId: parseInt(localStorage.getItem("logged_user")),
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
    },
    body: JSON.stringify(data),
  })
    .then((data) => {
      console.log("Books data:", data);
    })
    .catch((error) => {
      console.error("Fetch error:", error);
    });
};
