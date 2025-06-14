<template>
  <div class="delete-field">

    <h2>Delete Pallet</h2>

    <label>ID to delete:</label>
    <input v-model.number="deleteId" type="number" placeholder="Enter ID to delete" />

    <button @click="handleDelete">Delete</button>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import axios from 'axios';

const deleteId = ref(0);

async function handleDelete() {
  if (deleteId.value <= 0) {
    alert("Please enter a valid ID.");
    return;
  }

  try {
    const response = await axios.post('http://localhost:8080/api/action', {
      action: 2,
      data: { id: deleteId.value, aisle: 0, bay: 0, side: "", sku_list: [] }
    });
    console.log(response.data);
  } catch (err: any) {
    console.error(err);
    if (err.response && err.response.data) {
      alert("Delete failed: " + err.response.data);
    } else {
      alert('Delete failed.');
    }
  }
}
</script>

<style scoped>
.delete-field {
  display: flex;
  flex-direction: column;
  width: 300px;
  margin: 20px auto;
}

label {
  margin-top: 10px;
}

input {
  margin-bottom: 10px;
  padding: 5px;
}

button {
  padding: 5px 10px;
}
</style>