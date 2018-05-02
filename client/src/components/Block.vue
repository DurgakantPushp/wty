<template>
  <div>
    <tr>
      <td>#1</td>
      <td>{{blockData.timestamp}}</td>
      <td>{{blockData.hash}}</td>
      <td>{{blockData.prevHash}}</td>
      <td>
        <button
          type="button"
          class="btn btn-success"
          @click="showGrat">
          {{blockData.data}}
        </button>
      </td>
    </tr>
  </div>
</template>

<script>
import store from '../store'

export default {
  props: ['blockData'],
 
  methods: {
    showGrat () {
      let grat = JSON.parse(atob(this.blockData.data))
      store.commit('UpdateGrat', grat)
      $('#modalGrat').modal()
      console.log('show gratitude', grat)
    } 
  }
}
</script>

