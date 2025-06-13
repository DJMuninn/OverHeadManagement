<script setup>
    import { ref , watch } from 'vue'

    const tgl = ref(true)
    function tglBtn(){
        tgl.value = !tgl.value;
    }
</script>

<template>
    <div class = "comp">
        <div class = "but">
            <input type ="button" @click="tglBtn()" value = "Toggle!" style="margin: 0 1vw;"></input>    
            <plaintext v-if="tgl">True</plaintext >
            <plaintext  v-else>False</plaintext >
        </div>
    </div>
</template>

<style>
    @import './cssTemps/flexComp.css';
    .but{
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 0;
    }
</style>