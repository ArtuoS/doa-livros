const showAlert = (message, type = "danger") => {
    const alertContainer = document.getElementById("alert-container");
    const alert = document.createElement("div");
    alert.className = `alert alert-${type} alert-dismissible fade show`;
    alert.role = "alert";
    alert.innerHTML = `
      ${message}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    `;
    alertContainer.appendChild(alert);

    setTimeout(() => {
        alert.classList.remove("show");
        alert.addEventListener("transitionend", () => alert.remove());
    }, 5000);
}

function showModal({ title, message, buttonText = "Close", modalType = "info" }) {
    const modalTitle = document.getElementById("genericModalLabel");
    modalTitle.textContent = title;

    const modalMessage = document.getElementById("genericModalMessage");
    modalMessage.innerHTML = message;

    const modalButton = document.getElementById("genericModalButton");
    modalButton.textContent = buttonText;

    const modalHeader = modalTitle.parentElement;
    modalHeader.classList.remove("bg-danger", "bg-success", "bg-warning", "bg-primary");
    switch (modalType) {
        case "error":
            modalHeader.classList.add("bg-danger", "text-white");
            break;
        case "success":
            modalHeader.classList.add("bg-success", "text-white");
            break;
        case "warning":
            modalHeader.classList.add("bg-warning");
            break;
        default:
            modalHeader.classList.add("bg-primary", "text-white");
    }

    new bootstrap.Modal(document.getElementById("genericModal")).show();
}

export { showAlert, showModal };