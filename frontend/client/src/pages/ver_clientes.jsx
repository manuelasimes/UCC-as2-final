import React, { useContext, useEffect, useState } from 'react';
import './estilo/ver_clientes.css';
import { Link} from 'react-router-dom';

const HomePage = () => {
  const [clientes, setClientes] = useState([]);

  const getClientes = async () => {
    try {
      const request = await fetch("http://localhost/user-res-api/user");
      const response = await request.json();
      setClientes(response);
    } catch (error) {
      console.log("No se pudieron obtener los clientes:", error);
    }
  };

  useEffect(() => {
      getClientes();
  }, []);

  const handleGoBack = () => {
    // Función para ir atrás en la historia del navegador
    window.history.back();
  };

  return (
    <div className="bodyinicioC">
      <div className="containerIniC">
        <div className="hotels-containerC">
          {clientes.length ? (
            clientes.map((cliente) => (
              <div className="hotel-cardC" key={cliente.id}>
                <div className="hotel-infoC">
                  <h4>{cliente.name}</h4>
                  <p>{cliente.last_name}</p>
                </div>
                <div className="hotel-infoC">
                  <p>{cliente.username}</p>
                  <p>{cliente.email}</p>
                </div>
              </div>
            ))
          ) : (
            <p>No hay clientes</p>
          )}
        </div>
      </div>
      <div className="boton-atras-container">
            <Link to="/" className="botonAtras">
              Volver a Home
            </Link>
            <button onClick={handleGoBack} className="botonAtras">
            Volver Atrás
            </button>
      </div>
    </div>
  );
};

export default HomePage;
