import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

export default function SignUp() {
  const [email, setemail] = React.useState("");

  const [password, setpassword] = React.useState("");

  const [repass, setrepass] = React.useState("");

  const [uname, setuname] = React.useState("");

  const emailChange = (e) => {
    setemail(e.target.value);
  };

  const passChange = (e) => {
    setpassword(e.target.value);
  };
  const unameChange = (e) => {
    setuname(e.target.value);
  };

  const repassChange = (e) => {
    setrepass(e.target.value);
  };
  let submitHandler = (e) => {
    e.preventDefault();
    if (
      !(
        email.trim().length > 0 &&
        password.trim().length > 0 &&
        repass.trim().length > 0 &&
        uname.trim().length > 0
      )
    ) {
      alert("Invalid data");
      throw new Error("Invalid data");
    }
    if (password !== repass) {
      alert("pass donot match");
      throw new Error("Password Donot Match");
    }

    let requestBody = JSON.stringify({
      uname: uname,
      email: email,
      password: password,
      repassword: repass,
    });
    alert("OKKK");
    setuname("");
    setemail("");
    setpassword("");
    setrepass("");
    fetch("http://"+ window.location.hostname + ":8080/signup/", {
      method: "POST",
      credentials: 'include',
      body: requestBody,
    }).then((res) => {
      res.json().then(function (result) {
        console.log(result);
        if (result.status !== 200) {
          alert("Auth Failed!!");
          throw new Error("Authentication Failed !!");
        }
      });

      alert("Successful!!");
    });
  };
  return (
    <>
      <div>
        <Form className="form-control">
          <Form.Group className="mb-3 " controlId="formBasicEmail">
            <Form.Control
              style={{ "margin-bottom": "5px" }}
              type="uname"
              placeholder="Username"
              onChange={unameChange}
              value={uname}
            />
            <Form.Control
              style={{ "margin-bottom": "5px" }}
              type="email"
              placeholder="Email"
              onChange={emailChange}
              value={email}
            />
            <Form.Control
              style={{ "margin-bottom": "5px" }}
              type="password"
              placeholder="Password"
              onChange={passChange}
              value={password}
            />
            <Form.Control
              style={{ "margin-bottom": "5px" }}
              type="password"
              placeholder="Confirm Password"
              onChange={repassChange}
              value={repass}
            />
          </Form.Group>
          <Button
            style={{ width: "100%" }}
            variant="primary"
            size="wide"
            type="submit"
            onClick={submitHandler}
          >
            Sign up
          </Button>
        </Form>
      </div>
    </>
  );
}
