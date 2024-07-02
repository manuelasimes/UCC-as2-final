import React, { useContext, useEffect } from 'react';
import { AuthContext } from './login/auth';
import { Link} from 'react-router-dom';
import './estilo/admin_hoteles.css';

const AdminHotelesPage = () => {
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
    <div className="containerHotel">
      <div className="rectangulo">
        <h1 className="titulo">Hoteles🏨</h1>
        <div className="botones-container">
          <Link to="/agregar-hoteles" className="botonAH">
            Agregar Hoteles
          </Link>
          <Link to="/ver-hoteles" className="botonAH">
            Ver Hoteles
          </Link>
        </div>
      </div>
      <div className="boton-atras-container">
        <button onClick={handleGoBack} className="botonAtras">
          Volver Atrás
        </button>
        <Link to="/" className="botonAtras">
          Volver a Home
        </Link>
      </div>
    </div>
  );
};

export default AdminHotelesPage;
