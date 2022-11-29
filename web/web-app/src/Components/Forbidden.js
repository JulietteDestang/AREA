import * as React from 'react'
import { Grid, Card, Box, Link, Button } from '@mui/material'
import { createTheme, ThemeProvider, Typography } from '@material-ui/core'

const btmCard = {
    bgcolor: '#5CCCE2',
    borderColor: 'text.primary',
    m: 3,
    padding: '10%',
    borderRadius: '30px'
}
const parentCard = {
    bgcolor: '#3A3A3A',
    borderRadius: '30px',
    boxShadow: 3,
    padding: '0%'
}

const theme = createTheme({
    typography: {
      fontFamily: ['Titan One', 'cursive'].join(',')
    }
  })

export default function Forbidden () {
    return (
        <Box sx={{
            marginTop: 5,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center'
        }}>
            <ThemeProvider theme={theme}>
                <Typography variant='h2' gutterBottom> Introuvable...</Typography>
            </ThemeProvider>
            <Card sx={{ ...parentCard }}>
                <Card sx={{ ...btmCard }}>
                    <Typography gutterBottom variant="h5">
                        Vous n&apos;avez rien à faire ici. Sortez, madame.
                    </Typography>
                </Card>
                {/* <Button> */}
                    <Link href={'/login'}>
                        <Card sx={{ ...btmCard }}>
                            <Typography gutterBottom variant="h5">
                                Connexion
                            </Typography>
                        </Card>
                    </Link>
                    <Link href={'/register'}>
                        <Card sx={{ ...btmCard }}>
                            <Typography gutterBottom variant="h5">
                                Inscription
                            </Typography>
                        </Card>
                    </Link>

                {/* </Button> */}
            </Card>
        </Box>
    )
}
