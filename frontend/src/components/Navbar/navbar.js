import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";

function NavBar() {
  return (
    <Navbar bg="dark" variant="dark">
      <Container>
        <Navbar.Brand href="/">ChatApp</Navbar.Brand>
        <Nav className="me-auto">
          <Nav.Link href="/login">LogIn</Nav.Link>
          <Nav.Link href="/signup">SignUp</Nav.Link>
          <Nav.Link href="/about">About</Nav.Link>
        </Nav>
      </Container>
    </Navbar>
  );
}

export default NavBar;
