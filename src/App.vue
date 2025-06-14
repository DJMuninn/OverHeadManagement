<script setup lang="ts">

import Message from './components/Message.vue'
import Counter from './components/Counter.vue'
import Toggle from './components/Toggle.vue'
import Bay from './components/Bay.vue'
import axios from 'axios'

axios.get('http://localhost:8080/api/health').then(res => console.log(res.data)).catch(err => console.error(err));

const palletData = {
  aisle: 1,
  bay: 2,
  side: "left",
  sku_list: [
    { sku: 123, qty: 5 },
    { sku: 456, qty: 2 },
    { sku: 789, qty: 12 }
  ]
};

axios.post('http://localhost:8080/api/action', {
  action: 1,
  data: palletData
}, {
  headers: {
    'Content-Type': 'application/json'
  }
}).then(res => console.log(res.data)).catch(err => console.error(err));

const bay1="bay1";
const size=24;
const depth=99;
const shelfNum=4;

// --- Types reused ---
type ShelfSKUMap = {
  [shelfNumber: number]: number[]
}
type PalletMap = {
  [PalletID: number]: number
}

interface BayData {
  name: string
  size: number
  depth: number
  shelfNum: number
  shelfSKU: ShelfSKUMap
  palletMap: PalletMap
}

var pCnt=0;

// --- Random SKU generator ---
function generateRandomSKUs(shelfCount: number, skusPerShelf: number): ShelfSKUMap {
  const shelves: ShelfSKUMap = {}
  for (let i = 0; i < shelfCount; i++) {
    shelves[i] = []
    for (let j = 0; j < skusPerShelf; j++) {
      shelves[i].push(Math.floor(Math.random() * 1000))  // random SKU 0-999
    }
  }
  return shelves
}
function palletCnt(palletCount: number): PalletMap{
  const pallets: PalletMap ={}
  for(let i = 0; i < palletCount; i++){
    pallets[i] = pCnt++;
  }
  return pallets
}
// --- Generate 3 bays ---
const bays: BayData[] = []
for (let i = 1; i <= 3; i++) {
  bays.push({
    name: `Bay ${i}`,
    size: 30,
    depth: 200,
    shelfNum: 4,
    shelfSKU: generateRandomSKUs(4, 3),
    palletMap: palletCnt(3)
  })
}

</script>

<template>
  <div class = "wrapper">
    <!--<Message class="compWrap"/>
    <Counter class="compWrap"/>
    <Toggle class="compWrap"/>-->
    
    <Bay 
      v-for="(bay, index) in bays" 
      :key="index"
      :name="bay.name"
      :size="bay.size"
      :depth="bay.depth"
      :shelfNum="bay.shelfNum"
      :shelfSKU="bay.shelfSKU"
      :palletMap="bay.palletMap"
      class="compWrap"
    />
    
  </div>
</template>

<style scoped>

.compWrap{
  font-size: 5vw;
  text-align: center;

  /*background-color: rgb(10,10,10);*/
}

:deep(.shelves p){
  font-size: 30%;
}
:deep(.shelves h2){
  font-size: 50%;
}
.wrapper{
  display: flex;
  flex-wrap: wrap; 
  justify-content: center;
  align-items: center;
  height: 100vh;
  width: 100vw;
  
}

header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

/*@media (min-width: 1024px) {
  header .wrapper{
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
}*/
</style>
