import React from 'react';
import { Link } from 'react-router-dom';
import './estilo/opciones_admin.css';

const OpcionesAdminPage = () => {
  return (
    <div className= "bodyinicioADM">
     
    <div className="container">
    <div className= "cuadradointerno">
       <h1 className="titulo">Opciones</h1>
        <div className="botones-container">
        <Link to="/administrar-hoteles" className="botonAD">
          Administrar Hoteles
        </Link>
        <Link to="/administrar-clientes" className="botonAD">
          Administrar Clientes
        </Link>
        <Link to="/infraestructura" className="botonAD">
          Administrar Infraestructura
        </Link>
        </div>
      </div>
      </div>
    </div>
  );
};

export default OpcionesAdminPage;
