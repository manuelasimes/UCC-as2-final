import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/ver_reservas.css';

const VerReservas = () => {
  const [reservations, setReservations] = useState([]);
  const [hoteles, setHoteles] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);

  const getHoteles = useCallback(async () => {
    try {
      const hotelesArray = [];
      for (let i = 0; i < reservations.length; i++) {
        const reserva = reservations[i];
        const request = await fetch(`http://localhost:8090/cliente/hotel/${reserva.hotel_id}`);
        const response = await request.json();
        hotelesArray.push(response);
      }
      setHoteles(hotelesArray);
    } catch (error) {
      console.log("No se pudieron obtener los hoteles:", error);
    }
  }, [reservations]);

  const getReservations = useCallback(async () => {
    if (isLoggedAdmin) {
      try {
        const request = await fetch(`http://localhost:8090/admin/reservas`);
        const response = await request.json();
        setReservations(response);
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
    } else {
      window.location.href = '/';
    }
  }, [isLoggedAdmin]);

  useEffect(() => {
    getReservations();
  }, [getReservations]); // Se elimina la dependencia de getReservations

  useEffect(() => {
    getHoteles();
  }, [getHoteles]); // Se agrega getHoteles como dependencia separada

  return (
  
    <div className="reservations-container1">
      <h4>Recervas realizadas</h4>
      <div className="reservations-container2">
            {reservations.length ? (
              reservations.map((reserva) => {
                const hotel = hoteles.find((hotel) => hotel.id === reserva.hotel_id);
                const fechaInicio = `${reserva.dia_inicio}/${reserva.mes_inicio}/${reserva.anio_inicio}`;
                const fechaFin = `${reserva.dia_final}/${reserva.mes_final}/${reserva.anio_final}`;
                return (
                  <div className="reservation-card" key={reserva.ID}>
                    <p>Hotel: {hotel ? hotel.nombre : 'Hotel desconocido'}</p>
                    <p>Fecha de llegada: {fechaInicio}</p>
                    <p>Fecha de fin: {fechaFin}</p>
                    <p>Gracias por elegirnos!</p>
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