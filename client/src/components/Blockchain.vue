<template>
  <div>
    <div class="container-responsive" style="margin: 20px">
      <h2> Filter user</h2>
      <label>user name (sender)</label>
      <input
        v-model="userName"
        placeholder="user name"
      >
      <button
        class="btn btn-primary"
        @click="userGrats">
        show Gratitudes
      </button>
    </div>
    <table class="w3-table w3-striped w3-bordered w3-border">
      <tr>
        <th>#</th>
        <th>Timestamp</th>
        <th>Current Hash</th>
        <th>Prev Hash</th>
        <th>Data</th>
      </tr>
      <tbody
        v-for="block of blocks"
        :key="block.timestamp">
        <block
          :blockData="block"
        >
        </block>
      </tbody>
    </table>
    <gratitude />
  </div>
</template>

<script>
import Block from './Block'
import store from '../store'
import AdminGratitude from './AdminGratitude'

export default {
  data () {
    return {
      userName: ''
    }
  },
  components: {
    block: Block,
    gratitude: AdminGratitude
  },

  computed: {
    blocks() {
      return store.state.blocks
    }
  },
  methods: {
    userGrats () {
      console.log('user gratitudes for user', this.userName)
      let tmpBlocks = []
      for (let block of this.blocks) {
        let grat = JSON.parse(atob(block.data))
        if (grat.sender === this.sender) {
          tmpBlocks.push(block)
          console.log('block pushed', block)
        }
      }
      this.blocks = tmpBlocks
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

