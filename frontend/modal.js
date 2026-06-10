/*show modal*/
const deleteBtn = document.querySelector(".delete-btn");
const modal = document.querySelector(".modal-overlay");

deleteBtn.addEventListener("click", () => {
    modal.style.display = "flex";
});

/*hide  modal*/ 
const cancelBtn = document.querySelector(".cancel-modal");

cancelBtn.addEventListener("click", () => {
    modal.style.display = "none";
});