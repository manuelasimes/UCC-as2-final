import React, { useEffect, useState } from 'react';
import '../estilo/Register_cliente.css';

function goTo(path){
  setTimeout(() => {
      window.location = window.location.origin + path;
  }, 0)
}

function RegistrationPage() {
  const [clienteEmail, setClienteEmail] = useState({});
  const [clienteUsername, setClienteUsername] = useState({});
  const [error, setError] = useState("");

  const [formData, setFormData] = useState({
    name: "",
    last_name: "",
    username: "",
    password: "",
    email: ""
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const Register = (event) => {
    event.preventDefault();

    fetch('http://localhost:8070/user', {
    method: 'POST',
    headers: {
    'Content-Type': 'application/json'
    },
    body: JSON.stringify(formData)
    })
    .then(response => {
      console.log(response)
      if (response.status === 400)
      {
        alert('Se a ingresado un nombre de usuario o un email que ya se encuentra registrado.');
      }
      if (response.ok) {
        return response.json();
      } else {
        throw new Error('Error en el registro');
      }
    })
    .then(data => {
      // Maneja la respuesta JSON aquí si es necesario
      console.log('Datos de respuesta:', data);
      goTo('/login-cliente');
    })
    .catch(error => {
      console.error('Error en el registro:', error);
      alert('Credenciales incorrectas');
    });
  };

  /*function RegistrationPage() {
    const [clienteEmail, setClienteEmail] = useState("");
    const [clienteUsername, setClienteUsername] = useState("");
    const [formData, setFormData] = useState({
      name: "",
      last_name: "",
      username: "",
      password: "",
      email: ""
    });
    const [error, setError] = useState("");
  
    const handleChange = (e) => {
      setFormData({ ...formData, [e.target.name]: e.target.value });
    };
  
    const handleRegister = async () => {
      try {
        // Comprobar disponibilidad de email
        const responseEmail = await fetch(`http://localhost:8070/user?email=${formData.email}`);
        if (responseEmail.ok) {
          const dataEmail = await responseEmail.json();
          setClienteEmail(dataEmail.email);
        }
  
        // Comprobar disponibilidad de username
        const responseUsername = await fetch(`http://localhost:8070/user?username=${formData.username}`);
        if (responseUsername.ok) {
          const dataUsername = await responseUsername.json();
          setClienteUsername(dataUsername.username);
        }
  
        // Realizar registro si no hay conflictos
        if (formData.email === clienteEmail && formData.username === clienteUsername) {
          setError('El email y el username ya están en uso');
        } else if (formData.email === clienteEmail) {
          setError('El email ya pertenece a una cuenta');
        } else if (formData.username === clienteUsername) {
          setError('El username no está disponible');
        } else {
          const response = await fetch('http://localhost:8070/user', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
          });
          if (response.ok) {
            const data = await response.json();
            console.log('Registro exitoso:', data);
            window.location = '/login-cliente';
          } else {
            throw new Error('Error en el registro');
          }
        }
      } catch (error) {
        console.error('Error en el registro:', error);
        setError('Credenciales incorrectas');
      }
    };*/

  return (
    <div className="registration-container">
      <h2>Registro</h2>
      <form onSubmit={Register} className="registration-form">
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
          Apellido:
          <input
            type="text"
            name="last_name"
            value={formData.last_name}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Nombre de usuario:
          <input
            type="text"
            name="username"
            value={formData.username}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Contraseña:
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        <label>
          Email:
          <input
            type="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
            required
          />
        </label>
        <br />
        {error && <p className="error-message">{error}</p>}
        <button type="submit">Registrarse</button>
      </form>
    </div>
  );
}

export default RegistrationPage;
