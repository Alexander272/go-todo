<template>
    <div class="conrainer">
        <transition name="fade" mode="out-in">
            <div v-if="!list || list?.id !== listId" class="full">
                <loader />
            </div>
        </transition>
        <div class="header">
            <div class="header-top">
                <h2 class="title">{{ list?.title }}</h2>
                <p v-if="list?.createdAt" class="date">Дата создания: {{ created }}</p>
            </div>

            <p>{{ list?.description }}</p>
        </div>
        <div class="buttons">
            <app-button
                text="Создать задачу"
                rounded="round"
                icon="plus"
                size="small"
                @on-click="toggleOpenModal"
            />
        </div>
        <todo-list />
        <teleport to="#app">
            <modal :condition="isOpenModal" title="Создать список" @close="toggleOpenModal">
                <template v-slot:default>
                    <input-field
                        id="title"
                        labelText="Название"
                        name="title"
                        styleInput="black"
                        class="modal__input"
                        :errorText="todo.errorTitle"
                        v-model="todo.title"
                    />
                    <textarea-field
                        id="description"
                        labelText="Описание"
                        name="description"
                        styleInput="black"
                        :rows="4"
                        class="modal__input"
                        v-model="todo.description"
                    />
                    <input-field
                        type="datetime-local"
                        id="date"
                        name="date"
                        labelText="Срок выполнения"
                        styleInput="black"
                        :errorText="todo.errorDeadline"
                        v-model="todo.deadlineAt"
                    />
                    <range
                        id="range"
                        labelText="Приоритет задачи"
                        name="range"
                        :min="1"
                        :max="10"
                        maxText="Масимальный"
                        minText="Минимальный"
                        v-model="todo.priority"
                        class="modal__range"
                    />
                    <!-- <div class="tags">
                            <tag title="test" />
                        </div> -->
                </template>
                <template v-slot:footer>
                    <app-button
                        text="Отмена"
                        rounded="round"
                        variant="ghost"
                        class="modal__btn"
                        @on-click="toggleOpenModal"
                    />
                    <app-button
                        text="Сохранить"
                        rounded="round"
                        class="modal__btn"
                        icon="save"
                        @on-click="addTodo"
                    />
                </template>
            </modal>
        </teleport>
    </div>
</template>

<script>
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { computed, reactive, ref } from '@vue/reactivity'
import { dateFormat } from '@/utils/dateFormat'
import AppButton from '@/components/Button.vue'
import TodoList from '@/components/TodoList.vue'
import Modal from '@/components/Modal.vue'
import InputField from '@/components/InputField.vue'
import TextareaField from '@/components/TextareaField.vue'
import Range from '@/components/Range.vue'
import Loader from '@/components/Loader.vue'
import Tag from '@/components/Tag.vue'
import useCheckAuth from '@/composables/useCheckAuth'
export default {
    name: 'Todo',
    components: { AppButton, TodoList, Modal, InputField, TextareaField, Range, Loader, Tag },
    setup() {
        const store = useStore()
        const route = useRoute()
        const { loading, ready } = useCheckAuth()
        const listId = computed(() => route.params.listId)
        const list = computed(() => store.getters['lists/getListById'](listId.value))
        const created = computed(() => dateFormat(+list.value?.createdAt * 1000))

        const isOpenModal = ref(false)
        const toggleOpenModal = () => (isOpenModal.value = !isOpenModal.value)

        const todo = reactive({
            title: '',
            errorTitle: '',
            description: '',
            deadlineAt: '',
            errorDeadline: '',
            priority: '5',
            tags: [],
        })

        const checkForm = () => {
            let isValid = todo.title.trim() !== ''
            if (!isValid) todo.errorTitle = 'Название не может быть пустым'
            else todo.errorTitle = ''
            if (todo.deadlineAt !== '') {
                isValid = Date.parse(todo.deadlineAt) > Date.now()
                if (!isValid)
                    todo.errorDeadline = 'Срок окончания не можеть меньше текущего времени'
                else todo.errorDeadline = ''
            }
            return isValid
        }

        const addTodo = () => {
            if (!checkForm()) return
            store.dispatch('todo/createTodo', {
                listId: listId.value,
                title: todo.title,
                description: todo.description,
                deadlineAt: +(Date.parse(todo.deadlineAt) / 1000).toFixed(),
                priority: +todo.priority,
                tags: todo.tags,
            })
            toggleOpenModal()
        }

        return {
            loading,
            ready,
            listId,
            list,
            created,
            todo,
            isOpenModal,
            toggleOpenModal,
            addTodo,
        }
    },
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
$dangerColor: #e32525;
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
    opacity: 0;
}

.conrainer {
    display: flex;
    flex-direction: column;
    position: relative;
    height: 100%;
    width: 100%;
}

.full {
    position: absolute;
    z-index: 15;
    background: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
}

.header {
    padding: 0 15px 15px;
    border-bottom: 1px solid #bfbfbf;
}

.header-top {
    display: flex;
    align-items: flex-end;
    margin-bottom: 10px;
}
.title {
    color: $primaryColor;
    margin-right: 20px;
}
.date {
    font-size: 0.7em;
    color: rgb(116, 116, 116);
    line-height: 20px;
}

.modal__input {
    margin-top: 15px;
}
.modal__range {
    margin-bottom: 15px;
}
.tags {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    flex-wrap: wrap;
}
.modal__btn {
    flex-grow: 1;
    margin: 0 5px;
}

.buttons {
    padding: 10px;
    border-bottom: 1px solid #bfbfbf;
}
</style>