import React, { useContext, useEffect, useCallback, useState } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from './pages/login/auth';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';

const ProtectedRoute = ({ children, adminOnly = false }) => {
    const { auth, login, logout } = useContext(AuthContext);
    const [loading, setLoading] = useState(true);

    const isTokenExpired = (token) => {
        const { exp } = jwtDecode(token);
        return Date.now() >= exp * 1000;
    };

    const refreshToken = useCallback(async () => {
        try {
            const response = await axios.post('http://localhost/user-res-api/refresh', {
                refreshToken: auth.refreshToken
            });
            const { accessToken, refreshToken, type } = response.data;
            login(accessToken, refreshToken, type);
        } catch (error) {
            console.error('Error al refrescar el token:', error);
            logout();
        } finally {
            setLoading(false);
        }
    }, [auth.refreshToken, login, logout]);

    useEffect(() => {
        if (auth.accessToken) {
            if (isTokenExpired(auth.accessToken)) {
                refreshToken();
            } else {
                setLoading(false);
            }
        } else {
            setLoading(false);
        }
    }, [auth.accessToken, refreshToken]);

    if (loading) {
        return <div>Cargando...</div>;
    }

    if (!auth.accessToken) {
        return <Navigate to={adminOnly ? '/login-admin' : '/login-cliente'} />;
    }

    if (adminOnly && auth.userType !== true) {
        return <Navigate to="/" />;
    }

    return children;
};

export default ProtectedRoute;
