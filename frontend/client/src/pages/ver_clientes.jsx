import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_clientes.css';

const HomePage = () => {
  const [clientes, setClientes] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getClientes = async () => {
    try {
      const request = await fetch("http://localhost:8070/user");
      const response = await request.json();
      setClientes(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getClientes();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className="bodyinicioC" onLoad={Verificacion}>
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
    </body>
  );
};

export default HomePage;