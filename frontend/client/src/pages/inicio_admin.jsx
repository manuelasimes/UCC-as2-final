import React, { useContext } from 'react';
import { AuthContext } from './login/auth';
import { Link } from 'react-router-dom';
import './estilo/opciones_admin.css';

const OpcionesAdminPage = () => {
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className= "bodyinicioADM">
     
    <div className="container" onLoad={Verificacion}>
    <div className= "cuadradointerno">
       <h1 className="titulo">Opciones</h1>
        <div className="botones-container">
        <Link to="/administrar-hoteles" className="botonAD">
          Administrar Hoteles
        </Link>
        <Link to="/administrar-clientes" className="botonAD">
          Administrar Clientes
        </Link>
        <Link to="/insfraestructura" className="botonAD">
          Administrar Infraestructura
        </Link>
        </div>
      </div>
      </div>
    </body>
  );
};

export default OpcionesAdminPage;
