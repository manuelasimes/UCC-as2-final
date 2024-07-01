import React, { useContext, useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from './pages/login/auth';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

const ProtectedRoute = ({ children, adminOnly = false }) => {
    const { auth, login, logout } = useContext(AuthContext);

    // Función para verificar si el access token está vencido
    const isTokenExpired = (token) => {
        const { exp } = jwtDecode(token);
        return Date.now() >= exp * 1000;
    };

    // Función para renovar el access token utilizando el refresh token
    const refreshToken = async () => {
        try {
            const response = await axios.post('http://localhost/user-res-api/refresh', {
                refreshToken: auth.refreshToken
            });
            const { accessToken, refreshToken, type } = response.data;
            console.log("data refresh token", response.data);
            login(accessToken, refreshToken, type);
        } catch (error) {
            console.error('Error al refrescar el token:', error);
            logout();
        }
    };

    useEffect(() => {
        // Verificar si el access token está vencido y renovarlo si es necesario
        if (auth.accessToken && isTokenExpired(auth.accessToken)) {
            refreshToken();
        }
    }, [auth.accessToken]);

    if (!auth.accessToken) {
        return <Navigate to={adminOnly ? '/login-admin' : '/login-cliente'} />;
    }

    if (adminOnly && auth.userType !== true) {
        return <Navigate to="/" />;
    }

    return children;
};

export default ProtectedRoute;
