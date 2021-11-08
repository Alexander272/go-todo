<template>
    <transition name="fade" mode="out-in">
        <div class="loader" v-if="loading || !ready || !readyTodo">
            <loader size="middle" />
        </div>
        <div v-else class="lists">
            <template v-if="!isEmptyList">
                <div v-if="fetching" class="lists__loader">
                    <loader size="middle" />
                </div>
                <h5 class="title">Задачи:</h5>
                <todo-list-item
                    v-for="item in todos"
                    :key="item.id"
                    :id="item.id"
                    :title="item.title"
                    :description="item.description"
                    :done="item.done"
                    :deadlineAt="item.deadlineAt"
                    :createdAt="item.createdAt"
                />
            </template>
            <p class="empty" v-else>Задачи еще не созданы</p>
        </div>
    </transition>
</template>

<script>
import { computed } from '@vue/reactivity'
import { watch } from '@vue/runtime-core'
import { useStore } from 'vuex'
import { useRoute } from 'vue-router'
import useCheckAuth from '@/composables/useCheckAuth'
import TodoListItem from './TodoListItem.vue'
import Loader from './Loader.vue'
export default {
    components: { TodoListItem, Loader },
    name: 'TodoList',
    setup() {
        const store = useStore()
        const route = useRoute()
        const { loading, ready } = useCheckAuth()

        const listId = computed(() => route.params.listId)
        const todos = computed(() => store.state.todo.todos)
        const isEmptyList = computed(() => store.getters['todo/isEmptyList'])
        const fetching = computed(() => store.state.todo.loading)
        const readyTodo = computed(() => store.state.todo.ready)

        const getTodos = condition => {
            if (condition && store.getters['auth/isAuth'])
                store.dispatch('todo/getTodos', listId.value)
        }
        getTodos(ready)

        watch(ready, newValue => {
            getTodos(newValue)
        })
        watch(listId, () => {
            getTodos(ready)
        })

        return { loading, ready, todos, isEmptyList, fetching, readyTodo }
    },
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.loader {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-grow: 1;
}

.lists {
    margin-top: 15px;
    border-radius: 12px;
    display: flex;
    flex-direction: column;
    position: relative;
    flex-grow: 1;

    &__loader {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background-color: #121a7033;
        border-radius: 12px;
        display: flex;
        justify-content: center;
        align-items: center;
    }
}

.title {
    color: $primaryColor;
    font-size: 1.3rem;
    margin-bottom: 15px;
    margin-left: 1em;
}
</style>