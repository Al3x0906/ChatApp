import React from "react";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";

export default function SignUp() {
  const [email, setemail] = React.useState("");

  const [password, setpassword] = React.useState("");

  const [repass, setrepass] = React.useState();

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
    if(email.trim().length > 0 && password.trim().length > 0 && repass.trim().length> 0 && uname.trim().length > 0){
      
    }


  };
  return (
    <>
      <div>
        <Form className="form-control">
          <Form.Group className="mb-3" controlId="formBasicEmail">
            <Form.Control
              type="uname"
              placeholder="mark"
              onChange={unameChange}
              value={uname}
            />
            <Form.Control
              type="email"
              placeholder="mark@yahoo.com "
              onChange={emailChange}
              value={email}
            />
          </Form.Group>

          <Form.Group className="mb-3-dark" controlId="formBasicPassword">
            <Form.Control
              type="password"
              placeholder="********"
              onChange={passChange}
              value={password}
            />
            <Form.Control
              type="password"
              placeholder="********"
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
