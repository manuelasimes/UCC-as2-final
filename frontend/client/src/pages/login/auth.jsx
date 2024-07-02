import React, { createContext, useState, useEffect } from 'react';
import Cookies from 'universal-cookie';
import { jwtDecode } from 'jwt-decode';

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [auth, setAuth] = useState({ accessToken: null, refreshToken: null, userType: null });

    useEffect(() => {
        const cookies = new Cookies();
        const accessToken = cookies.get('accessToken');
        const refreshToken = cookies.get('refreshToken');
        if (accessToken && refreshToken) {
            try {
                const decoded = jwtDecode(accessToken);
                setAuth({ accessToken, refreshToken, userType: decoded.type });
                localStorage.setItem('user_id', decoded.user_id);
            } catch (error) {
                console.error('Error al decodificar el token:', error);
                cookies.remove('accessToken', { path: '/' });
                cookies.remove('refreshToken', { path: '/' });
                setAuth({ accessToken: null, refreshToken: null, userType: null });
            }
        }
    }, []);

    const login = (accessToken, refreshToken, userType) => {
        const cookies = new Cookies();
        cookies.set('accessToken', accessToken, { path: '/' });
        cookies.set('refreshToken', refreshToken, { path: '/' });
        setAuth({ accessToken, refreshToken, userType });
        
        const decoded = jwtDecode(accessToken);
        localStorage.setItem('user_id', decoded.user_id);
        console.log("token user", decoded);
    };

    const logout = () => {
        const cookies = new Cookies();
        cookies.remove('accessToken', { path: '/' });
        cookies.remove('refreshToken', { path: '/' });
        setAuth({ accessToken: null, refreshToken: null, userType: null });
        localStorage.removeItem('user_id');
    };

    return (
        <AuthContext.Provider value={{ auth, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export { AuthContext };
