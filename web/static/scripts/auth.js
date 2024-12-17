import { setCookie } from "./cookieManager.js";
import { showAlert } from "./utils.js";

document.addEventListener("DOMContentLoaded", function () {
  const button = document.getElementById("auth");
  button.addEventListener("click", function () {
    auth({
      Email: document.getElementById("email").value,
      Password: document.getElementById("password").value,
    });
  });

  const saveUserButton = document.getElementById("save-user-button");
  saveUserButton.addEventListener("click", function () {
    const form = document.getElementById("userForm");
    const userData = {
      firstName: form.querySelector("#firstName").value,
      lastName: form.querySelector("#lastName").value,
      email: form.querySelector("#email").value,
      password: form.querySelector("#password").value,
    };
    saveUser(userData);
  });
});



const auth = (data) => {
  fetch("/auth", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (!response.ok) {
        const errorData = response.json();
        throw new Error(errorData.message || "Usuário inválido.");
      }
      return response.json();
    })
    .then((data) => {
      setCookie("jwt", data.token, 1);
      setCookie("logged_user_id", data.loggedUserId, 1);
      window.location.href = "/";
    })
    .catch((error) => {
      showAlert(error.message, "danger");
    });
};



const saveUser = (data) => {
  fetch("/users", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  })
    .then((response) => {
      console.log(response);
      if (!response.ok) {
        console.log("response", response);
        const errorData = response.json();
        throw new Error(errorData.message || "Erro ao salvar usuário.");
      }
      bootstrap.Modal.getInstance(document.getElementById("userModal")).hide();
    })
    .then((data) => {
      window.location.href = "/auth";
    })
    .catch((error) => {
      console.log(error);
      showAlert(error.message, "danger");
    });
};
