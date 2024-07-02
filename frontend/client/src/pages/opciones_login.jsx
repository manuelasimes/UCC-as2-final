import React from 'react';
import './estilo/opciones_login.css';

const RoleSelection = () => {
  const handleLoginAdmin = () => {
    window.location.href = 'http://localhost:3000/login-admin';
  };

  const handleLoginCliente = () => {
    window.location.href = 'http://localhost:3000/login-cliente';
  };

  return (
  <div className= "bodyopciones">
    <div className="selection">
      
      <div className="contOPcio">
        <h2 className="welcome-text">Bienvenido! Elija una opción</h2>
        <div className="buttons-container">
          <button className="buttoninicio" variant="contained" size="large" onClick={handleLoginCliente}>
            Cliente
          </button>
          <button className="buttoninicio" variant="contained" size="large" onClick={handleLoginAdmin}>
            Administrador
          </button>
        </div>
      </div>
      </div>
   </div>
    
  );
};

export default RoleSelection;

