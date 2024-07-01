import React, { useContext, useEffect } from 'react';
import { AuthContext } from './login/auth';
import { Link, useNavigate } from 'react-router-dom';
import './estilo/admin_hoteles.css';

const AdminHotelesPage = () => {
  const navigate = useNavigate();


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
    </div>
  );
};

export default AdminHotelesPage;
