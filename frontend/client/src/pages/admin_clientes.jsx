import React, { useContext, useEffect } from 'react';
import { AuthContext } from './login/auth';
import { Link, useNavigate } from 'react-router-dom';
import './estilo/admin_clientes.css';

const AdminClientesPage = () => {
  const { auth } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (auth.userType !== true) {
      navigate('/login-admin');
    }
  }, [auth, navigate]);

  return (
    <div className="container" >
      <div className= "rectangulo1">
        <h1 className="titulo">Clientes👥</h1>
        <div className="botones-container">
          <Link to="/ver-reservas" className="botonAC">
            Ver Reservas
          </Link>
          <Link to="/ver-clientes" className="botonAC">
            Ver Clientes
          </Link>
        </div>
      </div>
    </div>
  );
};

export default AdminClientesPage;
