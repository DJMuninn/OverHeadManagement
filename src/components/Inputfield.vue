<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

// Reactive pallet object
const pallet = ref({
  id: 0,
  aisle: 0,
  bay: 0,
  side: '',
  sku_list: [] as { sku: number, qty: number }[]
});

// For raw SKU input
const skuInput = ref('');

// Parse SKU input (from text to array)
function parseSkuInput() {
  pallet.value.sku_list = skuInput.value.split(',').map(item => {
    const [sku, qty] = item.split(':').map(Number);
    return { sku, qty };
  });
}

// Insert handler
async function handleInsert() {
  parseSkuInput();
  try {
    const response = await axios.post('http://localhost:8080/api/action', {
      action: 1,
      data: pallet.value
    });
    console.log(response.data);
    alert('Insert successful!');
  } catch (err) {
    console.error(err);
    alert('Insert failed.');
  }
}

// Update handler
async function handleUpdate() {
  parseSkuInput();
  try {
    const response = await axios.post('http://localhost:8080/api/action', {
      action: 3,
      data: pallet.value
    });
    console.log(response.data);
  } catch (err) {
    console.error(err);
  }
}
</script>
<template>
    <div class="input-field">

    <h2>Pallet Input</h2>

    <label>ID (only for update):</label>
    <input v-model.number="pallet.id" type="number" placeholder="ID (for update)" />

    <label>Aisle:</label>
    <input v-model.number="pallet.aisle" type="number" placeholder="Aisle" />

    <label>Bay:</label>
    <input v-model.number="pallet.bay" type="number" placeholder="Bay" />

    <label>Side:</label>
    <input v-model="pallet.side" type="text" placeholder="Side" />

    <label>SKU List (comma-separated SKU:Qty):</label>
    <input v-model="skuInput" type="text" placeholder="Ex: 123:5,456:2" />

    <div class="buttons">
      <button @click="handleInsert">Insert</button>
      <button @click="handleUpdate">Update</button>
    </div>

  </div>
</template>
<style scoped>
    @import './cssTemps/flexComp.css';
    .input-field {
        display: flex;
        flex-direction: column;
        width: 300px;
        margin: auto;
    }

    label {
        margin-top: 10px;
    }

    input {
        margin-bottom: 10px;
        padding: 5px;
    }

    .buttons {
        display: flex;
        justify-content: space-between;
        margin-top: 10px;
    }
</style>