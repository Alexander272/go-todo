<template>
    <div v-if="loading || !ready">Loading ...</div>
    <div v-else class="home">
        <p>Home</p>
    </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import useCheckAuth from '@/composables/useCheckAuth'

export default {
    name: 'Home',
    components: {
        HelloWorld,
    },
    setup() {
        const { loading, ready } = useCheckAuth()

        return { loading, ready }
    },
}
</script>
