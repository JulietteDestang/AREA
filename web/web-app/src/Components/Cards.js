import * as React from 'react'
import Card from '@mui/material/Card'
import './style.css'
import { Typography, Grid, Box } from '@mui/material'
import IconButton from '@mui/material/IconButton'
import DeleteIcon from '@mui/icons-material/Delete'
import axios from 'axios'

const topCard = {
    bgcolor: '#5CCCE2',
    borderColor: 'text.primary',
    m: 3,
    padding: '10%',
    borderRadius: '30px'
}
const btmCard = {
    bgcolor: '#3A3A3A',
    borderColor: 'text.primary',
    m: 3,
    padding: '10%',
    borderRadius: '30px'
}

const parentCard = {
    bgcolor: '#262626',
    borderRadius: '20px',
    boxShadow: 3,
    paddingBottom: '10px'
}

export function AREACard ({ cards }) {
    function handleDeletion (card) {
        axios.get('http://localhost:8080/area/delete/' + card.ID)
        .then(function (response) {
            window.location.reload(true)
        }).catch(function (error) {
            console.log(error)
        })
    }
    return (
        <Grid container spacing={4} sx={{ padding: '0 10%', width: '100%', marginLeft: '0px', marginTop: '10px' }}>
            {cards.map((card, index) => (
                < Grid item key={index} xs={12} sm={4} lg={3} style={{ paddingRight: '32px' }}>
                    <Card sx={{ ...parentCard }}>
                    <Card sx={{ ...topCard }}>
                        <Typography gutterBottom variant="h5">
                            {card.action}
                        </Typography>
                    </Card>
                    <Card sx={{ ...btmCard }}>
                        <Typography gutterBottom variant="h5" color={'white'}>
                            {card.reaction}
                        </Typography>
                    </Card>
                    <Box display='flex' justifyContent='center'>
                        <Card sx={{ ...topCard, display: 'flex' }}>
                            <IconButton aria-label="delete" size="large" onClick={ () => { handleDeletion(card) } }>
                                <DeleteIcon fontSize="inherit" />
                            </IconButton>
                        </Card>
                    </Box>
                    </Card>
                </Grid>
            ))}
        </Grid>
    )
}
