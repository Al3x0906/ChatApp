import Login from "./components/Login/login.js";
import ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import React from "react";
import "./app.css";
import NavBar from "./components/Navbar/navbar";
import SignUp from "./components/Signup/signup.js";
function App() {
  return (
    <>
      <NavBar />
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
