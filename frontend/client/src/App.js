//app.js
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import LogCliente from './pages/login/login_cliente';
import LogAdmin from './pages/login/login_admin';
import Register from './pages/login/Register_cliente';
import Inicio from './pages/inicio';
import InicioAdmin from './pages/inicio_admin';
import AdminHoteles from './pages/admin_hoteles';
import AdminClientes from './pages/admin_clientes';
import Reservar from './pages/reservar';
import Cuenta from './pages/cuenta';
import ReservasCliente from './pages/reservas_cliente';
import InsertHoteles from './pages/insert_hoteles';
import VerHoteles from './pages/ver_hoteles';
import VerReservas from './pages/ver_reservas';
import VerClientes from './pages/ver_clientes';
import EditarHoteles from './pages/editar_hoteles';
import AdminInfra from './pages/admin_infra';
import { AuthProvider } from './pages/login/auth';
import ProtectedRoute from './ProtectedRoute';
import './App.css';

function App() {
    return (
        <div>

            <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css" />
            <Router>
                <AuthProvider>
                    <Routes>
                        <Route path='/' element={<Inicio />}></Route>
                        <Route path='/admin' element={
                            <ProtectedRoute adminOnly={true}>
                                <InicioAdmin />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/login-cliente' element={<LogCliente />}></Route>
                        <Route path='/login-admin' element={<LogAdmin />}></Route>
                        <Route path='/register' element={<Register />}></Route>
                        <Route path='/administrar-hoteles' element={
                            <ProtectedRoute adminOnly={true}>
                                <AdminHoteles />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/administrar-clientes' element={
                            <ProtectedRoute adminOnly={true}>
                                <AdminClientes />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/reservar/:hotelId' element={
                            <ProtectedRoute>
                                <Reservar />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/cuenta' element={
                            <ProtectedRoute>
                                <Cuenta />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/reservas-cliente' element={
                            <ProtectedRoute>
                                <ReservasCliente />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/agregar-hoteles' element={
                            <ProtectedRoute adminOnly={true}>
                                <InsertHoteles />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/ver-hoteles' element={
                            <ProtectedRoute adminOnly={true}>
                                <VerHoteles />
                            </ProtectedRoute>}></Route>
                        <Route path='/ver-reservas' element={
                            <ProtectedRoute adminOnly={true}>
                                <VerReservas />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/ver-clientes' element={
                            <ProtectedRoute adminOnly={true}>
                                <VerClientes />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/hotel/edit/:id' element={
                            <ProtectedRoute adminOnly={true}>
                                <EditarHoteles />
                            </ProtectedRoute>
                        }></Route>
                        <Route path='/infraestructura' element={
                            <ProtectedRoute adminOnly={true}>
                                <AdminInfra />
                            </ProtectedRoute>
                        }></Route>
                    </Routes>
                </AuthProvider>
            </Router>
        </div>
    )
}
export default App;