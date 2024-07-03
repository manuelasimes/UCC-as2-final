import React, { useContext, useState, useEffect } from 'react';
import { AuthContext } from './login/auth';
import './estilo/insert_hoteles.css';
import { ToastContainer, toast } from "react-toastify";
import { useParams } from 'react-router-dom';

const notifyRegistered = () => {
    toast.success("Hotel insertado correctamente!", {
        pauseOnHover: false,
        autoClose: 2000,
    });
};

const notifyError = () => {
    toast.error("Error al insertar hotel!", {
        pauseOnHover: false,
        autoClose: 2000,
    });
};

const handleHomeClick = () => {
    // Lógica para el botón "Home"
    window.location.href = '/';

};

const handleBackClick = () => {
    // Lógica para el botón "Atrás"
    window.history.back();

};

const UpdateHotel = () => {
    const { id } = useParams(); // Get the hotel ID from the URL parameters
    const { auth } = useContext(AuthContext);
    const [formData, setFormData] = useState({
        name: '',
        description: '',
        address: '',
        city: '',
        country: '',
        images: [],
        amenities: []
    });

    const { accessToken } = auth;

    useEffect(() => {
        if (!accessToken) {
            window.location.href = '/login-admin'; // Redirect if no accessToken
            return;
        }

        const fetchHotelData = async () => {
            try {
                const response = await fetch(`http://localhost/hotels-api/hotels/${id}`);
                const data = await response.json();

                setFormData({
                    name: data.name,
                    description: data.description,
                    address: data.address,
                    city: data.city,
                    country: data.country,
                    images: data.images.join(','), // Convert array to comma-separated string
                    amenities: data.amenities.join(','), // Convert array to comma-separated string
                });
            } catch (error) {
                console.error("Error fetching hotel data:", error);
                notifyError();
            }
        };

        fetchHotelData();
    }, [accessToken, id]);

    const handleChange = (event) => {
        event.preventDefault();
        const { name, value } = event.target;

        if (name === 'city' || name === 'country') {

            setFormData((prevFormData) => ({
                ...prevFormData,
                [name]: value.toLowerCase(),
            }));

        }

        setFormData((prevFormData) => ({
            ...prevFormData,
            [name]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log(formData)

        const updatedFormData = {
            ...formData,
            images: formData.images.split(','), // Convert comma-separated string back to array
            amenities: formData.amenities.split(',') // Convert comma-separated string back to array
        };

        try {
            const response = await fetch(`http://localhost/hotels-api/hotels/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    // 'Authorization': `Bearer ${accessToken}`
                },
                body: JSON.stringify(updatedFormData),
            });

            if (!response.ok) {
                notifyError();
            } else {
                notifyRegistered();
                setTimeout(() => {
                    window.location = window.location.origin; // Redirigir a la página principal después del registro
                }, 2000);
            }
        } catch (error) {
            console.error("Error al registrar el hotel:", error);
            notifyError();
        }
    };

    return (
        <div className="registration-container">
            <div className="button-container1">
                <button onClick={handleHomeClick}>Home</button>
                <button onClick={handleBackClick}>Atrás</button>
            </div>
            <h2>Registro De Hoteles</h2>
            <form onSubmit={handleSubmit} className="registration-form">
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
                <br />
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
                <br />
                <button type="submit">Actualizar Hotel</button>
            </form>
            <ToastContainer />
        </div>
    );
};

export default UpdateHotel;
