import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/cuenta.css';

function AccountDetails() {
  const [accountDetails, setAccountDetails] = useState({
    nombre: '',
    apellido: '',
    username: '',
    password: '*****',
    email: ''
  });

  const { isLoggedAdmin } = useContext(AuthContext);
  const { isLoggedCliente } = useContext(AuthContext);
  const { logout } = useContext(AuthContext);

  useEffect(() => {
    const getUser = () => {
      if (isLoggedCliente) {
        const accountId = localStorage.getItem("id_cliente");
        fetch(`http://localhost:8090/cliente/${accountId}`)
          .then(response => response.json())
          .then(data => {
            setAccountDetails({
              nombre: data.name,
              apellido: data.last_name,
              username: data.username,
              password: '*****',
              email: data.email
            });
          })
          .catch(error => {
            console.error('Error al obtener los datos del cliente:', error);
          });
      }
      else if (isLoggedAdmin) {
        const accountId = localStorage.getItem("id_admin");
        fetch(`http://localhost:8090/admin/${accountId}`)
          .then(response => response.json())
          .then(data => {
            setAccountDetails({
              nombre: data.name,
              apellido: data.last_name,
              username: data.username,
              password: '*****',
              email: data.email
            });
          })
          .catch(error => {
            console.error('Error al obtener los datos del administrador:', error);
          });
      }
      else {
        window.location.href = '/login-cliente';
      }
    };

    getUser();
  }, [isLoggedAdmin, isLoggedCliente]);

  const cerrarSesion = () => {
    logout();
  };

  const reservas = () => {
    window.location.href = '/reservas-cliente';
  };

  return (
    <body className="bodyCuenta">
    <div className="containerCU">
      <div className="account-form">
        <h2 className="tituloDC">Detalles de la cuenta</h2>
        <div className="account-field">
          <p>Nombre: {accountDetails.nombre}</p>
        </div>
        <div className="account-field">
          <p>Apellido: {accountDetails.apellido}</p>
        </div>
        <div className="account-field">
          <p>UserName: {accountDetails.username}</p>
        </div>
        <div className="account-field">
          <p>Password: {accountDetails.password}</p>
        </div>
        <div className="account-field">
          <p>Email: {accountDetails.email}</p>
        </div>
        <div className="account-buttons">
          <button className="logout-button" onClick={cerrarSesion}>Cerrar sesión</button>
          <button className="reservations-button" onClick={reservas}>Reservas</button>
        </div>
      </div>
    </div>
    </body>
  );
}

export default AccountDetails;