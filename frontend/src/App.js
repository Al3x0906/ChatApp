import Login from "./components/Login/login.js";
import ReactDOM from "react-dom";
import { BrowserRouter, Route, Routes } from "react-router-dom";

import React from "react";
import "./app.css";
import NavBar from "./components/Navbar/navbar";
import SignUp from "./components/Signup/signup.js";
import "./Login-Form-Dark.css";
function App() {
  return (
    <>
      <div className="login-dark">
        <NavBar />
          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<SignUp />} />
          </Routes>

      </div>
    </>
  );
}

export default App;
