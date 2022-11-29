/** @file App.js
 * @brief main file
 *
 * more detailed version here
 */

import React from 'react'
import Home from './Components/Home'
import Login from './Components/Login'
import Wallet from './Components/Wallet'
import NavBar from './Components/Navbar'
import { Route, Routes } from 'react-router-dom'
import Register from './Components/Register'
import Services from './Components/Services'
import Cookies from 'js-cookie'
import Forbidden from './Components/Forbidden'

export default function App () {
  const [loggedIn, setLoggedIn] = React.useState(Cookies.get('jwt') < 0)

  return (
    <div className='App'>
      <NavBar setLoggedIn={ setLoggedIn } loggedIn={ loggedIn }/>
      <Routes>
        <Route path='/' element={<Home />} />
        {loggedIn === true &&
        <React.Fragment>
          <Route path='/wallet' element={<Wallet />} />
          <Route path='/user/services' element={<Services />} />
        </React.Fragment>
        }
        {loggedIn === false &&
        <React.Fragment>
          <Route path='/user/services' element={<Forbidden />} />
          <Route path='/wallet' element={<Forbidden />} />
          <Route path='/register' element={<Register />} />
          <Route path='/login' element={<Login />} />
        </React.Fragment>
      }
      </Routes>
    </div>
  )
}
