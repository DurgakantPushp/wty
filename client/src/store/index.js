import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    minors: 4,
    validators: 5,
    blocks: [],
    tmpSndBlocks: [],
    tmpRcptBlocks: [],
    grat: {},
    senders: [],
    recipients: []
  },
  mutations: {
    IncrMinor (state) {
      state.minors += 1
    },
    DecrMinor (state) {
      state.minors -= 1
    },
    InitBlocks (state, sblocks) {
      console.log('blocks received for initialization', sblocks)
      state.blocks = JSON.parse(sblocks)

      for (let block of state.blocks) {
        let grat = JSON.parse(atob(block.data))
        state.senders.push(grat.sender)
        state.recipients.push(grat.recipient)
        console.log('sender pushed', grat.sender)
        console.log('recipient pushed', grat.recipient)
      }
    },
    UpdateBlocks (state, sblock) {
      console.log('blocks received for update', sblock)
      let block = JSON.parse(sblock)
      state.blocks.push(block)

      let grat = JSON.parse(atob(block.data))
      state.senders.push(grat.sender)
      state.recipients.push(grat.recipient)
      console.log('sender pushed', grat.sender)
      console.log('recipient pushed', grat.recipient)
    },
    UpdateGrat (state, grat) {
      console.log('grat updated', grat)
      state.grat = grat
    },
    FilterBySender (state, sender) {
      state.tmpSndBlocks = []

      for (let block of state.blocks) {
        let grat = JSON.parse(atob(block.data))
        if (grat.sender === sender) {
          state.tmpSndBlocks.push(block)
          console.log('block pushed', block)
        }
      }

      swapBlocks(state, 'snd')
    },
    FilterByRecipient (state, recipient) {
      state.tmpRcptBlocks = []

      for (let block of state.blocks) {
        let grat = JSON.parse(atob(block.data))
        if (grat.recipient === recipient) {
          state.tmpRcptBlocks.push(block)
          console.log('block pushed', block)
        }
      }

      swapBlocks(state, 'rcpt')
    },

    FilterReset (state, commEnd) {
      if (commEnd === 'snd') {
        state.blocks = state.tmpSndBlocks
        state.tmpSndBlocks = []
      } else {
        state.blocks = state.tmpRcptBlocks
        state.tmpRcptBlocks = []
      }
    }
  },
  strict: true
})

// swaps tmpBlocks with blocks
function swapBlocks (state, commEnd) {
  let tmp = state.blocks
  if (commEnd === 'snd') {
    state.blocks = state.tmpSndBlocks
    state.tmpSndBlocks = tmp
  } else {
    state.blocks = state.tmpRcptBlocks
    state.tmpRcptBlocks = tmp
  }
}
export default store
