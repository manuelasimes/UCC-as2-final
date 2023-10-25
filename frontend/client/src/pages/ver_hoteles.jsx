import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/hoteles_admin.css';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHotels = async () => {
    try {
      const request = await fetch("http://localhost:8090/admin/hoteles");
      const response = await request.json();
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getHotels();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  return (
    <body className="bodyinicioH" onLoad={Verificacion}>
      <div className="containerIniH">
        <div className="hotels-containerH">
          {hotels.length ? (
            hotels.map((hotel) => (
              <div className="hotel-cardH" key={hotel.id}>
                <img src={hotel.image} alt={hotel.nombre} className="hotel-imageH" />
                <div className="hotel-infoH">
                  <h4>{hotel.nombre}</h4>
                  <p>{hotel.email}</p>
                </div>
                <div className="hotel-descriptionH">
                    <label htmlFor={`description-${hotel.id}`}>Descripción:</label>
                    <p id={`description-${hotel.id}`}>{hotel.descripcion}</p>
                </div>
              </div>
            ))
          ) : (
            <p>No hay hoteles</p>
          )}
        </div>
      </div>
    </body>
  );
};

export default HomePage;