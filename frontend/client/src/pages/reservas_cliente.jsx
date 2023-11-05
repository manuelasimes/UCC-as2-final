import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedCliente } = useContext(AuthContext);

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        const request = await fetch(`http://localhost:8070/cliente/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push(response);
      }
      setHoteles(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations]);

  const getReservations = useCallback(async () => {
    if (isLoggedCliente) {
      const accountId = localStorage.getItem("id_cliente");
      try {
        const request = await fetch(`http://localhost:8070/booking/user/${accountId}`);
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
  }, [getReservations]); // Se elimina la dependencia de getReservations

  useEffect(() => {
    getHoteles();
  }, [getHoteles]); // Se agrega getHoteles como dependencia separada

  return (
    <div className="reservations-container1">
      <h4>Mis reservas</h4>
      <div className="reservations-container2">
        {reservations.length ? (
          reservations.map((reservation) => {
            const hotel = hoteles.find((hotel) => hotel.id === reservation.hotel_id);
            const fechaInicio = `${reservation.dia_inicio}/${reservation.mes_inicio}/${reservation.anio_inicio}`;
            const fechaFin = `${reservation.dia_final}/${reservation.mes_final}/${reservation.anio_final}`;
            return (
              <div className="reservation-card" key={reservation.ID}>
                <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
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