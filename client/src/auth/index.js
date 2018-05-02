/* eslint-disable */
import axios from 'axios'
import jwt_decode from 'jwt-decode'

const API_URL = '/api/'
const LOGIN_URL = API_URL + '/users/login'
const SIGNUP_URL = API_URL + '/users/signup'

export default {
  Signup(context, creds) {
    axios.post(SIGNUP_URL, creds).then((response) => {
      localStorage.setItem('id_token', response.data.id_token)

      var decoded = jwt_decode(response.data.id_token)
      console.log('decoded is ', decoded)

      redirect(context, decoded.role)
    }, (error) => {
      console.log('signup error', error)
      context.error = error
    })
  },
  Login(context, creds) {
    axios.post(LOGIN_URL, creds).then( (response) => {
      localStorage.setItem('id_token', response.data.id_token)

      var decoded = jwt_decode(response.data.id_token)
      console.log('decoded is ', decoded)

      redirect(context, decoded.role)
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

// redirect redirects as per role
function redirect(context, role) {
  switch (role) {
    case 'user':
      context.$router.replace('/users/home')
      break
    case 'entity':
      context.$router.replace('/users/entity/home')
      break
    case 'business':
      context.$router.replace('/users/business/home')
      break
    case 'team':
      context.$router.replace('/users/wty-team/home')
      break
    case 'admin':
      context.$router.replace('/users/admin/home')
      break
    default:
      console.log('invalid role received', role)
      break
  }
}