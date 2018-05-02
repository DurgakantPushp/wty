<template>
  <div class="modal fade" id="modalGrat" role="dialog">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h2>Gratitude detail</h2>
          <h4>
            <img src = "../assets/wty_white_logo-home.png">
          </h4>
        </div>
        <div class="modal-body">
          <form role="form">
            <div class="form-group">
              <label>
                <span class="glyphicon glyphicon-id"></span>
                id
              </label>
              <button 
                class="btn btn-info"
              >{{grat.id}}
              </button>
            </div>
            <div class="form-group">
              <label>
                <span class="glyphicon glyphicon-user"></span>
                Sender
              </label>
              <button 
                class="btn btn-primary"
              >{{grat.sender}}
              </button>
            </div>
            <div class="form-group">
              <label>
                <span class="glyphicon glyphicon-user"></span>
                Recipient
              </label>
              <button 
                class="btn btn-info"
              >{{grat.recipient}}
              </button>
            </div>
             <div class="form-group">
              <label>
                <span class="glyphicon glyphicon-cnt"></span>
                Content
              </label>
              <button 
                class="btn btn-success"
              >{{grat.cnt}}
              </button>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button type="submit" 
          class="btn btn-danger btn-default pull-left" 
          data-dismiss="modal">
            <span class="glyphicon glyphicon-remove"></span>
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import store from '../store'

export default {
  computed: {
    grat () {
      return store.state.grat
    }
  }
}
</script>

