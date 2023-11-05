import React, { useEffect, useState } from 'react';
import '../estilo/Register_cliente.css';

function RegistrationPage() {
  const [clienteEmail, setEmail] = useState({});
  const [clienteUsername, setUsername] = useState({});

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

  useEffect(() => {
    setEmail('');
  
    if (formData.email) {
      fetch(`http://localhost:8070/cliente/email/${formData.email}`)
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
    setUsername('');
  
    if (formData.username) {
      fetch(`http://localhost:8070/cliente/username/${formData.username}`)
        .then(response => response.json())
        .then(data => {
          setUsername(data);
        })
        .catch(error => {
          console.error('Error al obtener los datos del cliente:', error);
        });
    }
  }, [formData.username]);

  const Register = () => {
    if (formData.email === clienteEmail.email) {
      alert('El email ya pertenece a una cuanta');
    }
    else if (formData.username === clienteUsername.username) {
      alert('El username no esta disponible');
    }
    else
    {
      fetch('http://localhost:8070/user', {
      method: 'POST',
      headers: {
      'Content-Type': 'application/json'
      },
      body: JSON.stringify(formData)
      })
      .then(response => response.json())
      .then(data => {
        console.log('Registro exitoso:', data);
        window.location.href = '/login-cliente';
      })
      .catch(error => {
        console.error('Error en el registro:', error);
        alert('Credenciales incorrectas');
      });
    }
  };

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
        <button type="submit">Registrarse</button>
      </form>
    </div>
  );
}

export default RegistrationPage;
