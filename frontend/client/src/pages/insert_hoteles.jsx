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
    nombre: '',
    descripcion: '',
    email: '',
    imagen: '',
    cant_hab: ''
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
  
    if (name === "cant_hab" && value !== '') {
      const intValue = parseInt(value);
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: intValue
      }));
    } else {
      setFormData((prevFormData) => ({
        ...prevFormData,
        [name]: value
      }));
    }
  };

  useEffect(() => {
    setEmail('');
  
    if (formData.email) {
      fetch(`http://localhost:8090/admin/hotel/email/${formData.email}`)
        .then(response => response.json())
        .then(data => {
            setEmail(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [formData.email]);

  useEffect(() => {
    setNombre('');
  
    if (formData.nombre) {
      fetch(`http://localhost:8090/admin/hotel/nombre/${formData.nombre}`)
        .then(response => response.json())
        .then(data => {
            setNombre(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [formData.nombre]);

  const RegisterHotel = () => {
    if (formData.email === Email.email) {
      alert('El email ya pertenece a un hotel');
    }
    else if (formData.nombre === Nombre.nombre) {
      alert('El nombre no esta disponible');
    }
    else
    {
      fetch('http://localhost:8090/admin/hotel', {
      method: 'POST',
      headers: {
      'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
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
  };

  return (
    <div className="registration-container" onLoad={Verificacion}>
      <h2>Registro De Hoteles</h2>
      <form onSubmit={RegisterHotel} className="registration-form">
        <label>
          Nombre:
          <input
            type="text"
            name="nombre"
            value={formData.nombre}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Descripcion:
          <input
            type="text"
            name="descripcion"
            value={formData.descripcion}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Email:
          <input
            type="text"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
         Imagen:
          <input
            type="text"
            name="imagen"
            value={formData.imagen}
            onChange={handleChange}
          />
        </label>
        <br />
        <label>
          Cantidad de habitaciones:
          <input
            type="text"
            name="cant_hab"
            value={formData.cant_hab}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <button type="submit">Registrar Hotel</button>
      </form>
    </div>
  );
}

export default RegistrationHotel;