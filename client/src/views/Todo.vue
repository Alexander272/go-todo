<template>
    <div>Todo</div>
</template>

<script>
import useCheckAuth from '@/composables/useCheckAuth'
export default {
    name: 'Todo',
    setup() {
        const { loading, ready } = useCheckAuth()

        return { loading, ready }
    },
}
</script>

<style lang="scss" scoped>
</style>