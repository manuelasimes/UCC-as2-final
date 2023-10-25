import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import { useParams } from 'react-router-dom';
import './estilo/reservar.css';

const ReservaPage = () => {
  const { hotelId } = useParams();
  const [hotelData, setHotelData] = useState('');
  const { isLoggedCliente } = useContext(AuthContext);
  const [startDate, setStartDate] = useState('');
  const [endDate, setEndDate] = useState('');
  const accountId = localStorage.getItem("id_cliente");

  const Verificacion = () => {
    if (!isLoggedCliente) {
      window.location.href = '/login-cliente';
    }
  };

  const handleReserva = () => {
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(endDate);
    const Dias = Math.round((endDateObj - startDateObj) / (1000 * 60 * 60 * 24));
    const formData = {
      hotel_id: parseInt(hotelId),
      cliente_id: parseInt(accountId),
      anio_inicio: startDateObj.getFullYear(),
      anio_final: endDateObj.getFullYear(),
      mes_inicio: startDateObj.getMonth() + 1, 
      mes_final: endDateObj.getMonth() + 1, 
      dia_inicio: startDateObj.getDate() + 1,
      dia_final: endDateObj.getDate() + 1,
      dias: Dias
    };

    fetch('http://localhost:8090/cliente/reserva', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
    })
      .then(response => response.json())
      .then(data => {
        console.log('Registro exitoso:', data);
        alert(JSON.stringify(formData));
        /*window.location.href = '/';*/
      })
      .catch(error => {
        console.error('Error en el registro:', error);
        alert('Credenciales incorrectas');
      });
  };

  useEffect(() => {
    setHotelData('');
    if (hotelId) {
      fetch(`http://localhost:8090/cliente/hotel/${hotelId}`)
        .then(response => response.json())
        .then(data => {
          setHotelData(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [hotelId]);

  const handleStartDateChange = (event) => {
    setStartDate(event.target.value);
    const startDateObj = new Date(event.target.value);
    const endDateObj = new Date(endDate);
    if (startDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no válidas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const handleEndDateChange = (event) => {
    setEndDate(event.target.value);
    const startDateObj = new Date(startDate);
    const endDateObj = new Date(event.target.value);
    if (startDateObj > endDateObj) {
      setEndDate('');
      alert("Fechas no válidas");
    }
    if (startDate && endDate) {
      filterHotels();
    }
  };

  const filterHotels = async () => {
    const request = await fetch(`http://localhost:8090/cliente/disponibilidad/${hotelId}/${startDate}/${endDate}`);
    const response = await request.json();
    if (response === 0) {
      setEndDate('');
      alert("No hay habitaciones disponibles para esas fechas");
    }
  };

  const handleVolver = () => {
    window.history.back();
  };

  return (
    <div className="bodyReserva">
      <div>
        {typeof hotelData === 'undefined' ? (
          <>CARGANDO...</>
        ) : (
          <div className="container45" onLoad={Verificacion}>
            <div className="informacion">
              <div className="cuadroImag"><img src={hotelData.image} alt={hotelData.nombre} className="tamanoImag" /></div>
              <div className="descripcion">{hotelData["descripcion"]}</div>
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
    </div>
  );
};

export default ReservaPage;
