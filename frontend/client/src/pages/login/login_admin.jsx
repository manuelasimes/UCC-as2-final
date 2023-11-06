import React, { useContext, useState } from 'react';
import { AuthContext } from './auth';
import '../estilo/login_admin.css';
import Cookies from "universal-cookie";

const Cookie = new Cookies()

function goTo(path){
  setTimeout(() => {
      window.location = window.location.origin + path;
  }, 0,1)
}

const AdminLogin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const {loginAdmin } = useContext(AuthContext);
  const [userData, setUserData] = useState(null);

  const handleLoginAdmin = () => {

    fetch('http://localhost:8070/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: username,
        password: password,
      }),
    })
    .then(response => {
      if (response.status === 400 || response.status === 401 || response.status === 403)
      {
        return {"user_id": -1, "user_type": "false"}
      }
      return response.json()
    }).then(data => {
      if (data.type === true) {
        // Autenticación exitosa, almacenar datos del usuario y token
        setUserData(data); // Almacena los datos del usuario
        const token = 'TOKEN_Admin';
        loginAdmin(token, data.user_id);

        Cookie.set("user_id", data.user_id, { path: '/' });
        Cookie.set("user_type", data.type, { path: '/' });

        console.log("Data del usuario:", data);
        goTo('/');
      }else if (data.type === false){
        alert("No eres un administrador");
        goTo('/login-cliente')
      } else {
        alert("Parece que los datos ingresados son inexistentes. Deberas registrarte como usuario.");
        goTo('/login-cliente')
      }
    })
    /*.then(response => response.json())
    .then(data => {
      if (/*username === data.username && password === data.password*//*data.user_id === ) {
        const token = 'TOKEN_CLIENTE';
        loginCliente(token, data.id);
        window.location.href = '/';
      } else {
        alert('Credenciales incorrectas');
      }
    })*/
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
          value={username}
          onChange={(e) => setUsername(e.target.value)}
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
