<template>
  <div>
    <div class="container-responsive" style="margin: 20px">
      <h2> Filter user by sender</h2>
      <label>sender</label>
      <select
        v-model="sender"
      >
        <option disabled value="">select sender</option>
        <option
          v-for="sender of senders"
          :key="sender.id"
        >
          {{sender}}
        </option>
      </select>
      <button
        class="btn btn-primary"
        v-if="sndEnabled"
        @click="filterBySender">
        Gratitude as per sender
      </button>
      <button
        class="btn btn-danger"
        v-else
        @click="filterResetBySender">
        Reset
      </button>
    </div>
    <div class="container-responsive" style="margin: 20px">
      <h2> Filter user by recipient</h2>
      <label>recipient</label>
      <select
        v-model="recipient"
      >
        <option disabled value="">select recipient</option>
        <option
          v-for="recipient of recipients"
          :key="recipient.id">
          {{recipient}}
        </option>
      </select>
      <button
        class="btn btn-primary"
        v-if="rcptEnabled"
        @click="filterByRecipient">
        Gratitude as per recipient
      </button>
      <button
        class="btn btn-danger"
        v-else
        @click="filterResetByRecipient">
        Reset
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
      sender: '',
      recipient: '',
      sndEnabled: true,
      rcptEnabled: true
    }
  },
  components: {
    block: Block,
    gratitude: AdminGratitude
  },

  computed: {
    blocks () {
      return store.state.blocks
    },
    senders () {
      return store.state.senders
    },
    recipients () {
      return store.state.recipients
    }
  },

  methods: {
    filterBySender () {
      console.log('user gratitudes for sender', this.sender)
      store.commit('FilterBySender', this.sender)
      this.sndEnabled = false
    },
    filterByRecipient () {
      console.log('gratitudes received by recipient', this.recipient)
      store.commit('FilterByRecipient', this.recipient)
      this.rcptEnabled = false
    },
    filterResetBySender () {
      store.commit('FilterReset', 'snd')
      this.sndEnabled = true
    },
    filterResetByRecipient () {
      store.commit('FilterReset', 'rcpt')
      this.rcptEnabled = true
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

