/* eslint-disable */
import * as React from 'react'
import Button from '@mui/material/Button'
import CssBaseline from '@mui/material/CssBaseline'
import TextField from '@mui/material/TextField'
import Link from '@mui/material/Link'
import Grid from '@mui/material/Grid'
import Box from '@mui/material/Box'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import { createTheme, ThemeProvider } from '@mui/material/styles'
import { Alert, Snackbar } from '@mui/material'
import { GoogleLogin } from 'react-google-login'
import { GoogleLoginButton } from 'react-social-login-buttons'
import axios from 'axios'

const theme = createTheme()
const clientId = '78828642227-b3tlfon89t2j66b2a81c60mu8oe45ijb.apps.googleusercontent.com'

export default function Register () {
  const [wrongPassword, setWrongPassword] = React.useState(false)

  const handleSubmit = (event) => {
    event.preventDefault()

    const data = new FormData(event.currentTarget)
    const [firstname, lastname, email, password] = [data.get('firstName'), data.get('lastName'), data.get('email'), data.get('password')]
    axios.post('http://localhost:8080/register/', {
        firstname,
        lastname,
        email,
        password
    }, { headers: {'Content-Type': 'text/plain'} })
    .then(function (response) {
        localStorage.setItem('loggedIn', true)
        location.href = '/wallet'
    })
    .catch(function (error) {
        console.log(error)
    })
  }

  
  const googleResponse = (e) => {
    const [firstname, lastname, email, password] = [e.profileObj.givenName, e.profileObj.familyName, e.profileObj.email, e.profileObj.googleId]
    axios.post('http://localhost:8080/register/', {
        firstname,
        lastname,
        email,
        password
    }, { headers: {'Content-Type': 'text/plain'} })
    .then(function (response) {
    })
    .catch(function (error) {
        console.log(error)
    })
  }

  return (
        <React.Fragment>
            <Snackbar open={wrongPassword} autoHideDuration={6000} onClose={() => setWrongPassword(false)}>
                <Alert onClose={() => setWrongPassword(false)} severity="error" sx={{ width: '100%' }}>
                    Les mots de passe ne vont pas.
                </Alert>
            </Snackbar>
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
                        <Typography component="h1" variant="h5">
                            Inscription
                        </Typography>
                        <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
                            <Grid container spacing={2}>
                                <Grid item xs={12} sm={6}>
                                    <TextField
                                        autoComplete="given-name"
                                        name="firstName"
                                        required
                                        fullWidth
                                        id="firstName"
                                        label="First Name"
                                        autoFocus
                                    />
                                </Grid>
                                <Grid item xs={12} sm={6}>
                                    <TextField
                                        required
                                        fullWidth
                                        id="lastName"
                                        label="Last Name"
                                        name="lastName"
                                        autoComplete="family-name"
                                    />
                                </Grid>
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
                                <Grid item xs={12}>
                                    <TextField
                                        required
                                        fullWidth
                                        name="passwordconf"
                                        label="Password confirmation"
                                        type="password"
                                        id="passwordconf"
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
                                Sign Up
                            </Button>
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
                            <Grid container justifyContent="flex-end">
                                <Grid item>
                                    <Link href="/login" variant="body2">
                                        Already have an account? Sign in
                                    </Link>
                                </Grid>
                            </Grid>
                        </Box>
                    </Box>
                </Container>
            </ThemeProvider>
        </React.Fragment >
  )
}
