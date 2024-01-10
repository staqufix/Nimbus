const username = document.getElementById("username");
const password = document.getElementById("password");
const login = document.getElementById("login");

login.onclick = () => {
  console.log("clicked");

  let data = new FormData();
  data.append(
    "json",
    JSON.stringify({
      username: username.textContent,
      password: password.textContent,
    })
  );

  fetch("/auth", {
    method: "POST",
    headers: {
      Accept: "application/json, text/plain, */*",
      "Content-Type": "application/json",
    },
    body: data,
  });

  console.log("send fetch");
};
