import React, { useContext, useEffect, useState } from 'react';
import './estilo/hoteles_admin.css';
import { Link} from 'react-router-dom';

const HomePage = () => {
  const [hotels, setHotels] = useState([]);

  const getHotels = async () => {
    try {
      const request = await fetch("http://localhost:80/search-api/searchAll=*:*");
      const response = await request.json();
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getHotels();
  }, []);;

  const handleUpdate = (id) => {

    window.location.href = `/hotel/edit/${id}`;

  }

  const handleHomeClick = () => {
    // Lógica para el botón "Home"
    window.location.href = '/';
  
  };
  
  const handleBackClick = () => {
    // Lógica para el botón "Atrás"
    window.history.back();
  
  };
  

  return (
    <div className="bodyinicioH">
      <div className="button-container2">
          <button className="botonHome" onClick={handleHomeClick}>Home</button>
          <button className="botonAtras" onClick={handleBackClick}>Atrás</button>
      </div>
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
                <button className='admin-button' onClick={() => handleUpdate(hotel.id)}>Editar hotel</button>
              </div>
            ))
          ) : (
            <p>No hay hoteles</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default HomePage;
