import React, { useContext, useState } from 'react';
import { AuthContext } from './auth';
import { Link } from 'react-router-dom';
import '../estilo/login_cliente.css';
import Cookies from "universal-cookie";
import { ToastContainer, toast } from "react-toastify";


// probamos del front del sem pasado 
/*async function login(username, password) {
  return await fetch('http://localhost:8070/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({"username": username, "password": password})
  })
      .then(response => {
        if (response.status == 400 || response.status == 401 || response.status == 403)
        {
          return {"user_id": -1, "user_type": "false"}
        }
        return response.json()
      })
      .then(response => {
        Cookie.set("user_id", response.user_id, {path: '/'})
        Cookie.set("username", username, {path: '/login'})
        Cookie.set("user_type", response.type, {path: '/'})
      })
}*/

function goTo(path){
  setTimeout(() => {
      window.location = window.location.origin + path;
  }, 100)
}

const Cookie = new Cookies();

const ClienteLogin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const {loginCliente } = useContext(AuthContext);
  const [userData, setUserData] = useState(null);


  const handleLoginCliente = () => {

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
      if (data.type !== true) {
        // Autenticación exitosa, almacenar datos del usuario y token
        setUserData(data); // Almacena los datos del usuario
        const token = 'TOKEN_CLIENTE';
        loginCliente(token, data.user_id);

        console.log("Data del usuario:", data);
        // Corrige el nombre de la instancia de Cookies a "Cookie" y utiliza "Cookie" en lugar de "Cookies"
        Cookie.set("user_id", data.user_id, { path: '/' });
        Cookie.set("username", username, { path: '/login' });
        Cookie.set("user_type", data.type, { path: '/' });
        goTo('/');
      }else {
        alert("Usted es un administrador. Para iniciar sesion como administrador, dirigase al area de admin.")
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

  /*function logout(){
    Cookie.set("user_id", -1, {path: "/"})
    Cookie.set("user_type", false, {path:"/"})
    document.location.reload()
  }

  const renderGreetings = (
    <body className="bodylogclient">
    <div className="contLogClie1">
    <div className="contLogClien2">
      <h1 className="title">Bienvenido</h1>
      {Cookie.get("username")}
      </div>
    </div>
    <button className='buttonClient' onClick={logout}>Log Out</button>
    </body>
  )*/

  return (
  //Cookie.get("user_id") > -1 ? renderGreetings : 
   <>
    <body className="bodylogclient">
    <div className="contLogClie1">
    <div className="contLogClien2">
      <h1 className="title">Bienvenido Cliente</h1>
      <div className="form-container">
        <input
          type="text"
          placeholder="Nombre de usuario"
          id="inputLcli"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <input
          type="password"
          placeholder="Contraseña"
          id="inputLcli"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <div className="button-container">
          <button className="buttonClient" onClick={handleLoginCliente}>
            Iniciar Sesión
          </button>
          <Link to="/register" className="buttonClient">
            Registrarse
          </Link>
        </div>
      </div>
    </div>
    </div>
    </body> 
    <ToastContainer/>
    </>
  )
  };

export default ClienteLogin;