import React, { createContext, useState } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [isLoggedCliente, setIsLoggedCliente] = useState(localStorage.getItem('authC') === 'true');
  const [isLoggedAdmin, setIsLoggedAdmin] = useState(localStorage.getItem('authA') === 'true');

  const loginAdmin = (newToken, id) => {
    setIsLoggedAdmin(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_admin', id);
    localStorage.setItem('authA', true);
  };

  const loginCliente = (newToken, id) => {
    setIsLoggedCliente(true);
    localStorage.setItem('token', newToken);
    localStorage.setItem('id_cliente', id);
    localStorage.setItem('authC', true);
  };

  const logout = () => {
    setIsLoggedCliente(false);
    setIsLoggedAdmin(false);
    localStorage.removeItem('token');
    localStorage.removeItem('id_admin');
    localStorage.removeItem('id_cliente');
    localStorage.setItem('authA', false);
    localStorage.setItem('authC', false);
  };

  const propiedades = {
    isLoggedCliente,
    isLoggedAdmin,
    loginCliente,
    loginAdmin,
    logout,
  };

  return (
    <AuthContext.Provider value={propiedades}>
      {children}
    </AuthContext.Provider>
  );
};
