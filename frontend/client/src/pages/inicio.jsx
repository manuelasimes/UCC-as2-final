import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/inicio.css';
import swal from 'sweetalert2';
import Cookies from 'universal-cookie';
import { useNavigate } from 'react-router-dom';

const Cookie = new Cookies();

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [city, setCity] = useState('');
  const { auth, logout } = useContext(AuthContext);

  const navigate = useNavigate();

  function isEmpty(str) {
    return !str.trim().length;
  }

  function isJSONEmpty(obj) {
    return Object.keys(obj).length === 0;
  }

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
  }, []);

  const Verificacion = (hotelId) => {
      navigate(`/reservar/${hotelId}`);
  };

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    const selectedStartDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no válidas");
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const selectedStartDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no válidas");
    }
  };

  const handleCityChange = (event) => {
    setCity(event.target.value);
  };

  const filterHotels = async () => {
    var endDateValue = document.getElementById('end-date').value;
    var startDateValue = document.getElementById('start-date').value;
    
    var cityInLowerCase = city.toLowerCase();

    if (isEmpty(city)) {
      alert("No está permitido buscar solo por fecha, debe ingresar un destino!");
      getHotels();
    } else {
      let request;
      if (!endDateValue || !startDateValue) {
        request = await fetch(`http://localhost/search-api/search=city_${cityInLowerCase}`);
      } else {
        request = await fetch(`http://localhost/search-api/search=city_${cityInLowerCase}_${startDate}_${endDate}`);
      }
      const response = await request.json();
      if (response !== null) {
        const checkResponse = JSON.stringify(response);
        console.log(checkResponse);
        if (isJSONEmpty(response)) {
          alert("No se encontraron hoteles en esa ubicación");
          getHotels();
        } else {
          setHotels(response);
        }
      } else {
        alert("No se encontraron hoteles en esa ubicación");
        getHotels();
      }
    }
  };

  const Cuenta = () => {
    if (auth.accessToken) {
      navigate('/cuenta');
    } else {
      navigate('/login-cliente');
    }
  };

  const Admin = () => {
    if (auth.userType === false) {
      swal.fire({
        customClass: {
          popup: 'popup-custom',
          confirmButton: 'confButton-custom'
        },
        title: "Ingresa como admin",
        text: "Los clientes no pueden acceder al área de administración, por lo que deberás cerrar la sesión del usuario actual e ingresar como un usuario administrador. A continuación, puedes cerrar sesión y serás redirigido al login de administrador o puedes continuar como usuario regular haciendo click por fuera del recuadro.",
        confirmButtonText: 'Ingresar como admin',
        icon: "warning",
        padding: "20px",
        timerProgressBar: "true",
        allowOutsideClick: "true"
      }).then(response => {
        if (response.isConfirmed) {
          logout();
          navigate('/login-admin');
        }
      });
    } else if (auth.userType !== true) {
      navigate('/login-admin');
    } else {
      navigate('/admin');
    }
  };

  return (
    <body className="bodyinicio">
      <div className="header-content">
        <div className="cuenta-button-container">
          <button className="cuenta-button" onClick={Cuenta}>
            Tu Cuenta
          </button>
        </div>
        <div className="admin-button-container">
          <button className="admin-button" onClick={Admin}>
            Admin
          </button>
        </div>
      </div>
      <div className="contdeFechas">
        <div className="localidad">
          <label className="fecha">¿A dónde vas?</label>
          <input type="text" id="destino" placeholder="Su destino" value={city} onChange={handleCityChange} />
        </div>
        <div className="date-pickerINI1">
          <label htmlFor="start-date" className="fecha">Entrada</label>
          <input type="date" id="start-date" value={startDate} onChange={handleStartDateChange} />
        </div>
        <div className="date-pickerINI1">
          <label htmlFor="end-date" className="fecha">Salida</label>
          <input type="date" id="end-date" value={endDate} onChange={handleEndDateChange} />
        </div>
        <button className="botbusquedaFec" onClick={filterHotels}>Buscar</button>
      </div>
      <div className="containerIni">
        <div className="hotels-container">
          {hotels.length ? (
            hotels.map((hotel) => (
              <div className='hotel-card' key={hotel.id}>
                <img src={hotel.images[0]} alt={hotel.name} className="hotel-image" />
                <div className="hotel-info">
                  <h4>{hotel.name}</h4>
                  <p>{hotel.description}</p>
                  {console.log(hotel.images[0])}
                  <button onClick={() => Verificacion(hotel.id)}>
                    Reservar
                  </button>
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
