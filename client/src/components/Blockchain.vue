<template>
  <div>
    <table class="w3-table w3-striped w3-bordered w3-border">
      <thead>
      <tr>
        <th>#</th>
        <th>Timestamp</th>
        <th>Current Hash</th>
        <th>Prev Hash</th>
        <th>Data</th>
      </tr>
      </thead>
      <tbody>
        <block 
          :blockData="block"
          :key="block.Timestamp"
          v-for="block of blocks">
        </block>
      </tbody>
    </table>
  </div>
</template>

<script>
import Block from './Block'
import store from '../store'

export default {
  components: {
    block: Block
  },
  computed: {
    blocks() {
      return store.state.blocks
    }
  }
}
</script>
