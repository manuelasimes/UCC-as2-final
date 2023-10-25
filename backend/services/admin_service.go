package services

import (
	adminClient "backend/clients/admin"
	telefonoClient "backend/clients/telefono"
	clienteClient "backend/clients/cliente"
	hotelClient "backend/clients/hotel"
	reservaClient "backend/clients/reserva"

	"backend/dto"
	"backend/model"
	e "backend/utils/errors"
)

type adminService struct{}

type adminServiceInterface interface {
	GetAdminById(id int) (dto.AdminDto, e.ApiError)
	GetAdminByUsername(username string) (dto.AdminDto, e.ApiError)
	GetAdminByEmail(email string) (dto.AdminDto, e.ApiError)
	GetAdmins() (dto.AdminsDto, e.ApiError)
	InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError)
	GetClienteById(id int) (dto.ClienteDto, e.ApiError)
	GetClientes() (dto.ClientesDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotelByEmail(email string) (dto.HotelDto, e.ApiError)
	GetHotelByNombre(nombre string) (dto.HotelDto, e.ApiError)
	GetHoteles() (dto.HotelesDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	AddTelefono(telefonoDto dto.TelefonoDto) (dto.HotelDto, e.ApiError)
	GetReservas() (dto.ReservasDto, e.ApiError)
}

var (
	AdminService adminServiceInterface
)

func init() {
	AdminService = &adminService{}
}

func (s *adminService) GetAdminById(id int) (dto.AdminDto, e.ApiError) {

	var admin model.Admin = adminClient.GetAdminById(id)
	var adminDto dto.AdminDto

	if admin.ID == 0 {
		return adminDto, e.NewBadRequestApiError("Administrador No Encontrado")
	}

	adminDto.ID = admin.ID
	adminDto.Name = admin.Name
	adminDto.LastName = admin.LastName
	adminDto.UserName = admin.UserName
	adminDto.Password = admin.Password
	adminDto.Email = admin.Email

	return adminDto, nil
}

func (s *adminService) GetAdminByUsername(username string) (dto.AdminDto, e.ApiError) {

	var admin model.Admin = adminClient.GetAdminByUsername(username)
	var adminDto dto.AdminDto

	if admin.UserName == "" {
		return adminDto, e.NewBadRequestApiError("Administrador No Encontrado")
	}

	adminDto.ID = admin.ID
	adminDto.Name = admin.Name
	adminDto.LastName = admin.LastName
	adminDto.UserName = admin.UserName
	adminDto.Password = admin.Password
	adminDto.Email = admin.Email

	return adminDto, nil
}

func (s *adminService) GetAdminByEmail(email string) (dto.AdminDto, e.ApiError) {

	var admin model.Admin = adminClient.GetAdminByEmail(email)
	var adminDto dto.AdminDto

	if admin.Email == "" {
		return adminDto, e.NewBadRequestApiError("Administrador No Encontrado")
	}

	adminDto.ID = admin.ID
	adminDto.Name = admin.Name
	adminDto.LastName = admin.LastName
	adminDto.UserName = admin.UserName
	adminDto.Password = admin.Password
	adminDto.Email = admin.Email

	return adminDto, nil
}

func (s *adminService) GetAdmins() (dto.AdminsDto, e.ApiError) {

	var admins model.Admins = adminClient.GetAdmins()
	var adminsDto dto.AdminsDto

	for _, admin := range admins {
		var adminDto dto.AdminDto
		adminDto.ID = admin.ID
		adminDto.Name = admin.Name
		adminDto.LastName = admin.LastName
		adminDto.UserName = admin.UserName
		adminDto.Password = admin.Password
		adminDto.Email = admin.Email

		adminsDto = append(adminsDto, adminDto)
	}

	return adminsDto, nil
}

func (s *adminService) InsertAdmin(adminDto dto.AdminDto) (dto.AdminDto, e.ApiError) {

	var admin model.Admin

	admin.Name = adminDto.Name
	admin.LastName = adminDto.LastName
	admin.UserName = adminDto.UserName
	admin.Password = adminDto.Password
	admin.Email = adminDto.Email

	admin = adminClient.InsertAdmin(admin)

	adminDto.ID = admin.ID

	return adminDto, nil
}

func (s *adminService) GetClienteById(id int) (dto.ClienteDto, e.ApiError) {

	var cliente model.Cliente = clienteClient.GetClienteById(id)
	var clienteDto dto.ClienteDto

	if cliente.ID == 0 {
		return clienteDto, e.NewBadRequestApiError("Cliente No Encontrado")
	}

	clienteDto.ID = cliente.ID
	clienteDto.Name = cliente.Name
	clienteDto.LastName = cliente.LastName
	clienteDto.UserName = cliente.UserName
	clienteDto.Password = cliente.Password
	clienteDto.Email = cliente.Email

	return clienteDto, nil
}

func (s *adminService) GetClientes() (dto.ClientesDto, e.ApiError) {

	var clientes model.Clientes = clienteClient.GetClientes()
	var clientesDto dto.ClientesDto

	for _, cliente := range clientes {
		var clienteDto dto.ClienteDto
		clienteDto.ID = cliente.ID
		clienteDto.Name = cliente.Name
		clienteDto.LastName = cliente.LastName
		clienteDto.UserName = cliente.UserName
		clienteDto.Password = cliente.Password
		clienteDto.Email = cliente.Email

		clientesDto = append(clientesDto, clienteDto)
	}

	return clientesDto, nil
}

func (s *adminService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.ID == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel No Encontrado")
	}

	hotelDto.ID = hotel.ID
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Image = hotel.Image
	hotelDto.Cant_Hab = hotel.Cant_Hab

	for _, telefono := range hotel.Telefonos {
		var dtoTelefono dto.TelefonoDto

		dtoTelefono.Codigo = telefono.Codigo
		dtoTelefono.Numero = telefono.Numero

		hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
	}

	return hotelDto, nil
}

