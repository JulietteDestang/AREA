/*eslint-disable*/
import { Box, Button, Card, CardContent, CardMedia, Grid, ButtonBase, Dialog } from '@mui/material'
import * as React from 'react'
import axios from 'axios'
import githubImg from '../resources/github.png'
import spotifyImg from '../resources/spotify.png'
import deezerImg from '../resources/deezer.png'
import discordImg from '../resources/discord.svg'
import gmailImg from '../resources/gmail.png'
import { DialogTitle, TextField } from '@material-ui/core'
import { createTheme, ThemeProvider, Typography } from '@material-ui/core'

export default function Services () {

    const theme = createTheme({
        typography: {
          fontFamily: ['Titan One', 'cursive'].join(',')
        }
    })
    const services = [{ name: 'Spotify', token: null, img: spotifyImg }, { name: 'Deezer', token: null, img: deezerImg }, { name: 'Github', token: null, img: githubImg }, { name: 'Discord', token: null, img: discordImg }, { name: 'Gmail', token: null, img: gmailImg }]
    return (
        <React.Fragment>
            <Box sx={{
                marginTop: 8,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center'
            }}>
            <ThemeProvider theme={theme}>
                <Typography variant='h2' gutterBottom>Services</Typography>
            </ThemeProvider>
            </Box>
            <ServicesCard services={ services }/>
        </React.Fragment >
    )
}

function ServicesCard ({ services }) {
    const [service, setService] = React.useState(null)
    const [serviceToken, setServiceToken] = React.useState(null)
    const [showInputs, setShowInputs] = React.useState(false)
    const [validemail, setValidEmail] = React.useState(false)
    const [password, setPassword] = React.useState("")
    const [email, setEmail] = React.useState("")
    const [emailError, setEmailError] = React.useState(true)
    const [passwordError, setPasswordError] = React.useState(true)
    const [isButtonDisabled, setIsButtonDisabled] = React.useState(true)

    const getDiscordToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/discord/auth/url', { headers })
        .then(function (response) {
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }

    const getSpotifyToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/spotify/auth/url', { headers })
        .then(function (response) {
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }

    const getDeezerToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/deezer/auth/url', { headers })
        .then(function (response) {
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }

    const getGithubToken = () => {
        const headers = { 'Content-Type': 'text/plain' }
        axios.get('http://localhost:8080/github/auth/url', { headers })
        .then(function (response) {
            location.href = response.data
        }).catch(function (error) {
            console.log(error)
        })
    }

    const submitGmail = () => {
        axios.post('http://localhost:8080/email/login', {
            email: email,
            password: password,
        }, { headers: { 'Content-Type': 'text/plain' }, withCredentials: true })
        .then(function (response) {
            // location.href = response.data
        }).catch(function (error) {
                console.log(error)
        })
        setShowInputs(false)
    }

    const getGmailToken = () => {
        setShowInputs(true)
    }

    const handleClick = (service) => {
        setService(service.name)
        switch (service.name) {
            case 'Spotify':
                getSpotifyToken()
                break
            case 'Deezer':
                getDeezerToken()
                break
            case 'Discord':
                getDiscordToken()
                break
            case 'Github':
                getGithubToken()
                break
            case 'Gmail':
                getGmailToken()
                break
            default:
                break
        }
    }
       
    function handleError () {
        if (!emailError && !passwordError) {
            setIsButtonDisabled(false)
        } else {
            setIsButtonDisabled(true)
        }
    }

    const handleEmailChange = e => {
        setEmail(e.currentTarget.value)
        if (/\S+@\S+\.\S+/.test(e.currentTarget.value)) {
            setEmailError(false)
        } else {
            setEmailError(true)
        }
        handleError()
    }

    const handlePasswordChange = e => {
        setPassword(e.currentTarget.value)
        if (e.currentTarget.value.length <= 4) {
            setPasswordError(true)
        } else {
            setPasswordError(false)
        }
        handleError()
    }

    return (
        <Grid container spacing={4} sx={{ padding: '0 10%', width: '100%', marginLeft: '0px' }}>
            {services.map((service, index) => (
                < Grid item key={index} xs={12} sm={6} md={4} style={{ paddingRight: '32px' }}>
                    <ButtonBase onClick={e => handleClick(service)} style={{ width: '100%' }}>
                    <Card
                        sx={{ display: 'flex', flexDirection: 'column' }}
                    >
                        <CardContent
                        sx={{ flexGrow: 1 }}
                        style={ service.token === null ? { backgroundColor: '#d5dbe6' } : { backgroundColor: 'white' } }
                        >
                            <Typography gutterBottom variant="h5" component="h2" align='center'>
                                {service.name}
                            </Typography>
                        </CardContent>
                        <CardMedia
                            component="img"
                            image={ service.img }
                            alt="serviceImg"
                        />
                    </Card>
                        </ButtonBase>
                        
                </Grid>
            ))
            }
            {showInputs &&
            <Dialog open={showInputs} >
                <DialogTitle>Log in with Gmail</DialogTitle>
                <Box component="form" noValidate sx={{ mt: 3 }}>
                    <Grid container spacing={2}>
                        <Grid item xs={12}>
                            <TextField
                                required
                                fullWidth
                                id="email"
                                label="Email Address"
                                name="email"
                                onChange={ handleEmailChange }
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
                                onChange={ handlePasswordChange }
                                autoComplete="new-password"
                            />
                        </Grid>
                    </Grid>
                        <Button
                            onClick={ submitGmail }
                            disabled={isButtonDisabled}
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign In
                        </Button>
                </Box>
            </Dialog>}
        </Grid >
    )
}
