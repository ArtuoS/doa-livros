import { setCookie } from "./cookieManager.js";

document.addEventListener("DOMContentLoaded", function () {
  const button = document.getElementById("auth");
  button.addEventListener("click", function () {
    auth({
      Email: document.getElementById("email").value,
      Password: document.getElementById("password").value,
    });
  });
});

const auth = (data) => {
  fetch("/users/auth", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => response.json())
    .then((data) => {
      setCookie("jwt", data.token, 1);
      window.location.href = "/";
    })
    .catch((error) => {
      console.log(error);
    });
};
