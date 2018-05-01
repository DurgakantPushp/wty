import mqtt from 'mqtt'

import store from '../store'

export default function Init () {
  let optMqtt = {
    clientId: 'js-utility-3P4xG',
    username: 'ajay',
    password: 'akkk'
  }

  let client = mqtt.connect('ws://iot.eclipse.org:80/ws')
  let topicCentreBlocksInit = 'centre/blocks/init'
  let topicCentreBlocksUpdate = 'centre/blocks/update'
  let topicMinorsConnected = 'minors/connected'

  client.on('connect', () => {
    console.log('connected for subscribe', optMqtt)
    client.subscribe(topicMinorsConnected)
    client.subscribe(topicCentreBlocksInit)
    client.subscribe(topicCentreBlocksUpdate)
  })

  client.on('message', (topic, message) => {
    if (topic === 'minors/connected') {
      console.log('minors received', message.toString())
      store.commit('IncrMinor')
    } else if (topic === topicCentreBlocksInit) {
      console.log('blocks received', message.toString())
      store.commit('InitBlocks', message.toString())
    } else if (topic === topicCentreBlocksUpdate) {
      store.commit('UpdateBlocks', message.toString())
    }
  })
}

$(function () {
  Init()
})
