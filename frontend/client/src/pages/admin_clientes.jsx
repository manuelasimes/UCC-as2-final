import React, { useContext, useEffect } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/admin_clientes.css';

const AdminClientesPage = () => {
  const { auth } = useContext(AuthContext);

  useEffect(() => {
    if (auth.userType !== true) {
      // Redirige si no es un usuario autorizado
      window.location.href = '/login-admin';
    }
  }, [auth]);

  const handleGoBack = () => {
    // Función para ir atrás en la historia del navegador
    window.history.back();
  };

  return (
    <div className="container">
      <div className="rectangulo1">
        <div className="tituloAD">Clientes👥</div>
        <div className="botones-container">
          <Link to="/ver-reservas" className="botonAC">
            Ver Reservas
          </Link>
          <Link to="/ver-clientes" className="botonAC">
            Ver Clientes
          </Link>
        </div>
      </div>
      <div className="boton-atras-container">
        <button onClick={handleGoBack} className="botonAtras">
          Volver
        </button>
        <Link to="/" className="botonAtras">
          Volver a Home
        </Link>
      </div>
    </div>
  );
};

export default AdminClientesPage;
