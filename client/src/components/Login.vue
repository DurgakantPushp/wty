<template>
  <div>
    <div class="modal fade" id="loginModal" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button"
              class="close"
              data-dismiss="modal"
            >&times;</button>
            <h4>
              <img src = "../assets/wty_white_logo-home.png">
            </h4>
          </div>
          <div class="modal-body">
            <form role="form">
              <div class="form-group">
                <label for="userName">
                  <span class="glyphicon glyphicon-user"></span>
                  Username
                </label>
                <input
                  type="text"
                  class="form-control"
                  id="userName"
                  v-model="userInfo.userName"
                  placeholder="Enter username"
                >
              </div>
              <div class="form-group">
                <label for="psw1">
                  <span class="glyphicon glyphicon-eye-open"></span>
                  Password
                </label>
                <input
                  type="password"
                  class="form-control"
                  id="psw1"
                  v-model="userInfo.password"
                  placeholder="Enter password"
                >
              </div>
              <div class="checkbox">
                <label>
                  <input
                    type="checkbox"
                    value=""
                    checked
                  >
                    Remember me
                  </label>
              </div>
              <button
                class="btn btn-success btn-block"
                id="btnLogin"
                @click.prevent="login"
              >
                <span class="glyphicon glyphicon-off"></span>
                Login
              </button>
            </form>
          </div>
          <div class="modal-footer">
            <button type="submit"
            class="btn btn-danger btn-default pull-left"
            data-dismiss="modal">
              <span class="glyphicon glyphicon-remove"></span>
              Cancel
            </button>
            <p>Not a member? <a href="#">Sign Up</a></p>
            <p>Forgot <a href="#">Password?</a></p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  data () {
    return {
      userInfo: {
        userName: '',
        password: ''
      }
    }
  },
  methods: {
    login () {
      let userInfo = {
        userName: this.userInfo.userName,
        password: this.userInfo.password
      }
      auth.Login(this, userInfo)
    }
  }
}
</script>

<style scoped>
.modal-header {
  padding: 35px 50px;
  background-color:rgb(6, 110, 128)!important
}
.modal-body {
  padding: 40px 50px
}
#btnLogin {
  background-color:rgb(6, 110, 128)!important
}
</style>
