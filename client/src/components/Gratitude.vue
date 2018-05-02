<template>
<div class="modal fade" id="modalGrat" role="dialog">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal">&times;</button>
        <h4><img src = "../assets/wty_white_logo-home.png"></h4>
      </div>
      <div class="modal-body">
        <form role="form">
          <div class="form-group">
            <label for="usr">
              <span class="glyphicon glyphicon-user"></span>
              Receiver's Name
            </label>
            <input
              type="text"
              class="form-control"
              id="usr"
              v-model="recipient"
              placeholder="Enter receiver's email/User ID/Profile Name"
            >
          </div>
          <div class="form-group">
            <label for="msg">
              <span class="glyphicon fa fa-envelope"></span>
              Your Gratitude
            </label>
            <input
              type="text"
              class="form-control"
              id="msg"
              v-model="msgGrat"
              placeholder="Type your message"
            >
          </div>
          <button
            class="btn btn-success btn-block"
            @click="sendGratitude"
            style="background-color:rgb(6, 110, 128)!important"
          >
            <span class="glyphicon glyphicon-off"></span>
            Send Gratitude
          </button>
        </form>
      </div>
      <div class="modal-footer">
        <h4><img src = "../assets/wty_white_logo-home.png"></h4>
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
      sender: '',
      recipient: '',
      msgGrat: ''
    }
  },
  methods: {
    sendGratitude () {
      let dataGrat = {
        recipient: this.recipient,
        msgGrat: this.msgGrat
      }
      console.log('send gratitude called', dataGrat)

      axios.post('/api/users/gratitudes',
        dataGrat,
        {headers: auth.getAuthHeader()})
      .then((response) => {
        console.log('response is', response)
      },
      (error) => {
        console.log('error is', error)
      })
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
</style>


