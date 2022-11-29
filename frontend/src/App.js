import React, { useEffect } from "react";
import { Route, Routes, Navigate } from "react-router-dom";
import { useDispatch, useSelector } from "react-redux";
import { login, selectUser } from "./features/userSlice";

import "./app.css";
import "./Login-Form-Dark.css";

import SignUp from "./components/Signup/signup.js";
import Login from "./components/Login/login.js";
import NavBar from "./components/Navbar/navbar";
import Chat from "./components/chat/chat.js";

function App() {
  const dispatch = useDispatch();
  const user = useSelector(selectUser);
  useEffect(() => {
    if (!user) {
      fetch("http://" + window.location.hostname + ":8080/login/", {
        method: "GET",
        credentials: "include",
      }).then((res) => {
        res.json().then(function (result) {
          if (result.status == 200) {
            const userinfo = result.userinfo;
            dispatch(
              login({
                username: userinfo.username,
                userid: userinfo.id,
                email: userinfo.email,
              })
            );
          }
        });
      });
    }
  }, []);

  return (
    <>
      <div className="login-dark">
        <NavBar />
        <Routes>
          {console.log("user :: " + user)}
          {user && <Route path="/" element={<Chat />} />}

          {!user && (
            <Route path="/" element={<Navigate to="/login" replace />} />
          )}

          {user && (
            <Route path="/login" element={<Navigate to="/" replace />} />
          )}

          {!user && <Route path="/login" element={<Login />} />}

          {!user && <Route path="/signup" element={<SignUp />} />}

          {user && (
            <Route path="/signup" element={<Navigate to="/" replace />} />
          )}
        </Routes>
      </div>
    </>
  );
}

export default App;
