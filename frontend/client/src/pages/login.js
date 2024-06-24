import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import {
  Grid,
  Paper,
  TextField,
  Typography,
  Button,
  FormHelperText,
  Snackbar,
  Alert,
  FormControl,
  IconButton,
  Input,
  InputLabel,
  InputAdornment,
  MenuItem,
  CircularProgress,
  Link,
} from '@mui/material';
import { Visibility, VisibilityOff } from '@mui/icons-material';
import ResumeDataModal from '@/components/ResumeDataModal';
import AuthService from '@/services/auth.service';
import AttentionCentersService from '@/services/attentionCenters.service';
import { setUserData } from '../stores/slices/user';
import { setCenterConfig } from '../stores/slices/centerConfig';
import logoFCM from '../assets/logos_fcm.png';
import logoSiHosp from '../assets/logos_sihosp.png';
import logoUNC from '../assets/logos_unc.png';
import {Formik, useFormik } from "formik"

export default function Login() {
  const [attentionCenters, setAttentionCenters] = useState([]);
  const [errorUser, setErrorUser] = useState(false);
  const [username, setUser] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [center, setCenter] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [showCenterModal, setShowCenterModal] = useState(false);
  const [showRecoverModal, setShowRecoverModal] = useState(false);
  /*const formik = useFormik ({
    initialValues {
      username:"",
      email :"",
      password: ""
    }
  })*/
  const [showSnackbar, setShowSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });
  const [isLoading, setIsLoading] = useState(false);
  const dispatch = useDispatch(); 
  const passRef = useRef();
  const navigate = useNavigate();

  const getConfig = async () => {
    try {
      const { data } = await AuthService.login(username, password, center);
      const res = await AttentionCentersService.getCenterConfig(center);
      window.localStorage.setItem('access_token', data.access_token);
      window.localStorage.setItem('refresh_token', data.refresh_token);
      window.localStorage.setItem('user', JSON.stringify(data.user));
      window.localStorage.setItem('center', center);
      dispatch(setUserData(data.user));
      dispatch(setCenterConfig(res.data));
      setShowCenterModal(false);
      navigate('/inicio');
    } catch (error) {
      null;
    }
  };

  const handleCloseCenterModal = () => {
    setShowCenterModal(false);
  };

  const onChangePassword = () => {
    setShowPassword((prevState) => !prevState);
  };

  useEffect(() => {
    const fetchData = async () => {
      const { data } = await AuthService.login(username, password);
      if (data && data.user.institutions.length > 0) {
        setCenter(data.user.institutions[0].id);
        setAttentionCenters(data.user.institutions);
      }
    };

    if (username && password) {
      fetchData();
    }
  }, [username, password]);

  const handleLogin = async () => {
    if (username && password) {
      setIsLoading(true);
      try {
        const { data } = await AuthService.login(username, password);
        if (data && data.user.institutions.length > 1) {
          setShowCenterModal(true);
          setAttentionCenters(data.user.institutions);
        } else {
          getConfig();
        }
      } catch (error) {
        setErrorUser(error);
      }
      setIsLoading(false);
    }
  };

  const validateEmail = () => {
    return String(email)
      .toLowerCase()
      .match(
        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
      );
  };

  const handleRecoverPassword = () => {
    if (email && validateEmail()) {
      AuthService.sendRecoverEmail({ email, attentionCenterId: 1 })
        .then(() =>
          setShowSnackbar({
            open: true,
            message:
              'Se envió un mail a su casilla de correo con las instrucciones para recuperar su contraseña',
            severity: 'success',
          }),
        )
        .catch(() =>
          setShowSnackbar({
            open: true,
            message: 'Ocurrió un error enviando el mail para recuperar su contraseña',
            severity: 'error',
          }),
        )
        .finally(handleCloseModal);
    } else {
      setShowSnackbar({
        open: true,
        message: 'Ingrese un mail válido para recuperar su contraseña',
        severity: 'error',
      });
      handleCloseModal();
    }
  };

  const handleCloseSnackbar = () => {
    setShowSnackbar({
      open: false,
      message: '',
      severity: 'success',
    });
  };

  const handleCloseModal = () => {
    setShowRecoverModal(false);
    setEmail('');
  };

  return (
    <Grid
      container
      display="flex"
      justifyContent="center"
      alignItems="center"
      width="100vw"
      height="100vh"
      data-testid="login-test">
      <Paper elevation={2}>
        <Grid
          container
          item
          xs={12}
          direction="row"
          minWidth="780px"
          minHeight="400px"
          alignItems="center">
          <Grid xs={6} item justifyContent="center" alignItems="center" paddingX="2rem">
            <Typography
              variant="p"
              component="h3"
              marginBottom="20px"
              paddingTop="20px"
              color="gray"
              fontWeight="normal">
              Iniciar Sesión
            </Typography>
            <TextField
              helperText={!username && 'Por favor ingrese usuario'}
              error={errorUser && true}
              id="username"
              label="Usuario"
              variant="standard"
              value={username}
              margin="normal"
              size="normal"
              onKeyDown={(e) => {
                if (e.code === 'Enter') {
                  password ? handleLogin() : passRef.current.focus();
                }
              }}
              onChange={(e) => setUser(e.target.value)}
              sx={{ marginBottom: '20px', width: '100%' }}
            />
            <FormControl
              sx={{ width: '100%', marginBottom: '50px', marginTop: '20px' }}
              variant="standard"
              error={errorUser && true}>
              <InputLabel htmlFor="password">Contraseña</InputLabel>
              <Input
                id="password"
                inputRef={passRef}
                type={showPassword ? 'text' : 'password'}
                value={password}
                onKeyDown={(e) => e.code === 'Enter' && handleLogin()}
                onChange={(e) => setPassword(e.target.value)}
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton
                      aria-label="toggle password visibility"
                      onClick={onChangePassword}
                      onMouseDown={onChangePassword}>
                      {showPassword ? <VisibilityOff /> : <Visibility />}
                    </IconButton>
                  </InputAdornment>
                }
              />
              {!password && <FormHelperText>Por favor ingrese contraseña</FormHelperText>}
            </FormControl>
            {errorUser && (
              <Typography color="red" marginBottom="20px">
                Usuario o contraseña incorrecta
              </Typography>
            )}
            <Grid
              xs={12}
              container
              item
              direction="row"
              justifyContent="space-between"
              alignItems="center">
              <Button
                variant="text"
                sx={{ textTransform: 'none' }}
                onClick={() => setShowRecoverModal(true)}>
                Recuperar contraseña
              </Button>
              <Button
                variant="contained"
                onClick={handleLogin}
                sx={{ textTransform: 'none', paddingX: '30px', minWidth: 120, minHeight: 35 }}>
                {isLoading ? <CircularProgress color={'inherit'} size={20} /> : 'Ingresar'}
              </Button>
            </Grid>
          </Grid>
          <Grid xs={6} item textAlign="center">
            <Grid paddingTop="60px">
              <img src={logoSiHosp} alt="logo" width={'265px'} />
            </Grid>
            <Grid
              display="flex"
              alignItems="center"
              justifyContent="space-around"
              paddingX="50px"
              marginTop="50px">
              <Link
                href="https://fcm.unc.edu.ar/"
                target="_blank"
                onMouseOver={(e) => (e.target.style.transform = 'scale(1.1)')}
                onMouseOut={(e) => (e.target.style.transform = 'scale(1)')}>
                <img src={logoFCM} alt="logo" width={'80px'} />
              </Link>
              <Link
                href="https://www.unc.edu.ar/"
                target="_blank"
                onMouseOver={(e) => (e.target.style.transform = 'scale(1.1)')}
                onMouseOut={(e) => (e.target.style.transform = 'scale(1)')}>
                <img src={logoUNC} alt="logo" width={'80px'} />
              </Link>
            </Grid>
          </Grid>
        </Grid>
      </Paper>
      <ResumeDataModal
        title="Ingreso"
        open={showCenterModal}
        onClose={handleCloseCenterModal}
        onConfirm={getConfig}
        disabledConfirmButton={center === ''}>
        <TextField
          label="Elegir un centro"
          variant="filled"
          value={center}
          onChange={(e) => setCenter(e.target.value)}
          sx={{ marginBottom: '20px', width: '100%' }}
          select>
          {attentionCenters.map((data) => (
            <MenuItem key={data.id} value={data.id}>
              {data.name}
            </MenuItem>
          ))}
        </TextField>
      </ResumeDataModal>
      <ResumeDataModal
        title="INGRESE EL EMAIL PARA RECUPERAR SU CONTRASEÑA"
        open={showRecoverModal}
        onClose={handleCloseModal}
        onConfirm={handleRecoverPassword}>
        <TextField
          label="email"
          variant="filled"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          sx={{ marginBottom: '20px', width: '100%' }}
        />
      </ResumeDataModal>
      <Snackbar
        open={showSnackbar.open}
        anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
        autoHideDuration={5000}
        onClose={handleCloseSnackbar}>
        <Alert
          onClose={handleCloseSnackbar}
          severity={showSnackbar.severity}
          sx={{ width: '100%' }}>
          {showSnackbar.message}
        </Alert>
      </Snackbar>
    </Grid>
  );
}