import React, { useState } from 'react';
import axios from 'axios';
import Cookies from 'universal-cookie';

const cookies = new Cookies();

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

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
                cookies.set("token", response.data.data.token, { path: "/" });
                cookies.set("rol", response.data.data.role, { path: "/" });
                // Redirigir al usuario a la página de inicio
                window.location.href = "/";
            }
            console.log(response)
        }).catch(error => {
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

    return (
        <div>
            <h1>Inicio de sesión</h1>
            <form onSubmit={handleSubmitLogin}>
                <label>
                    Email:
                    <input type="text" name="email" value={email} onChange={(e) => setEmail(e.target.value)} />
                </label>
                <label>
                    Contraseña:
                    <input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </label>
                <button type="submit">Iniciar sesión</button>
            </form>
        </div>
    );
};

export default Login;
