/* eslint-disable */
import axios from 'axios'
import jwt_decode from 'jwt-decode'

const API_URL = '/api/'
const LOGIN_URL = API_URL + '/users/login'
const SIGNUP_URL = API_URL + '/users/signup'

export default {
  Signup(context, creds, redirect) {
    axios.post(SIGNUP_URL, creds).then((response) => {
      localStorage.setItem('id_token', response.body.id_token)
     
      if (redirect) {
        context.$router.replace(redirect)
      }
    }, (error) => {
      console.log('signup error', error)
      context.error = error
    })
  },
  Login(context, creds, redirect) {
    axios.post(LOGIN_URL, creds).then( (response) => {
      localStorage.setItem('id_token', response.data.id_token)

      var decoded = jwt_decode(response.data.id_token)
      console.log('decoded is ', decoded)
      if (redirect) {
        if (decoded.role === "entity") {
          context.$router.replace('/users/entity/home')          
        }
      }
    }, (error) => {
      console.log('login error', error)
      context.error = error
    })
  },
  logout(context) {
    localStorage.removeItem('id_token')
    context.$router.replace('/home')
  },

  isAuthenticated(role) {
    const jwt = localStorage.getItem('id_token')
    if (jwt) {
      let decoded = jwt_decode(jwt)
      if (decoded.role === role) {
        return true
      }
    }
    return false
  },

  getAuthHeader() {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('id_token')
    }
  }
}

