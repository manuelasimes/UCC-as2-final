// login_cliente.js
import React, { useContext, useState } from 'react';
import { AuthContext } from './auth';
import { Link, useNavigate } from 'react-router-dom';
import Cookies from 'universal-cookie';
import { ToastContainer, toast } from 'react-toastify';
import '../estilo/login_cliente.css';

const ClienteLogin = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const { login } = useContext(AuthContext);
    const navigate = useNavigate();

    const handleLoginCliente = () => {
        fetch('http://localhost/user-res-api/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password }),
        })
        .then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                throw new Error('Invalid credentials');
            }
            return response.json();
        })
        .then(data => {
            if (data.type === false) {
                login(data.accessToken, data.refreshToken, data.type);
                navigate('/');
            } else {
                alert("Usted es un administrador. Para iniciar sesión como administrador, diríjase al área de admin.");
            }
        })
        .catch(error => {
            toast.error('Error al iniciar sesión: ' + error.message);
        });
    };

    return (
        <>
            <div className="bodylogclient">
                <div className="contLogClie1">
                    <div className="contLogClien2">
                        <h1 className="title">Bienvenido Cliente</h1>
                        <div className="form-container">
                            <input
                                type="text"
                                placeholder="Nombre de usuario"
                                id="username"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                            />
                            <input
                                type="password"
                                placeholder="Contraseña"
                                id="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                            <div className="button-container">
                                <button className="buttonClient" onClick={handleLoginCliente}>
                                    Iniciar Sesión
                                </button>
                                <Link to="/register" className="buttonClient">
                                    Registrarse
                                </Link>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <ToastContainer />
        </>
    );
};

export default ClienteLogin;
