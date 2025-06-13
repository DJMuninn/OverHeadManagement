<script setup>
    import { reactive , watch } from 'vue'

    const counter = reactive({
        count: 0
    })
    function cntBtn(){
        counter.count++;
    }
    watch(() => counter.count, (newVal) => {
        if (newVal === 10) {
            counter.count = 0
        }
    })
</script>

<template>
    <div class = "comp">
        <div class = "but">
            <input type ="button" @click="cntBtn()" value = "Click Me!" style="margin: 0 1vw;"></input>    
            {{ counter.count }}
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