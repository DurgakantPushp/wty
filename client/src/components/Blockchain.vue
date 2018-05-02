<template>
  <div>
    <table class="w3-table w3-striped w3-bordered w3-border">
      <tr>
        <th>#</th>
        <th>Timestamp</th>
        <th>Current Hash</th>
        <th>Prev Hash</th>
        <th>Data</th>
      </tr>
        <block 
          :blockData="block"
          :key="block.timestamp"
          v-for="block of blocks">
        </block>
        <gratitude />
    </table>
  </div>
</template>

<script>
import Block from './Block'
import store from '../store'
import AdminGratitude from './AdminGratitude'

export default {
  components: {
    block: Block,
    gratitude: AdminGratitude
  },
  
  computed: {
    blocks() {
      return store.state.blocks
    }
  }
}
</script>

<style scoped>
table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    width: 100%;
}

td, th {
    border: 1px solid #dddddd;
    text-align: left;
    padding: 8px;
}

tr:nth-child(even) {
    background-color: #dddddd;
}
</style>