func (s *adminService) GetHotelByEmail(email string) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelByEmail(email)
	var hotelDto dto.HotelDto

	if hotel.Email == "" {
		return hotelDto, e.NewBadRequestApiError("Hotel No Encontrado")
	}

	hotelDto.ID = hotel.ID
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Image = hotel.Image
	hotelDto.Cant_Hab = hotel.Cant_Hab

	for _, telefono := range hotel.Telefonos {
		var dtoTelefono dto.TelefonoDto

		dtoTelefono.Codigo = telefono.Codigo
		dtoTelefono.Numero = telefono.Numero

		hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
	}

	return hotelDto, nil
}

func (s *adminService) GetHotelByNombre(nombre string) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel = hotelClient.GetHotelByNombre(nombre)
	var hotelDto dto.HotelDto

	if hotel.Nombre == "" {
		return hotelDto, e.NewBadRequestApiError("Hotel No Encontrado")
	}

	hotelDto.ID = hotel.ID
	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Image = hotel.Image
	hotelDto.Cant_Hab = hotel.Cant_Hab

	for _, telefono := range hotel.Telefonos {
		var dtoTelefono dto.TelefonoDto

		dtoTelefono.Codigo = telefono.Codigo
		dtoTelefono.Numero = telefono.Numero

		hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
	}

	return hotelDto, nil
}

func (s *adminService) GetHoteles() (dto.HotelesDto, e.ApiError) {

	var hoteles model.Hoteles = hotelClient.GetHoteles()
	var hotelesDto dto.HotelesDto

	for _, hotel := range hoteles {
		var hotelDto dto.HotelDto
		hotelDto.ID = hotel.ID
		hotelDto.Nombre = hotel.Nombre
		hotelDto.Descripcion = hotel.Descripcion
		hotelDto.Email = hotel.Email
		hotelDto.Image = hotel.Image
		hotelDto.Cant_Hab = hotel.Cant_Hab

		for _, telefono := range hotel.Telefonos {
			var dtoTelefono dto.TelefonoDto
	
			dtoTelefono.Codigo = telefono.Codigo
			dtoTelefono.Numero = telefono.Numero
	
			hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
		}

		hotelesDto = append(hotelesDto, hotelDto)
	}

	return hotelesDto, nil
}

func (s *adminService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Nombre = hotelDto.Nombre
	hotel.Descripcion = hotelDto.Descripcion
	hotel.Email = hotelDto.Email
	hotel.Image = hotelDto.Image
	hotel.Cant_Hab = hotelDto.Cant_Hab

	hotel = hotelClient.InsertHotel(hotel)

	hotelDto.ID = hotel.ID

	return hotelDto, nil
}

func (s *adminService) AddTelefono(telefonoDto dto.TelefonoDto) (dto.HotelDto, e.ApiError) {

	var telefono model.Telefono

	telefono.Codigo = telefonoDto.Codigo
	telefono.Numero = telefonoDto.Numero
	telefono.HotelID = telefonoDto.HotelID
	
	telefono = telefonoClient.AddTelefono(telefono)

	
	var hotel model.Hotel = hotelClient.GetHotelById(telefonoDto.HotelID)
	var hotelDto dto.HotelDto

	hotelDto.Nombre = hotel.Nombre
	hotelDto.Descripcion = hotel.Descripcion
	hotelDto.Email = hotel.Email
	hotelDto.Image = hotel.Image
	hotelDto.Cant_Hab = hotel.Cant_Hab

	for _, telefono := range hotel.Telefonos {
		var dtoTelefono dto.TelefonoDto

		dtoTelefono.Codigo = telefono.Codigo
		dtoTelefono.Numero = telefono.Numero

		hotelDto.TelefonosDto = append(hotelDto.TelefonosDto, dtoTelefono)
	}

	return hotelDto, nil
}

func (s *adminService) GetReservas() (dto.ReservasDto, e.ApiError) {

	var reservas model.Reservas = reservaClient.GetReservas()
	var reservasDto dto.ReservasDto

	for _, reserva := range reservas {
		var reservaDto dto.ReservaDto
		reservaDto.ID = reserva.ID
		reservaDto.HotelID = reserva.Hotel.ID
		reservaDto.ClienteID = reserva.Cliente.ID
		reservaDto.AnioFinal = reserva.AnioFinal
		reservaDto.MesInicio = reserva.MesInicio
		reservaDto.MesFinal = reserva.MesFinal
		reservaDto.DiaInicio = reserva.DiaInicio
		reservaDto.DiaFinal = reserva.DiaFinal
		reservaDto.Dias = reserva.Dias

		reservasDto = append(reservasDto, reservaDto)
	}

	return reservasDto, nil
}