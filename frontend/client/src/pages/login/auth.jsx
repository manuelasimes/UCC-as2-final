// auth.js
import React, { createContext, useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { jwtDecode } from 'jwt-decode';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [auth, setAuth] = useState({ accessToken: null, refreshToken: null, userType: null });

    useEffect(() => {
        const accessToken = new Cookies().get('accessToken');
        const refreshToken = new Cookies().get('refreshToken');
        if (accessToken && refreshToken) {
            setAuth({ accessToken, refreshToken, userType: jwtDecode(accessToken).type });
        }
    }, []);

    const login = (accessToken, refreshToken, userType) => {
        const cookies = new Cookies();
        cookies.set('accessToken', accessToken, { path: '/' });
        cookies.set('refreshToken', refreshToken, { path: '/' });
        setAuth({ accessToken, refreshToken, userType });
        
        // Almacenar el ID del cliente o administrador en localStorage
        localStorage.setItem('user_id', jwtDecode(accessToken).user_id);
    };

    const logout = () => {
        const cookies = new Cookies();
        cookies.remove('accessToken', { path: '/' });
        cookies.remove('refreshToken', { path: '/' });
        setAuth({ accessToken: null, refreshToken: null, userType: null });

        // Limpiar el localStorage al cerrar sesión
        localStorage.removeItem('user_id');
    };

    return (
        <AuthContext.Provider value={{ auth, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export { AuthContext };
