import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/cuenta.css';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { ToastContainer, toast } from "react-toastify";

const errorNotAClient = () => {
  toast.error("Debes ser un cliente para ver y realizar reservas", {
    pauseOnHover: false,
    autoClose: 2000,
  });
};


function AccountDetails() {
  const [accountDetails, setAccountDetails] = useState({
    Name: '',
    LastName: '',
    username: '',
    password: '*******',
    Email: ''
  });

  const { auth, logout } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    const getUser = async () => {
      if (auth.accessToken) {
        try {
          const accountId = localStorage.getItem("user_id");
          const request = await fetch(`http://localhost/user-res-api/user/${accountId}`);
          const response = await request.json();
          setAccountDetails({
            Name: response.name,
            LastName: response.last_name,
            username: response.username,
            password: '*******',
            Email: response.email
          });
        } catch (error) {
          console.error('Error al obtener los datos del cliente:', error);
        }
      } else {
        navigate('/login-cliente');
      }
    };

    getUser();
  }, [auth, navigate]);

  console.log("log de token", auth.accessToken);
  console.log("log de type", auth.userType);

  const cerrarSesion = () => {
    logout();
  };

  const reservas = () => {
    if (auth.userType === false) {
      navigate('/reservas-cliente');
    } else {
      errorNotAClient();
    }
  };

  return (
    <div className="bodyCuenta">
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
      <div className="boton-atras-container">
        <Link to="/" className="botonAtras">
          Volver a Home
        </Link>
      </div>
      <ToastContainer/>
    </div>
  );
}

export default AccountDetails;
