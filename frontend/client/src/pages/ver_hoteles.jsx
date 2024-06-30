import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/hoteles_admin.css';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);
  
  function isEmpty(str) {
    return !str.trim().length;
  }

  function isJSONEmpty(obj){
    return Object.keys(obj).length === 0;
  }

  const getHotels = async () => {
    try {
      // const request = await fetch("http://localhost:8091/cliente/hoteles");
      const request = await fetch("http://localhost:80/search-api/searchAll=*:*");
      //const request = await fetch("http://localhost:8070/hotel");
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
                <div className='img-name'>
                  <img src={hotel.images[0]} alt={hotel.name} className="hotel-imageH" />
                  <div className="hotel-infoH">
                    <h4>{hotel.name}</h4>
                  </div>
                </div>
                <div className="hotel-descriptionH">
                    <label htmlFor={`description-${hotel.id}`}>Descripción:</label>
                    <p id={`description-${hotel.id}`}>{hotel.description}</p>
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