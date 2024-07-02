import React, { useContext, useState, useEffect } from 'react';
import { AuthContext } from './login/auth';
import './estilo/insert_hoteles.css';
import { ToastContainer, toast } from "react-toastify";

const notifyRegistered = () => {
  toast.success("Actualizado!", {
    pauseOnHover: false,
    autoClose: 2000,
  });
};

const notifyError = () => {
  toast.error("Error al actualizar!", {
    pauseOnHover: false,
    autoClose: 2000,
  });
};

const RegistrationHotel = () => {
  const { auth } = useContext(AuthContext);
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    address: '',
    city: '',
    country: '',
    images: [],
    amenities: []
  });

  const { accessToken } = auth;

  useEffect(() => {
    if (!accessToken) {
      window.location.href = '/login-admin'; // Redirigir si no hay accessToken de autenticación
    }
  }, [accessToken]);

  const handleChange = (event) => {
    event.preventDefault();
    const { name, value } = event.target;

    var flag = false;

    if (name === 'city' || name === 'country') {

      flag = true;

      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value.toLowerCase(),
      }));

   } 

    if (name === 'images' || name === 'amenities') {
      const valueArray = value.split(',');
      flag = true;
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: valueArray,
      }));
    } else if (!flag) {
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value,
      }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();



    try {
      const response = await fetch('http://localhost/hotels-api/hotels', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          // 'Authorization': `Bearer ${accessToken}`
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        notifyError();
      } else {
        notifyRegistered();
        setTimeout(() => {
          window.location = window.location.origin; // Redirigir a la página principal después del registro
        }, 2000);
      }
    } catch (error) {
      console.error("Error al registrar el hotel:", error);
      notifyError();
    }
  };

  return (
    <div className="registration-container">
      <h2>Registro De Hoteles</h2>
      <form onSubmit={handleSubmit} className="registration-form">
        <label>
          Nombre:
          <input
            type="text"
            name="name"
            value={formData.name}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Descripcion:
          <input
            type="text"
            name="description"
            value={formData.description}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Direccion:
          <input
            type="text"
            name="address"
            value={formData.address}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Ciudad:
          <input
            type="text"
            name="city"
            value={formData.city}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Pais:
          <input
            type="text"
            name="country"
            value={formData.country}
            onChange={handleChange}
          />
        </label>
        <br />
        <label>
          Imagenes (ingresa urls separados por una coma):
          <input
            type="text"
            name="images"
            value={formData.images}
            onChange={handleChange}
            multiple size="50"
          />
        </label>
        <br/>
        <label>
          Amenities (ingresa amenities separados por una coma):
          <input
            type="text"
            name="amenities"
            value={formData.amenities}
            onChange={handleChange}
            multiple size="50"
          />
        </label>
        <br/>
        <button type="submit">Registrar Hotel</button>
      </form>
      <ToastContainer />
    </div>
  );
};

export default RegistrationHotel;
