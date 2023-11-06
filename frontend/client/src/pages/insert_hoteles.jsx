import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/insert_hoteles.css'

function RegistrationHotel() {
  const [Email, setEmail] = useState({});
  const [Nombre, setNombre] = useState({});
  const { isLoggedAdmin } = useContext(AuthContext);
  
  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  const [formData, setFormData] = useState({
    name: '',
    description: '',
    address: '',
    city: '',
    country: '',
    images: [],
    amenities: []
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
  
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value
      }));
    }

  const RegisterHotel = () => {

    const jsonData = JSON.stringify(formData)
    console.log(jsonData)

    fetch('http://localhost:8060/hotels', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json'
    },
    body: jsonData
    })
    .then(response => response.json())
    .then(data => {
      console.log('Registro exitoso:', data);
      window.location.href = '/ver-hoteles';
    })
    .catch(error => {
      console.error('Error en el registro:', error);
      alert('Hotel no registrado');
    });
  }

return (
  <div className="registration-container" onLoad={Verificacion}>
    <h2>Registro De Hoteles</h2>
    <form onSubmit={RegisterHotel} className="registration-form">
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
      <label>
        Amenidades (ingresa amenidades separados por una coma):
        <input
          type="text"
          name="amenities"
          value={formData.amenities}
          onChange={handleChange}
          multiple size="50"
          
        />
      </label>
      <br />
      <button type="submit">Registrar Hotel</button>
    </form>
  </div>
);
}

export default RegistrationHotel;