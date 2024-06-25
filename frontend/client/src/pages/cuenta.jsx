import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/cuenta.css';
import Cookies from "universal-cookie";


function AccountDetails() {
  const [accountDetails, setAccountDetails] = useState({
    Name: '',
    LastName: '',
    username: '',
    password: '*******',
    Email: ''
  });

  const { isLoggedAdmin } = useContext(AuthContext);
  const { isLoggedCliente } = useContext(AuthContext);
  const { logout } = useContext(AuthContext);

  useEffect(() => {
    const getUser = () => {
      if (isLoggedCliente) {
        const accountId = localStorage.getItem("id_cliente");
        fetch(`http://localhost/user-res-api/user/${accountId}`)
          .then(response => response.json())
          .then(data => {
            setAccountDetails({
              Name: data.name,
              LastName: data.last_name,
              username: data.username,
              password: '*******',
              Email: data.email
            });
          })
          .catch(error => {
            console.error('Error al obtener los datos del cliente:', error);
          });
      }
      else if (isLoggedAdmin) {
        const accountId = localStorage.getItem("id_admin");
        fetch(`http://localhost/user-res-api/user/${accountId}`)
          .then(response => response.json())
          .then(data => {
            setAccountDetails({
              Name: data.name,
              LastName: data.last_name,
              username: data.username,
              password: '*******',
              Email: data.email
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
    const Cookie = new Cookies();
    Cookie.set("user_id", -1, {path: "/"})
    Cookie.set("user_type", false, {path:"/"})
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
          <p>Nombre: {accountDetails.Name}</p>
        </div>
        <div className="account-field">
          <p>Apellido: {accountDetails.LastName}</p>
        </div>
        <div className="account-field">
          <p>Nombre de Usuario: {accountDetails.username}</p>
        </div>
        <div className="account-field">
          <p>Contraseña: {accountDetails.password}</p>
        </div>
        <div className="account-field">
          <p>Email: {accountDetails.Email}</p>
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