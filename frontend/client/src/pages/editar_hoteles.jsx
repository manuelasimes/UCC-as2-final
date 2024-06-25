import React, { useState } from "react";
import FormInput from './FormInput'
import './estilo/SignInStyle.css'
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
import { useNavigate } from "react-router-dom";
import {useParams} from "react-router-dom";
import Cookies from "universal-cookie";


const Cookie = new Cookies()

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

const notifyDelete = () => {
    toast.success("Hotel borrado!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}


const EditarHotel = (hotel_id) => {

    const navigate = useNavigate();

    const {id} = useParams()

    // let deleteHotelURL = `http://localhost:8090/hotel/delete/${id}/${Cookie.get("user_id")}`

    const postUser = `http://localhost/hotels-api/hotels/${id}`

    // console.log(deleteHotelURL)

    /* function deleteHotel() {
         fetch(deleteHotelURL, { 
            method: 'DELETE' 
        })
        .then(res => {
            res.json();
            notifyDelete();
        })
        .then(data => console.log(data))
        .catch(error => console.error(error));
    }
 */

    async function updateHotel(jsonData) {

        const response = await fetch(postUser, {
            method: "PUT",
            headers:{"content-type":"application/json"},
            body: JSON.stringify(jsonData),
            mode: "no-cors"
        }).then(response => {
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

    const [values, setValues] = useState({
        name:"",
        description:"",
        rooms:"",
        address:"",
        ImageURL:"",    
    })

    const inputs = [
        {
            id:1,
            name:"name",
            type:"text",
            placeholder:"Nombre",
            label:"Nombre",
        },
        {
            id:2,
            name:"description",
            type:"text",
            placeholder:"Descripcion del hotel",
            label:"Descripcion del hotel",
        },
        {
            id:3,
            name:"country",
            type:"text",
            placeholder:"Pais",
            label:"Pais",
        },
        {
            id:4,
            name:"address",
            type:"text",
            placeholder:"Direccion del hotel",
            label:"Direccion del hotel",
        },
        {
            id:5,
            name:"city",
            type:"text",
            placeholder:"Ciudad",
            label:"Ciudad",
        },
        {
            id:6,
            name:"images",
            type:"text",
            placeholder:"Imagenes",
            label:"Images",
            multiple: true,
            size:"50",
        },
        {
            id:7,
            name:"amenities",
            type:"text",
            placeholder:"Amenidades",
            label:"Amenidades",
            multiple: true,
            size:"50",
        },
    ]

    console.log(id)

     const jsonData = {
        "id": id,
        "name": values.name,
        "description": values.description,
        "address": values.address,
        "images": values.images,
        "amenities": values.amenities,
        "city": values.city,
        "country": values.country
    } 

    // const userObj = { name, last_name, username, password, email, type}

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Dividir el campo 'images' en un array usando una coma como delimitador
        const imagesArray = values.images.split(",");
        const amenitiesArray = values.amenities.split(",");

        // Actualizar el campo 'images' en el objeto 'jsonData' con el array resultante
        jsonData.images = imagesArray;
        jsonData.amenities = amenitiesArray;

        console.log(jsonData)

        updateHotel(jsonData)

    }

  /*   const handleDelete = async (e) => {
        e.preventDefault();

        deleteHotel();
    } */

    const onChange = (e) => {
        setValues({...values, [e.target.name]: e.target.value})
    }

    const handleInsert = () => {
        navigate(`/agregar-hoteles`)
    }

    const handleAmenitie = () => {
        navigate(`/hotel/amenitie/${id}`)
    }

    console.log(values);

    return (
        <>
        <div className="Page">
            <form className = "registration-form" onSubmit={handleSubmit}>
                <h1>Actualizacion!</h1>
                {inputs.map((input) => (
                    <FormInput key={input.id} {...input} value={values[input.name]} onChange={onChange}/>
                ))}
                <button className="RegisterButton">Actualizar!</button>
            </form>
            <div className="AdminHotelOptions">
                {/* <button className="DeleteButton" onClick={handleDelete}>Borrar!</button> */}
                <button className="claseBoton" onClick={handleInsert}>Agregar Hotel!</button>
                {/*<button className="AddHotelButton" onClick={handleAmenitie}>Agregar Amenities!</button>*/}
            </div>
        </div>
        <ToastContainer />
        </>
    )
        
};

export default EditarHotel ;