import React from "react";
import { login } from "utils/api";
import {browserHistory} from "utils/history";
const Login = () => {
    const [username, usernameInput] = Input({type: "text", name: "username"});
    const [password, passwordInput] = Input({type: "password", name: "password"});
    const submit = () =>{
        login(username, password).then(data => {
            const accessToken = data.access_token;
            localStorage.setItem("access_token", accessToken);
            browserHistory.push("/");
        }).catch(err => {
            if (err.response) {
                const data = err.response.data;
                alert(data.message);
            }
        })
    }
    return (
        <div className="content">
            <div className="container">
                <div className="row">
                    <div className="col">
                        {usernameInput} 
                        {passwordInput}
                        <input type="submit" value="Login" onClick={submit}></input>
                    </div>
                </div>
            </div>
        </div>
    );
};

const Input = ({ type, name }) => {
    const [value, setValue] = React.useState("");
    const input = <input type={type} name={name} placeholder={name} onChange={e => setValue(e.target.value)} required />
    return [value, input]
};

export default Login;