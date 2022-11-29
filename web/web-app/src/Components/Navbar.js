import * as React from 'react'
import { styled } from '@mui/material/styles'
import AREALogo from './Icons/AREALogo'
import AccountIcon from './Icons/AccountIcon'
import { AppBar, Link, Button } from '@mui/material'
import { Box } from '@mui/system'
import axios from 'axios'
import LogoutIcon from './Icons/LogoutIcon'
import LogInIcon from './Icons/LogInIcon'

const StyledAppBar = styled(AppBar)(({ theme }) => ({
    backgroundColor: '#262626'
}))
export default function NavBar ({ setLoggedIn, loggedIn }) {
    React.useEffect(() => {
        const cookie = document.cookie.indexOf('jwt')
        if (cookie !== -1) {
            setLoggedIn(true)
        }
    }, [])

    const handleLogout = (event) => {
        event.preventDefault()
        axios.get('http://localhost:8080/logout/', { withCredentials: true })
        .then(function () {
            localStorage.setItem('loggedIn', false)
            location.href = '/'
        }).catch(function (error) {
            console.log(error)
        })
    }

    return (
        <StyledAppBar position='sticky'>
            <Box style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <Box >
                    {loggedIn
                    ? <Link href={'/wallet'}>
                        <AccountIcon/>
                    </Link>
                    : <Link href={'/login'}>
                        <LogInIcon />
                    </Link>
                    }
                </Box>
                <Box >
                    <Link href="/">
                        <AREALogo/>
                    </Link>
                </Box>
                <Box >
                    <Button onClick={handleLogout}>
                        <LogoutIcon/>
                    </Button>
                </Box>
            </Box>
        </StyledAppBar>
    )
}
