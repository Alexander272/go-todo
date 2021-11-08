<template>
    <div class="item" :class="{ complete: done }">
        <div class="done">
            <checkbox :value="done" :id="id" @update="setDone" />
        </div>
        <p class="title">{{ title }}</p>
        <p class="date">Сделать до: {{ deadline }}</p>
        <p class="description">{{ description }}</p>
    </div>
</template>

<script>
import { dateFormat } from '@/utils/dateFormat'
import { useStore } from 'vuex'
import { computed } from '@vue/reactivity'
import Checkbox from './Checkbox.vue'
export default {
    components: { Checkbox },
    name: 'TodoListItem',
    props: {
        id: String,
        title: String,
        description: String,
        done: Boolean,
        createdAt: Number,
        deadlineAt: Number,
        // tags: [],
    },
    setup(props) {
        const store = useStore()
        const deadline = computed(() => dateFormat(+props.deadlineAt * 1000))
        const setDone = target => {
            console.log(target)
            const id = target.dataset.id
            // todo добавить listId  в запрос, без него не работает
            store.dispatch('todo/setTodoDone', { id, done: target.value })
        }

        return { deadline, setDone }
    },
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
.item {
    display: grid;
    grid-template-columns: 30px 0.9fr 1fr 3fr;
    grid-column-gap: 9px;
    padding: 7px;
    border-bottom: 1px solid $primaryColor;
}
.done {
    display: flex;
    justify-content: center;
    align-items: center;
}
.complete {
    .title,
    .description,
    .date {
        text-decoration: line-through;
        color: #99b3ff;
    }
}

.title {
    font-size: 1.3rem;
    display: flex;
    align-items: center;
}
.description {
    font-size: 0.9rem;
}
</style>