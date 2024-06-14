import React, { useEffect, useState } from "react";
import axios from "axios";
import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";
import './Login.css';

const Cookie = new Cookies();
const Login = ({ token, setToken }) => {

    const navigate = useNavigate();
    const [register, setRegister] = useState(false);
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    useEffect(
        () => {
            if (token) {
                navigate("/");
            }
        }, [token]
    )


    const handleSubmitRegister = async (event) => {
        event.preventDefault();
        const targetUrl = 'http://localhost:8080/auth/register';
        axios.post(targetUrl, {
            name,
            email,
            password,
        }, {
            headers: {
                'Content-Type': 'application/json',
            },
        }).then(response => {

            if (response.data.data.token) {
                setToken(response.data.data.token);
            }

        }).catch(function (error) {
            if (error.response) {
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
            } else if (error.request) {
                console.log(error.request);
            } else {
               
                console.log('Error', error.message);
            }
            console.log(error.config);
        });

    };

    const handleSubmitLogin = async (event) => {
        event.preventDefault();
        const targetUrl = 'http://localhost:8080/auth/login';
        axios.post(targetUrl, {
            email,
            password,
        }, {
            headers: {
                'Content-Type': 'application/json',
            },
        }).then(response => {

            if (response.data.data.token) {
                setToken(response.data.data.token);
                Cookie.set("rol", response.data.data.data.role, { path: "/" });
            }
            console.log(response)

        }).catch(function (error) {
            if (error.response) {
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
            } else if (error.request) {
               
                console.log(error.request);
            } else {
                
                console.log('Error', error.message);
            }
            console.log(error.config);
        });

    };


    const LoginPage = () => {
        return (
            <div>
                <h1>Inicio de sesion</h1>
                <form onSubmit={handleSubmitLogin}>
                    <label>
                        Email:
                        <input type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
                    </label>
                    <label>
                        Contraseña:
                        <input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                    </label>
                    <button type="submit">Iniciar sesion</button>
                </form>
            </div>
        );

    }


    const CreateUserPage = () => {

        return (
            <div>
                <h1>Crear usuario</h1>
                <form onSubmit={handleSubmitRegister}>
                    <label>
                        Nombre:
                        <input type="text" name="name" value={name} onChange={(e) => setName(e.target.value)} />
                    </label>
                    <label>
                        Email:
                        <input type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
                    </label>
                    <label>
                        Contraseña:
                        <input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                    </label>
                    <button type="submit">Crear usuario</button>
                </form>
            </div>
        );
    }

    return (
        <div>
            <button onClick={() => {
                setRegister(!register);
                console.log(Cookie.get("rol"));
            }}> CREAR USUARIO / REGISTRATE</button>
            {register ? CreateUserPage() : LoginPage()}
        </div>
    )

};

export default Login;