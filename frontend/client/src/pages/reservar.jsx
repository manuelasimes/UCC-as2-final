import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import { useParams, useNavigate } from 'react-router-dom';
import './estilo/reservar.css';
import { ToastContainer, toast } from "react-toastify";
import Cookies from "universal-cookie";

const Cookie = new Cookies();

const notifyBooked = () => {
  toast.success("Reservado!", {
    pauseOnHover: false,
    autoClose: 2000,
  });
};

const notifyError = () => {
  toast.error("Hotel no disponible para reserva en fecha seleccionada!", {
    pauseOnHover: false,
    autoClose: 2000,
  });
};

function convertirFecha(fecha) {
  let fechaString = fecha.toString();
  let year = fechaString.substring(0, 4);
  let month = fechaString.substring(5, 7);
  let day = fechaString.substring(8, 10);
  let yearPlusMonth = year.concat("", month);
  let fechaStringFinal = yearPlusMonth.concat("", day);
  var fechaEntero = Number(fechaStringFinal);

  console.log(fechaEntero);

  return fechaEntero;
}

const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const accountId = localStorage.getItem("user_id");

  const handleReserva = async (e) => {
    e.preventDefault();

    const formData = {
      booked_hotel_id: hotelId,
      user_booked_id: parseInt(accountId),
      start_date: convertirFecha(startDate),
      end_date: convertirFecha(endDate),
    };

    console.log("formData", formData);

    fetch('http://localhost/user-res-api/booking', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    }).then((response) => {
      if (response.status === 400 || response.status === 401 || response.status === 403) {
        console.log("Error al reservar");

        notifyError();

        return response.json();
      } else {
        console.log("Booking added");

        notifyBooked();

        return response.json();
      }
    });
  };

  useEffect(() => {

    setHotelData('');
    if (hotelId) {
      fetch(`http://localhost/hotels-api/hotels/${hotelId}`)
        .then(response => response.json())
        .then(data => {
          if (data.images && typeof data.images === 'string') {
            data.images = JSON.parse(data.images); // Convert the string to an array
          }
          console.log("Todos los datos del hotel", data);
          console.log("imagenes del hotel", data.images);
          setHotelData(data);
        })
        .catch(error => {
          console.error('Error al obtener datos del hotel: ', error);
        });
    }
  }, [hotelId]);

  const handleStartDateChange = (event) => {
    const selectedStartDate = event.target.value;
    setStartDate(selectedStartDate);
    const startDateObj = new Date(selectedStartDate);
    const endDateObj = new Date(endDate);
    if (endDate && startDateObj > endDateObj) {
      setEndDate('');
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const handleEndDateChange = (event) => {
    const selectedEndDate = event.target.value;
    setEndDate(selectedEndDate);
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(selectedEndDate);
    if (startDate && startDateObj > endDateObj) {
      setStartDate('');
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const filterHotels = async () => {
    const startDateParsed = convertirFecha(startDate);
    const endDateParsed = convertirFecha(endDate);

    const request = await fetch(`http://localhost/user-res-api/hotel/availability/${hotelId}/${startDateParsed}/${endDateParsed}`);

    const response = await request.json();
    if (response === 0) {
      setEndDate('');
      notifyError();
    }
  };

  const handleVolver = () => {
    window.history.back();
  };

  const today = new Date().toISOString().split('T')[0];

  return (
    <div className="bodyReserva">
      <div>
        {typeof hotelData === 'undefined' ? (
          <>CARGANDO...</>
        ) : (
          <div className="container45">
            <div className="informacion">
              <div className="cuadroImag"><img src={hotelData.images} alt={hotelData.name} className="tamanoImag" /></div>
              <div className="descripcion">{hotelData.description}</div>
            </div>
            <div className="reserva-form">
              <h6>Realice reserva del Hotel</h6>
              <h6>{hotelData["nombre"]}</h6>
              <form onSubmit={handleReserva}>
                <div className="form-group">
                  <label htmlFor="fechaInicio">Fecha de inicio:</label>
                  <input
                    type="date"
                    id="fechaInicio"
                    min={today}
                    value={startDate}
                    onChange={handleStartDateChange}
                    required
                  />
                </div>
                <div className="form-group">
                  <label htmlFor="fechaFin">Fecha de fin:</label>
                  <input
                    type="date"
                    id="fechaFin"
                    min={today}
                    value={endDate}
                    onChange={handleEndDateChange}
                    required
                  />
                </div>
                <div>
                  <button type="submit" className="confReserva">Confirmar</button>
                  <button type="button" className="confReserva" onClick={handleVolver}>Volver</button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
      <ToastContainer />
    </div>
  );
};

export default ReservaPage;
