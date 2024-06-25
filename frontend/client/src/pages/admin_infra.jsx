import React, { useContext, useEffect, useState } from 'react';
import { AuthContext } from './login/auth';
import './estilo/hoteles_admin.css';

const AdminInfra = () => {
  const [contenedores, setContenedores] = useState([]);
  const { isLoggedAdmin } = useContext(AuthContext);
  const { isLoggedCliente } = useContext(AuthContext);
  
  function isEmpty(str) {
    return !str.trim().length;
  }

  function isJSONEmpty(obj){
    return Object.keys(obj).length === 0;
  }

  const getContenedores = async () => {
    try {
      // const request = await fetch("http://localhost:8091/cliente/hoteles");
      const request = await fetch("http://localhost:8040/containers");
      //const request = await fetch("http://localhost:8070/hotel");
      const response = await request.json();
      setContenedores(response);
    } catch (error) {
      console.log("No se pudieron obtener los contenedores:", error);
    }
  };

  useEffect(() => {
    getContenedores();
  }, []);

  const Verificacion = () => {
    if (!isLoggedAdmin) {
      window.location.href = '/login-admin';
    }
  };

  const Cuenta = () => {
    if (isLoggedAdmin || isLoggedCliente) {
      window.location.href = '/cuenta';
    }
    else
    {
      window.location.href = '/login-cliente'
    }
  }

  const Home = () => {

    window.location.href = '/'

  }

  const handleVolver = () => {
    window.history.back();
  };

  const handleCrear = (imageName, containerName, containerNumber, runningContainerId) => {
    const newContainerName = `${imageName}-${Number(containerNumber)+1}`;

    fetch(`http://localhost:8040/containers/${imageName}/${newContainerName}/${runningContainerId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        imageName: imageName,
        containerName: newContainerName,
        runningContainerId: runningContainerId
      })
    }).then(response => {
      if (!response.ok) {
        throw new Error('Failed to create container');
      }
      return response.json();
    }).then(data => {
      console.log("Created container ID:", data.containerId);
      // Start the container using the retrieved container ID
      handleStartContainer(data.container_id);
    }).catch(error => {
      console.error("Error creating container:", error);
      alert("Error al crear el contenedor");
    });
    }

  const handleStartContainer = (containerId) => {
    fetch(`http://localhost:8040/containers/start/${containerId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(response => {
      if (!response.ok) {
        throw new Error('Failed to start container');
      }
      console.log("Started container:", containerId);
    }).catch(error => {
      console.error("Error starting container:", error);
      alert("Error al iniciar el contenedor");
    });
  };

  const handleApagar = (contenedorId) => {
    
    console.log(contenedorId);

    fetch(`http://localhost:8040/containers/stop/${contenedorId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(response => {
      if (response.status === 400 || response.status === 401 || response.status === 403 || response.status === 500) {
        alert("Error al parar el contenedor");
      }
    }).catch(error => {
      console.error("Error stopping container:", error);
      alert("Error al parar el contenedor");
    });
  };

  const handlePrender = (contenedorId) => {

    fetch(`http://localhost:8040/containers/start/${contenedorId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(response => {
      if (response.status === 400 || response.status === 401 || response.status === 403 || response.status === 500) {
        alert("Error al iniciar el contenedor");
      }
    }).catch(error => {
      console.error("Error stopping container:", error);
      alert("Error al iniciar el contenedor");
    });
  };

  const handleBorrar = (contenedorId) => {

    fetch(`http://localhost:8040/containers/remove/${contenedorId}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    }).then(response => {
      if (response.status === 400 || response.status === 401 || response.status === 403 || response.status === 500) {
        alert("Error al iniciar el contenedor");
      }
    }).catch(error => {
      console.error("Error stopping container:", error);
      alert("Error al iniciar el contenedor");
    });
  };

  return (
    <body className="bodyinicioH" onLoad={Verificacion}>
     <div className="header-content">
        <div className="admin-button-container">
            <button className="admin-button" onClick={Home}>
                Inicio
            </button>
        </div>
        <div className="cuenta-button-container">
            <button className="cuenta-button" onClick={Cuenta}>
                Tu Cuenta
            </button>
            </div>
        <div className="admin-button-container">
            <button className="admin-button" onClick={handleVolver}>
                Volver
            </button>
        </div>
      </div>
      <div className="containerIniH">
        <div className="hotels-containerH">
          {contenedores.length ? (
            contenedores.map((contenedor) => (
            <div className="hotel-cardH" key={contenedor.Id}>
                <div className='img-name'>
                  <div className="hotel-infoH">
                    <h4> Contenedor: {contenedor.Names} </h4>
                    <div className="hotel-descriptionH">
                    <label htmlFor={`description-${contenedor.Id}`}> Imagen: {contenedor.Image} </label>
                    <label htmlFor={`description-${contenedor.Id}`}> Estado: {contenedor.State}</label>
                    </div>
                    {contenedor.Names && contenedor.Names[0].slice(-1) !== "1" && contenedor.State !== "exited" && (
                      <button className="botonAC" onClick={() => handleApagar(contenedor.Id)}>Apagar</button>
                    )}
                    {contenedor.State === "exited" && (
                      <button className="botonAC" onClick={() => handlePrender(contenedor.Id)}>Prender</button>
                    )}
                    {contenedor.Names && contenedor.Names[0].slice(-1) !== "1" && (
                      <button className="botonAC" onClick={() => handleBorrar(contenedor.Id)}>Borrar</button>
                    )}
                    <button className="botonAC" onClick={() => handleCrear(contenedor.Image, contenedor.Names, contenedor.Labels["com.docker.compose.container-number"], contenedor.Id)}> Crear nuevo contenedor </button>
                  </div>
                </div> 
            </div>
            ))
          ) : (
            <p>No hay contenedores disponibles</p>
          )}
        </div>
      </div>
    </body>
  );
};

export default AdminInfra;