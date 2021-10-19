<template>
    <router-view />
</template>

<script>
import { onMounted } from '@vue/runtime-core'
import { useStore } from 'vuex'
export default {
    setup() {
        const store = useStore()
        onMounted(() => {
            store.dispatch('auth/refresh').then(() => store.commit('setReady', true))
        })
    },
}
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css2?family=Montserrat&display=swap');

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    font-family: 'Montserrat', sans-serif;
}

// #app {

// }

.wrapper {
    min-height: 100vh;
    min-width: 320px;
    max-width: 1920px;
    margin: 0 auto;
    display: flex;
}
</style>
