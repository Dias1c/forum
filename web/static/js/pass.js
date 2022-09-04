function confirmPasswordOnChange() {
  let password = document.querySelector('#password-original');
  let confirm = document.querySelector('#password-confirm');
  if (confirm.value !== password.value) {
    alert("Passwords not equal")
  }
}