import React from 'react';
import {useCookies} from "react-cookie";

const Home = () => {

    const [cookies, setCookies, removeCookies] = useCookies()

    setCookies('test', 'Nmae1',{ path: '/test' })

    const Test = () => {
        removeCookies('test', {path: '/test'})
    }
    return (
        <div>
            <button onClick={Test}>Clear</button>
        </div>
    );
};

export default Home;
