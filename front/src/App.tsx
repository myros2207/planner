import React from 'react';
import {LoginInput} from "./Login-Register/LoginStyle";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Login from "./Login-Register/LoginComponent";
import Register from "./Login-Register/RegisterComponent";
import Home from "./Home/HomeComponent";
import {GlobalStyle} from "./globalStyle";
import StartPage from "./StartPage/StartPageComponent";

function App() {
  return (
    <div className="App">
        <BrowserRouter>
            <GlobalStyle/>
            <Routes>
                <Route element={<StartPage/>} path={"/"}/>
                <Route element={<Login/>} path={"/login"}/>
                <Route element={<Register/>} path={"/register"}/>
                <Route element={<Home/>} path={"/home"}/>
            </Routes>
        </BrowserRouter>
    </div>
  );
}

export default App;
