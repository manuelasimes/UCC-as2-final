import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/insert_hoteles.css'
import { ToastContainer, toast } from "react-toastify";

function goTo(path){
  setTimeout(() => {
      window.location = window.location.origin + path;
  }, 3000)
}

const notifyRegistered = () => {
  toast.success("Actualizado!", {
      pauseOnHover: false,
      autoClose: 2000,
  })
}

const notifyError = () => {
  toast.error("Error al actualizar!", {
      pauseOnHover: false,
      autoClose: 2000,
  })
}


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
    event.preventDefault();
    const { name, value } = event.target;
  
    // Handle images and amenities input fields separately
    if (name === 'images' || name === 'amenities') {
    // Split the comma-separated values into an array
      const valueArray = value.split(',');
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: valueArray,
      }));
    } else {
    // Handle other input fields as usual
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value,
      }));
    }
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
    .then(response => {
      if (response.status === 400 || response.status === 401 || response.status === 403) {
          console.log("Error al actualizar hotel"); 

          notifyError();

          return response.json();

      } else {
          console.log("Hotel updated"); 

          notifyRegistered();

          goTo("/");

          return response.json();
      }

  })
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
}

export default RegistrationHotel;