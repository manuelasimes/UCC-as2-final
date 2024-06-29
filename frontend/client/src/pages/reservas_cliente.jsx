import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);

  const getReservations = useCallback(async () => {
    if (isLoggedCliente) {
      const accountId = localStorage.getItem("id_cliente");
      try {
        const request = await fetch(`http://localhost/user-res-api/booking/user/${accountId}`);
        const response = await request.json();
        setReservations(response);
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
    } else {
      window.location.href = '/';
    }
  }, [isLoggedCliente]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  return (
    <div className="reservations-container1">
      <h4>Mis reservas</h4>
      <div className="reservations-container2">
        {reservations.length ? (
          reservations.map((reservation) => {
            const fechaInicio = `${String(reservation.start_date).slice(6, 8)}/${String(reservation.start_date).slice(4, 6)}/${String(reservation.start_date).slice(0, 4)}`;
            const fechaFin = `${String(reservation.end_date).slice(6, 8)}/${String(reservation.end_date).slice(4, 6)}/${String(reservation.end_date).slice(0, 4)}`;
            return (
              <div className="reservation-card" key={reservation.booking_id}>
                <p>Hotel: {reservation.hotel_name}</p>
                <p>Fecha de llegada: {fechaInicio}</p>
                <p>Fecha de fin: {fechaFin}</p>
                <p>Gracias por elegirnos!</p>
              </div>
            );
          })
        ) : (
          <p>No tienes reservas</p>
        )}
      </div>
    </div>
  );
};

export default HomePage;
