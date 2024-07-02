import React, { useContext, useEffect, useState, useCallback } from 'react';
import { AuthContext } from './login/auth';
import './estilo/reservas_cliente.css';
import { useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from "react-toastify";

const HomePage = () => {
  const [reservations, setReservations] = useState([]);
  const { auth } = useContext(AuthContext);
  const navigate = useNavigate();

  const getReservations = useCallback(async () => {
    if (auth.accessToken && auth.userType === false) { // userType === false significa que es un cliente
      const accountId = localStorage.getItem("user_id");
      try {
        const request = await fetch(`http://localhost/user-res-api/booking/user/${accountId}`);
        const response = await request.json();
        setReservations(response);
      } catch (error) {
        console.log("No se pudieron obtener las reservas:", error);
      }
    } else {
      navigate('/');
    }
  }, [auth, navigate]);

  useEffect(() => {
    getReservations();
  }, [getReservations]);

  const handleDeleteReserva =  async (bookingId) => {

    try {
      const response =  await fetch(`http://localhost/user-res-api/booking/${bookingId}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          // Add authorization headers if needed
          // 'Authorization': `Bearer ${accessToken}`
        },
      });
  
      if (!response.ok) {
        throw new Error('Error deleting reservation');
      }
  
      // Optionally update the UI after a successful deletion
      setReservations((prevReservations) => 
        prevReservations.filter((reservation) => reservation.booking_id !== bookingId)
      );
  
      // Notify success
      toast.success("Reservation deleted successfully!", {
        pauseOnHover: false,
        autoClose: 2000,
      });
    } catch (error) {
      console.error("Error deleting reservation:", error);
      toast.error("Error deleting reservation!", {
        pauseOnHover: false,
        autoClose: 2000,
      });
    }
  };

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
                <button className='admin-button' onClick={() => handleDeleteReserva(reservation.booking_id)}>Cancelar Reserva</button>
                <p>Gracias por elegirnos!</p>
              </div>
            );
          })
        ) : (
          <p>No tienes reservas</p>
        )}
      </div>
      <ToastContainer />
    </div>
  );
};

export default HomePage;
