<template>
    <div class="wrapper">
        <div class="full" v-if="loading || !ready">
            <loader size="large" />
        </div>

        <div v-else class="container">
            <side-bar />
            <main class="main">
                <router-view></router-view>
            </main>
        </div>
    </div>
</template>

<script>
// @ is an alias to /src
import { computed, watch } from '@vue/runtime-core'
import { useStore } from 'vuex'
import useCheckAuth from '@/composables/useCheckAuth'
import TodoList from '@/components/TodoList.vue'
import Loader from '../components/Loader.vue'
import SideBar from '../components/SideBar.vue'

export default {
    name: 'Home',
    components: { TodoList, Loader, SideBar },
    setup() {
        const { loading, ready } = useCheckAuth()
        const store = useStore()
        const lists = computed(() => store.state.lists.lists)
        const fetching = computed(() => store.state.lists.loading)
        const isEmptyList = computed(() =>
            store.state.lists.lists.value ? store.state.lists.lists.length === 0 : true
        )

        const getList = condition => {
            if (condition && store.getters['auth/isAuth']) store.dispatch('lists/getLists')
        }
        getList(ready)

        watch(ready, newValue => {
            getList(newValue)
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
    width: 100%;
}

.container {
    display: flex;
    flex-basis: 100%;
}

.main {
    margin: 15px;
    background-color: #fff;
    border-radius: 12px;
    padding: 20px;
    flex-basis: 85%;
    box-sizing: border-box;
    box-shadow: 3px 3px 6px #0000002e;
}
</style>
