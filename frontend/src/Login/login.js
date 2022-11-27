import React from "react";
import ReactDOM from "react-dom";
import "./login.css";
export default function Login() {
  const [isLogin, setisLogin] = React.useState(false);

  const [user, setuser] = React.useState(true);

  let switchModeHandler = (e) => {
    setisLogin(!isLogin);
  };

  let submitHandler = (e) => {
    e.preventDefault();
    const email = e.target.email.value;
    const password = e.target.password.value;
    if (email.trim().length > 0 && password.trim().length > 0) {
    } else {
      console.alert("Password Or email too small");
    }

    let requestBody = JSON.stringify({
      email: email,
      password: password,
    });

    if (user) {
      fetch("http://localhost:8080/login/", {
        method: "POST",
        body: requestBody,
      }).then((res) => {
        if (res.status !== 200) {
          alert("Auth Failed!!");
          throw new Error("Authentication Failed !!");
        }
        setuser(true);
        alert("Successful!!");
      });
    } else {
      fetch("http://localhost:8000/signup/", {
        method: "POST",
        body: requestBody,
      }).then((res) => {
        if (res.status !== 200) {
          alert("Auth Failed!!");
          throw new Error("Authentication Failed !!");
        }
        setuser(true);
        alert("Successful!!");
      });
    }
  };
  return (
    <form className="auth-form" onSubmit={submitHandler}>
      <div className="form-control">
        <label htmlFor="email">Email</label>
        <input type="email" id="email" placeholder="Email"></input>
      </div>

      <div className="form-control">
        <label htmlFor="password">Password</label>
        <input type="password" id="password" placeholder="Password"></input>
      </div>

      <div className="form-actions">
        <button type="submit">Submit</button>
        <button type="button" onClick={switchModeHandler}>
          Switch To {isLogin ? "Sign Up " : "Login"}
        </button>
      </div>
    </form>
  );
}
