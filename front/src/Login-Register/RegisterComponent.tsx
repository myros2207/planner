import React, {useState} from 'react';
import {LoginInput} from "./LoginStyle";

const Register = () => {
    const [typePassword, setTypePassword] = useState("password")

    const ShowPassword = () => {
        if (typePassword == "password"){
            setTypePassword("text")
        }
        else {
            setTypePassword("password")
        }
    }
    return (
        <div>
            <LoginInput placeholder={"Name"}/>
            <LoginInput placeholder={"Surname"}/>
            <LoginInput placeholder={"Password"} type={typePassword}/>
            <button onClick={ShowPassword}>show </button>
            <LoginInput placeholder={"repeat password"} type={typePassword}/>
            <LoginInput placeholder={"e-mail"}/>

        </div>
    );
};

export default Register;
