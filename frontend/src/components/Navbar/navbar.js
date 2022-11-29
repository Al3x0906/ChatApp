import Container from "react-bootstrap/Container";
import React from "react";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import { Link } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import { selectUser } from "../../features/userSlice";
import { logout } from "../../features/userSlice";

function NavBar() {
  const user = useSelector(selectUser);
  const dispatch = useDispatch();
  const logoutHandler = () => {
    dispatch(logout());
    fetch("http://" + window.location.hostname + ":8080/logout/", {
      method: "POST",
      credentials: "include",
    }).then((res) => {
      console.log(res);
    });
  };
  return (
    <Navbar bg="dark" variant="dark">
      <Container>
        <Navbar.Brand as={Link} to="/">
          ChatApp
        </Navbar.Brand>
        <Nav className="me-auto">
          {!user && (
            <Nav.Link as={Link} to="/login">
              LogIn
            </Nav.Link>
          )}

          {!user && (
            <Nav.Link as={Link} to="/signup">
              SignUp
            </Nav.Link>
          )}
          {user && <Nav.Link onClick={logoutHandler}>logout</Nav.Link>}
          <Nav.Link as={Link} to="/about">
            About
          </Nav.Link>
        </Nav>
      </Container>
    </Navbar>
  );
}

export default NavBar;
