import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

import "./login.css";
export default function Login() {
  const [email, setemail] = React.useState("");

  const [password, setpassword] = React.useState("");

  const emailChange = (e) => {
    setemail(e.target.value);
  };

  const passChange = (e) => {
    setpassword(e.target.value);
  };

  let submitHandler = (e) => {
    e.preventDefault();

    if (!(email.trim().length > 0 && password.trim().length > 0)) {
      alert("Password Or email too small");
      throw new Error("Password or email too small ");
    }

    let requestBody = JSON.stringify({
      email: email,
      password: password,
    });
    setemail("");
    setpassword("");
    fetch("http://localhost:8080/login/", {
      method: "POST",
      body: requestBody,
    }).then((res) => {
      console.log(res);
      if (res.status !== 200) {
        alert("Auth Failed!!");
        throw new Error("Authentication Failed !!");
      }

      alert("Successful!!");
    });
  };

  return (
    <>
      <div>
        <Form className="form-control">
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Label>Email</Form.Label>
            <Form.Control
              type="email"
              placeholder="mark@yahoo.com "
              onChange={emailChange}
              value={email}
            />
          </Form.Group>

          <Form.Group className="mb-3-dark" controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              placeholder="********"
              onChange={passChange}
              value={password}
            />
          </Form.Group>
          <Button
            style={{ width: "100%" }}
            variant="primary"
            size="wide"
            type="submit"
            onClick={submitHandler}
          >
            Log in
          </Button>
        </Form>
      </div>
    </>
  );
}
