import React, { useContext, useState } from 'react';
import { AuthContext } from './auth';
import '../estilo/login_admin.css';

const AdminLogin = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { loginAdmin } = useContext(AuthContext);

  const handleLoginAdmin = () => {
    fetch(`http://localhost:8090/admin/email/${email}`)
    .then(response => response.json())
    .then(data => {
      if (email === data.email && password === data.password) {
        const token = 'TOKEN_Admin';
        loginAdmin(token, data.id);
        window.location.href = '/';
      } else {
        alert('Credenciales incorrectas');
      }
    })
    .catch(error => {
      console.error('Error al obtener los datos del cliente:', error);
    });
  };

  return (
  <body className= "bodyAdmin"> 
    <div className="container">
      <div className="container2A">
      <h1 className="title">Bienvenido Administrador</h1>
      <div className="form-container">
        <input
          id="inputAD"
          type="text"
          placeholder="Correo electrónico"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          id="inputAD"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="button-container">
          <button className="button" onClick={handleLoginAdmin}>
            Iniciar Sesión
          </button>
        </div>
      </div>
      </div>
    </div>
  </body>
  );
};

export default AdminLogin;
