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
    <gratitude :grat="grat"/>
  </div>
</template>

<script>
import AdminGratitude from './AdminGratitude'

export default {
  data () {
    return {
      grat: {}
    }
  },
  props: ['blockData'],
  components: {
    gratitude: AdminGratitude
  },
  methods: {
    showGrat () {
      this.grat = JSON.parse(atob(this.blockData.data))
      $('#modalGrat').modal()
      console.log('show gratitude', this.grat)
    } 
  }
}
</script>

