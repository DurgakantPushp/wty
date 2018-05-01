<template>
  <div id="register">
    <div class="modal fade" id="modalReg" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">&times;</button>
            <h4><span class="glyphicon glyphicon-lock"></span> Register</h4>
          </div>
          <div class="modal-body">
            <form role="form">
              <div class="form-group">
                <label>First Name*</label><br>
                <input 
                  type="text"
                  class="form-control"
                  placeholder="Enter First Name"
                  v-model="userInfo.firstName"
                  required
                >
                <br>
                <label>Middle Name</label>
                <br>
                <input 
                  type="text"
                  class="form-control"
                  placeholder="Enter Middle Name"
                  name="middleName"
                  v-model="userInfo.middleName"
                >
                <br>
                <label>Last Name*</label>
                <br>
                <input 
                  type="text"
                  class="form-control"
                  placeholder="Enter Last Name"
                  name="lastName"
                  v-model="userInfo.lastName"
                  required
                >
              </div>
              <br>
              <div class="form-group">
                <label style=" color:Black;">Gender</label>
                <br>
                <input 
                  type="radio"
                  class="inGender"
                  name="gender"
                  value="Male"
                  checked required
                > Male
                <input 
                  class="inGender"
                  type="radio"
                  name="gender"
                  value="Female"
                > Female
                <input 
                  class="inGender"
                  type="radio"
                  value="Other"
                >Other
              </div>
              <br>
              <div class="form-group">
                <label>
                  Select Your Role
                </label>
                <br>
                <input
                  type="radio"
                  class="inRole"
                  name="role"
                  value="User"
                  checked required
                > User 
                <input 
                  type="radio"
                  class="inRole"
                  value="BusinessUser"
                  required
                > Business User
                <br>
                <input 
                  type="radio"
                  class="inRole"
                  value="Entity"
                  required
                >Entity
                <input 
                  type="radio"
                  class="inRole"
                  value="WTYTeam"
                  required
                > WTY Team
                <br>
              </div>
              <br>
              <div class="form-group">
                <label for="email">
                  <span class="glyphicon glyphicon-user"></span>
                  Email id*
                </label>
                <input 
                  type="text"
                  class="form-control"
                  id="email"
                  v-model="userInfo.email"
                  placeholder="Enter your email"
                  required
                >
              </div>
              
              <div class="form-group">
                <label for="username">
                  <span class="glyphicon glyphicon-user"></span>
                  Username*
                </label>
                <input 
                  type="text"
                  class="form-control"
                  id="username"
                  v-model="userInfo.userName"
                  placeholder="Enter your username"
                  required
                >
              </div>

              <div class="form-group">
                <label for="psw">
                  <span class="glyphicon glyphicon-eye-open"></span>
                  Password*
                </label>
                <input
                  type="password" 
                  class="form-control" 
                  id="psw"
                  v-model="userInfo.password"
                  placeholder="Enter password"
                  required
                >
              </div>
              <div class="form-group">
                <label for="pswr">
                  <span class="glyphicon glyphicon-eye-open"></span>
                  Again Password
                </label>
                <input 
                  type="password" 
                  class="form-control" 
                  id="pswr" 
                  placeholder="Confirm password"
                >
              </div>
              <div class="checkbox">
                <label>
                  <input type="checkbox" value="" checked>
                  Remember me
                </label>
              </div>
                <button
                  id="btnReg"
                  class="btn btn-success btn-block" 
                  @click.prevent="submit"
                >
                  <span class="glyphicon glyphicon-off"></span>
                  Register
                </button>
            </form>
          </div>
          <div class="modal-footer">
            <button 
              class="btn btn-danger btn-default pull-left" 
              data-dismiss="modal"
              @click="cancel"
            >
              <span class="glyphicon glyphicon-remove"></span>
              Cancel
            </button>
            <p>Terms and conditions <a href="#">click here</a></p>
            <p> <a href="#">Help me</a></p>
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
        firstName: '',
        middleName: '',
        lastName: '',
        gender: 'male',
        role: 'user',
        email: '',
        userName: '',
        password: ''
      },
      error: ''
    }
  },
  methods: {
    submit() {
      var userInfo = {
        firstName: this.userInfo.firstName,
        middleName: this.userInfo.middleName,
        lastName: this.userInfo.lastName,
        gender: this.userInfo.gender,
        role: this.userInfo.role,
        email: this.userInfo.email,
        userName: this.userInfo.userName,
        password: this.userInfo.password
      }
      auth.Signup(this, userInfo, '/home')
    },
    cancel(){
      console.log("cancelled")
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
  padding: 40px 50px;
}
#btnReg {
  background-color:rgb(6, 110, 128)!important
}
.inRole {
  margin-left:135px;
  color:white
}
.inGender {
  margin-left:135px;
  color:white 
}
</style>


