import React from 'react';
import {LoginInput} from "./Login-Register/LoginStyle";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "./Login-Register/LoginComponent";
import Register from "./Login-Register/RegisterComponent";
import Home from "./Home/HomeComponent";

function App() {
  return (
    <div className="App">
        <BrowserRouter>
            <Routes>
                <Route element={<Home/>} path={"/"}/>
                <Route element={<Login/>} path={"/login"}/>
                <Route element={<Register/>} path={"/register"}/>
            </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;
