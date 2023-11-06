import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/inicio.css';
import swal from 'sweetalert2';
import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";

const Cookie = new Cookies();

const HomePage = () => {
  const [hotels, setHotels] = useState([]);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const [city, setCity] = useState('');
  const [idHotelEdit, setIdHotelEdit] = useState('1');
  const { isLoggedCliente } = useContext(AuthContext);
  const { isLoggedAdmin } = useContext(AuthContext);
  const { logout } = useContext(AuthContext);

  const navigate = useNavigate();

  console.log(Cookie.get("user_id"))

  function isEmpty(str) {
    return !str.trim().length;
  }

  function isJSONEmpty(obj){
    return Object.keys(obj).length === 0;
  }

  const getHotels = async () => {
    try {
      // const request = await fetch("http://localhost:8090/cliente/hoteles");
       const request = await fetch("http://localhost:8090/searchAll=*:*");
      // const request = await fetch("http://localhost:8070/hotel");
      const response = await request.json();
      console.log(response)
      setHotels(response);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  };

  useEffect(() => {
    getHotels();
  }, []);

  const Verificacion = (hotelId) => {
    if (Cookie.get("user_id") !== -1) {
      window.location.href = `/reservar/${hotelId}`;
    }
    else
    {
      window.location.href = '/login-cliente';
    }
  };

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    const selectedStartDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const selectedStartDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (selectedStartDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no validas");
    }
  };

  const handleCityChange = (event) => {
    setCity(event.target.value);
  }

  const filterHotels = async () => {

      var endDateValue = document.getElementById('end-date').value;
      var startDateValue = document.getElementById('start-date').value;

      if (isEmpty(city)) {
        alert("No esta permitido buscar solo por fecha, debe ingresar un destino!")
        getHotels();
      } else {
        if (!endDateValue || !startDateValue ){
          const request = await fetch(`http://localhost:8090/search=city_${city}`);
          const response = await request.json();
          if (response !== null) {
            const checkResponse = JSON.stringify(response)
            console.log(checkResponse)
          if (isJSONEmpty(response)) {
            alert("No se encontraron hoteles en esa ubicacion")
            getHotels();
          } else {
            setHotels(response);
          }
          } 
        } else {
        const request = await fetch(`http://localhost:8090/search=city_${city}_${startDate}_${endDate}`);
        const response = await request.json();
        if (response !== null) {
          const checkResponse = JSON.stringify(response)
          console.log(checkResponse)
          if (isJSONEmpty(response)) {
            alert("No se encontraron hoteles en esa ubicacion disponibles en esa fecha")
            getHotels();
          } else {
            setHotels(response);
          }
          } else {
          alert("No se encontraron hoteles en esa ubicacion disponibles en esa fecha")
          getHotels();
          }
        }
      }
    };

  const Admin = () => {
    if (isLoggedCliente){
      swal.fire ({
        customClass: {
          popup: 'popup-custom',
          confirmButton: 'confButton-custom'
        },
        title: "Ingresa como admin",
        //html: '<div><p class="textAlert">Los clientes no pueden acceder al area de administracion, por lo que deberas cerrar la sesion del usuario actual e ingresar como un usuario administrador. A continuacion puedes cerrar sesion y seras redirigido al login de administrador o puedes continuar como usuario regular.</p></div>',
        text:"Los clientes no pueden acceder a el area de administracion, por lo que deberas cerrar la sesion del usuario actual e ingresar como un usuario administrador. A continuacion puedes cerrar sesion y seras redirigido al login de administrador o puedes continuar como usuario regular haciendo click por fuera del recuadro.",
        confirmButtonText: 'Ingresar como admin',
        icon: "warning",
        padding: "20px",
        timerProgressBar: "true",
        allowOutsideClick: "true"
      }).then(response => {
        if(response.isConfirmed){
          logout();
          window.location.href = '/login-admin';
        }
      });
    }
    else if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
    else
    {
      window.location.href = '/admin';
    }
  }

  const Cuenta = () => {
    if (isLoggedAdmin || isLoggedCliente) {
      window.location.href = '/cuenta';
    }
    else
    {
      window.location.href = '/login-cliente'
    }
  }

  console.log(Cookie.get("user_type"))

  const editHotel = (id) => {
    window.location.href = `/hotel/edit/${id}`
  }

  const renderButton= (id) => (
    <>
    <button onClick={ () => editHotel(id) }>Editar</button>
    </>
  )
  

  return (
    <body className= "bodyinicio">
      <div className="header-content">
        <div className="cuenta-button-container">
          <button className="cuenta-button" onClick={Cuenta}>
            Tu Cuenta
          </button>
        </div>
      </div>
        <div className="contdeFechas">
        <div className="localidad">
            <label className="fecha">¿a dónde vas?</label>
            <input type="text" id="destino" placeholder="Su destino" value={city} onChange={handleCityChange}/>
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
            {hotels.length ? 
              ( hotels.map((hotel) => (
                <div className='hotel-card' key={hotel.id}>
                  <img src={hotel.images[0]} alt={hotel.name} className="hotel-image" />
                  <div className="hotel-info">
                    <h4>{hotel.name}</h4>
                    <p>{hotel.description} </p>
                    {console.log(hotel.images[0])}
                    { Cookie.get("user_type") === true ? renderButton(hotel.id) : null } 
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
