<template>
    <div class="list">
        <div class="header">
            <h5 class="title">{{ title }}</h5>
            <p class="date">Дата создания: {{ created }}</p>
        </div>
        <div>
            <p class="description">{{ description }}</p>
        </div>
        <div class="footer"></div>
    </div>
</template>

<script>
import { computed } from '@vue/reactivity'
import { dateFormat } from '@/utils/dateFormat'
export default {
    name: 'TodoList',
    props: {
        title: {
            type: String,
            required: true,
        },
        description: String,
        createdAt: String,
    },
    setup(props) {
        const created = computed(() => dateFormat(+props.createdAt))

        return { created }
    },
}
</script>

<style lang="scss" scoped>
.list {
    border-radius: 12px;
    padding: 15px;
}
.header {
    padding: 0 10px 10px;
    border-bottom: 1px soild #eee;
}
.title {
    font-size: 1.2rem;
}
.date {
    font-size: 0.9rem;
    color: rgb(78, 80, 97);
}
.description {
    text-align: justify;
}
</style>