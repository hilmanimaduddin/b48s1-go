console.log("coba bisa gak?");

function submitData(event) {
  event.preventDefault();

  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let number = document.getElementById("input-number").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name === "") {
    return alert("Nama diisi dulu ya!");
  } else if (email === "") {
    return alert("Email diisi dulu ya!");
  } else if (number === "") {
    return alert("Nomer diisi dulu ya!");
  } else if (subject === "") {
    return alert("Pilih subject salah satu!");
  } else if (message === "") {
    return alert("Pesan harap diisi!");
  }

  let dataAll = {
    name,
    email,
    number,
    subject,
    message,
  };

  console.log(dataAll);

  let emailReceiver = "phaethoele@gmail.com";

  let a = document.createElement("a");
  a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo nama saya ${name}, \n${message}, silahkan kontak saya di nomor berikut : ${number}, terima kasih.`;
  a.click();
}
