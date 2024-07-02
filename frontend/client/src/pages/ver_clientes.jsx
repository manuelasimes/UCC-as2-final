import React, { useContext, useEffect, useState } from 'react';
import './estilo/ver_clientes.css';

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
    </div>
  );
};

export default HomePage;
