<template>
    <div class="full" v-if="loading || !ready">
        <loader size="large" />
    </div>

    <div v-else class="container">
        <p>Home</p>
        <div v-if="fetching" class="part"><loader /></div>

        <div class="lists" v-if="isEmptyList">
            <todo-list
                v-for="item in lists"
                :key="item.id"
                :title="item.title"
                :description="item.description"
                :createdAt="item.createdAt"
            />
        </div>
        <p v-else>Ни одного списка не создано</p>
    </div>
</template>

<script>
// @ is an alias to /src
import { computed, watch } from '@vue/runtime-core'
import { useStore } from 'vuex'
import useCheckAuth from '@/composables/useCheckAuth'
import TodoList from '@/components/TodoList.vue'
import Loader from '../components/Loader.vue'

export default {
    name: 'Home',
    components: { TodoList, Loader },
    setup() {
        const { loading, ready } = useCheckAuth()
        const store = useStore()
        const lists = computed(() => store.state.lists.lists)
        const fetching = computed(() => store.state.lists.loading)
        const isEmptyList = computed(() =>
            store.state.lists.lists.value ? store.state.lists.lists.length === 0 : true
        )

        watch(ready, newValue => {
            if (newValue && store.getters['auth/isAuth']) {
                store.dispatch('lists/getLists')
            }
        })

        return { loading, ready, lists, isEmptyList, fetching }
    },
}
</script>

<style lang="scss" scoped>
.full {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
}

.part {
    display: flex;
    justify-content: center;
    align-items: center;
}
</style>
