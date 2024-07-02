import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_reservas.css';

const VerReservas = () => {
  const [reservations, setReservations] = useState([]);
  const { auth } = useContext(AuthContext);

  const getReservations = useCallback(async () => {
    try {
      const request = await fetch(`http://localhost/user-res-api/booking`);
      const response = await request.json();
      setReservations(response);
    } catch (error) {
      console.log("No se pudieron obtener las reservas:", error);
    }
  });

  useEffect(() => {
    getReservations();
  }, []);

  return (
    <div className="reservations-container1">
      <h4>Reservas realizadas</h4>
      <div className="reservations-container2">
        {reservations.length ? (
          reservations.map((reserva) => {
            const fechaInicio = `${String(reserva.start_date).slice(6, 8)}/${String(reserva.start_date).slice(4, 6)}/${String(reserva.start_date).slice(0, 4)}`;
            const fechaFin = `${String(reserva.end_date).slice(6, 8)}/${String(reserva.end_date).slice(4, 6)}/${String(reserva.end_date).slice(0, 4)}`;
            return (
              <div className="reservation-card" key={reserva.booking_id}>
                <p>Usuario: {reserva.user_name}</p>
                <p>Hotel: {reserva.hotel_name}</p>
                <p>Fecha de llegada: {fechaInicio}</p>
                <p>Fecha de fin: {fechaFin}</p>
              </div>
            );
          })
        ) : (
          <p>No hay reservas</p>
        )}
      </div>
    </div>
  );
};

export default VerReservas;
