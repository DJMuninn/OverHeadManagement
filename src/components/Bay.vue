<script setup lang="ts">
    import { reactive , watch} from 'vue'

    type ShelfSKUMap = {
        [shelfID: number]: number[]
    }
    type PalletMap = {
        [PalletID: number]: number
    }

    interface BayDim {
        name: string,
        size: number,
        depth: number,
        shelfNum: number,
        shelfSKU: ShelfSKUMap,
        palletMap: PalletMap
    }
    
    const bayDim = defineProps<BayDim>()
    const heightPercent = `${bayDim.size}px`
    const widthPercent = `${bayDim.depth}px`

    function getGridStyle(count: number, boxWidth: number, boxHeight: number): Record<string, string> {
        let columns = 1
        let rows = 1
        
        if (count < 3) {
            columns = count
            rows = 1
        } else {
            console.log(count)
            columns = Math.ceil(Math.sqrt(count))
            rows = Math.ceil(count / columns)
        }
        const fontSize = `${Math.min(boxWidth / columns, boxHeight / rows) * 0.5}px`
        console.log(fontSize)
        return {
            display: 'grid',
            gridTemplateColumns: `repeat(${columns}, 1fr)`,
            gridTemplateRows: `repeat(${rows}, 1fr)`,
            width: '100%',
            height: '100%',
            fontSize: fontSize
        }
    }

</script>

<template>
    <div class = "comp">
        <div class = "but" :style="{ width: widthPercent, height: heightPercent }">
            <div class = "shelves">
                <div class = "bayName">
                    <h2>{{ bayDim.name }}</h2>
                </div>
                <div class="Pallets" :style="getGridStyle(Object.keys(bayDim.palletMap).length, bayDim.size,bayDim.depth)">
                    <div class="pallet" v-for="[PalletID, PalletValue] in Object.entries(bayDim.palletMap)" :key="PalletID">
                        {{ PalletValue }}
                    </div>
                </div>
                <!--<p v-for="[shelfID, skus] in Object.entries(bayDim.shelfSKU)" :key="shelfID">
                    Shelf {{ shelfID }}: {{ skus.join(', ') }}
                </p>-->
            </div>
        </div>
    </div>
</template>

<style>
    @import './cssTemps/flexComp.css';
    .but {
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: black;
        margin: 0;
        border: 1px solid white;
        width: 100%;    /* <-- ADD THIS */
        height: 100%;   /* <-- ADD THIS */
    }   

    .bayName {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        font-size: 2rem;
        color: rgba(255, 255, 255, 0.3); /* Transparent white */
        z-index: 1;
        pointer-events: none;
    }
    .shelves {
        width: 100%;
        height: 100%;
        position: relative;  /* this is important because your bayName is absolutely positioned inside it */
    }
    .pallet {
        display: flex;
        justify-content: center;
        align-items: center;
        border: 1px solid rgba(244,244,244,0.2);
    }

</style>