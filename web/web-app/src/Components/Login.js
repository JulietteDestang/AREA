/*eslint-disable*/
import * as React from 'react'
import Avatar from '@mui/material/Avatar'
import Button from '@mui/material/Button'
import CssBaseline from '@mui/material/CssBaseline'
import TextField from '@mui/material/TextField'
import Link from '@mui/material/Link'
import Grid from '@mui/material/Grid'
import Box from '@mui/material/Box'
import LockOutlinedIcon from '@mui/icons-material/LockOutlined'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import { GoogleLogin } from 'react-google-login'
import { GoogleLoginButton } from 'react-social-login-buttons'
import { gapi } from 'gapi-script'
import axios from 'axios'

const theme = createTheme()
const clientId = '78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com'

export default function SignIn () {

  const handleSubmit = (event) => {
    event.preventDefault()
    const data = new FormData(event.currentTarget)

    const [email, password] = [data.get('email'), data.get('password')]
    axios.post('http://localhost:8080/login/', {
      email,
      password,
    }, {headers: {'Content-Type': 'text/plain'}, withCredentials: true} ) 
    .then(function (response) {
      localStorage.setItem('loggedIn', true)
      location.href = '/wallet'

    })
    .catch(function (error) {
      console.log(error)
    })
  }

  const googleResponse = (e) => {
    const [email, password] = [e.profileObj.email, e.profileObj.googleId]
    axios.post('http://localhost:8080/login/', {
      email,
      password,
    }, {headers: {'Content-Type': 'text/plain'}, withCredentials: true} ) 
    .then(function (response) {
      localStorage.setItem('loggedIn', true)
      location.href = '/wallet'

    })
    .catch(function (error) {
      console.log(error)
    })
  }

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center'
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        required
                        fullWidth
                        id="email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        required
                        fullWidth
                        name="password"
                        label="Password"
                        type="password"
                        id="password"
                        autoComplete="new-password"
                    />
                </Grid>
              </Grid>
            <Button
                type="submit"
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
            >
                Sign In
            </Button>
          </Box>
          {/* <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}> */}
          <GoogleLogin
            clientId={clientId}
            render={renderProps => (
              <GoogleLoginButton onClick={renderProps.onClick} disabled={renderProps.disabled} />
            )}
            buttonText="Login"
            onSuccess={googleResponse}
            onFailure={googleResponse}
            cookiePolicy={'single_host_origin'}
          />
          <Grid container>
            <Grid item>
              <Link href="/register" variant="body2">
                {"Don't have an account? Sign Up"}
              </Link>
            </Grid>
          </Grid>
          {/* </Box> */}
        </Box>
      </Container>
    </ThemeProvider>
  )
}
