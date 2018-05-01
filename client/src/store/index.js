import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    minors: 4,
    validators: 5,
    blocks: []
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
    },
    UpdateBlocks (state, sblock) {
      console.log('blocks received for update', sblock)
      let block = JSON.parse(sblock)
      state.blocks.push(block)
    }
  },
  strict: true
})

export default store
