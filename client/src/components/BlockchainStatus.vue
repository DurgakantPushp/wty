<template>
  <div>
    <h1> Blockchain Status</h1>
    <div class="form-group">
      <label for="btnMinor">Minors</label>
      <button
        type="button"
        class="btn btn-circle btn-success btn-xl"
        id="btnMinor"
      >
        {{minors}}
      </button>
    </div>
    <br>
    <div class="form-group">
      <label for="btnValidator">Validators</label>
      <button
        type="button"
        class="btn btn-circle btn-info btn-xl"
        id="btnValidator"
      >
        {{validators}}
      </button>
    </div>
  </div>
</template>

<script>
import store from '../store'
import blockchain from '../blockchain'

export default {
  computed: {
    minors () {
      return store.state.minors
    },
    validators () {
      return store.state.validators
    }
  },
  mounted: function () {
//    blockchain.Init()
    console.log('blockchain web socket init')
  }
}

</script>

<style scoped>
.btn-circle.btn-xl {
  width: 70px;
  height: 70px;
  padding: 10px 16px;
  border-radius: 35px;
  font-size: 24px;
  line-height: 1.33;
}

.btn-circle {
  width: 30px;
  height: 30px;
  padding: 6px 0px;
  border-radius: 15px;
  text-align: center;
  font-size: 12px;
  line-height: 1.42857;
}
</style>
