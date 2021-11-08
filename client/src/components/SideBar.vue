<template>
    <aside class="sidebar">
        <div class="full" v-if="fetching">
            <loader size="small" />
        </div>
        <h4 class="title" @click="toggleOpenList">
            <app-icon icon="list-ul" class="icon" /> Мои списки
            <app-icon
                v-if="!isEmptyList"
                icon="caret-down"
                class="arrow"
                :class="{ inverse: !isOpenList }"
            />
        </h4>
        <transition name="fade" mode="out-in">
            <nav class="nav" v-if="!isEmptyList" :class="{ hide: !isOpenList }">
                <transition-group name="list">
                    <router-link
                        v-for="item in lists"
                        :key="item.id"
                        :to="`/list/${item.id}/`"
                        class="link"
                    >
                        {{ item.title }}
                        <span
                            :data-id="item.id"
                            @click.prevent="openEditModal($event)"
                            class="link-icon"
                        >
                            <app-icon icon="pen" />
                        </span>
                    </router-link>
                </transition-group>
            </nav>
            <p class="empty" v-else>Ни одного списка еще не создано</p>
        </transition>
        <teleport to="#app">
            <modal :condition="isOpenModal" title="Создать список" @close="toggleOpenModal">
                <template v-slot:default>
                    <input-field
                        id="title"
                        labelText="Название"
                        name="title"
                        styleInput="black"
                        :errorText="newList.errorTitle"
                        v-model="newList.title"
                        class="modal__input"
                    />
                    <textarea-field
                        id="description"
                        labelText="Описание"
                        name="description"
                        styleInput="black"
                        v-model="newList.description"
                        :rows="4"
                        class="modal__input"
                    />
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
                        v-if="newList.id"
                        text="Удалить"
                        rounded="round"
                        variant="danger"
                        icon="trash"
                        class="modal__btn"
                        @click="openConfirm('remove')"
                    />
                    <app-button
                        text="Сохранить"
                        rounded="round"
                        class="modal__btn"
                        icon="save"
                        @click="newList.id !== '' ? openConfirm('update') : addList()"
                    />
                </template>
            </modal>
        </teleport>
        <teleport to="#app">
            <confirm
                :type="typeConfirm === 'update' ? 'warn' : 'danger'"
                :condition="isOpenConfirm"
                :description="`Вы уверены что хотите ${
                    typeConfirm === 'update' ? 'обновить' : 'удалить'
                } список?`"
                :textBtn="typeConfirm === 'update' ? 'Обновить' : 'Удалить'"
                :iconBtn="typeConfirm === 'update' ? 'pen' : 'trash'"
                :variantBtn="typeConfirm === 'update' ? 'primary' : 'danger'"
                @close="closeOpenConfirm"
                @confirm="typeConfirm === 'update' ? updateList() : removeList()"
            ></confirm>
        </teleport>
        <app-button
            text="Создать список"
            rounded="round"
            icon="plus"
            class="btn"
            @on-click="toggleOpenModal"
        />
    </aside>
</template>

<script>
import { computed, reactive, ref } from 'vue'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'
import { FontAwesomeIcon as AppIcon } from '@fortawesome/vue-fontawesome'
import AppButton from './Button.vue'
import Modal from './Modal.vue'
import InputField from './InputField.vue'
import TextareaField from './TextareaField.vue'
import Loader from './Loader.vue'
import Confirm from './Confirm.vue'
export default {
    components: { AppButton, AppIcon, Modal, InputField, TextareaField, Loader, Confirm },
    name: 'SideBar',
    setup() {
        const store = useStore()
        const router = useRouter()
        const route = useRoute()
        const lists = computed(() => store.state.lists.lists)
        const fetching = computed(() => store.state.lists.loading)
        const isEmptyList = computed(() => store.getters['lists/isEmptyList'])

        const newList = reactive({
            id: '',
            title: '',
            errorTitle: '',
            description: '',
        })

        const checkForm = () => {
            if (newList.title.trim() === '') {
                newList.errorTitle = 'Название не может быть пустым'
                return true
            } else newList.errorTitle = ''
            return false
        }

        const isOpenModal = ref(false)
        const toggleOpenModal = () => {
            newList.id = ''
            newList.title = ''
            newList.description = ''
            newList.errorTitle = ''
            isOpenModal.value = !isOpenModal.value
        }

        const isOpenConfirm = ref(false)
        const typeConfirm = ref('update')
        const closeOpenConfirm = () => {
            isOpenConfirm.value = false
        }
        const openConfirm = type => {
            isOpenConfirm.value = true
            if (type === 'update') typeConfirm.value = 'update'
            else typeConfirm.value = 'remove'
        }

        const addList = () => {
            if (checkForm()) return
            store.dispatch('lists/createList', {
                title: newList.title,
                description: newList.description,
            })
            toggleOpenModal()
        }

        const isOpenList = ref(true)
        const toggleOpenList = () => (isOpenList.value = !isOpenList.value)

        const openEditModal = event => {
            const id = event.target.dataset.id
            toggleOpenModal()
            const list = store.getters['lists/getListById'](id)
            newList.id = list.id
            newList.title = list.title
            newList.description = list.description
        }

        const updateList = () => {
            if (checkForm()) return
            store.dispatch('lists/updateList', {
                id: newList.id,
                title: newList.title,
                description: newList.description,
            })
            closeOpenConfirm()
            toggleOpenModal()
        }

        const removeList = () => {
            store.dispatch('lists/removeList', newList.id)
            closeOpenConfirm()
            toggleOpenModal()
            const listId = route.params.listId
            if (listId === newList.id) router.push('/')
        }

        return {
            lists,
            isEmptyList,
            fetching,
            isOpenModal,
            toggleOpenModal,
            newList,
            addList,
            updateList,
            isOpenList,
            toggleOpenList,
            openEditModal,
            removeList,
            isOpenConfirm,
            closeOpenConfirm,
            openConfirm,
            typeConfirm,
        }
    },
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
.sidebar {
    margin: 15px;
    background-color: #fff;
    border-radius: 12px;
    padding: 20px;
    flex-basis: 15%;
    box-sizing: border-box;
    box-shadow: 3px 3px 6px #0000002e;
    position: relative;
    overflow: hidden;
}
.full {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    // background-color: #0000002e;
    background-color: #121a7033;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 10;
}

.icon {
    margin-right: 10px;
}
.arrow {
    margin-left: auto;
    transform: rotate(0deg);
    transition: 0.4s all ease-in-out;
}
.inverse {
    transform: rotate(180deg);
}

.title {
    color: $primaryColor;
    display: flex;
    align-items: center;
    margin-bottom: 5px;
    cursor: pointer;
}

.empty {
    padding: 15px;
    font-size: 0.9rem;
    color: #727272;
    text-align: center;
    padding-bottom: 0;
}

.btn {
    width: 100%;
    margin-top: 20px;
}

.nav {
    padding: 10px 15px;
    // margin: 0 10px;
    border-bottom: 1px dashed $primaryColor;
    transition: 0.4s all ease-in-out;
    overflow: auto;
    max-height: 45vh;

    &::-webkit-scrollbar {
        width: 10px;
    }
    &::-webkit-scrollbar-track {
        box-shadow: 5px 5px 5px -5px rgba(34, 60, 80, 0.2) inset;
        background-color: #f9f9fd;
        border-radius: 10px;
    }
    &::-webkit-scrollbar-thumb {
        border-radius: 10px;
        background: linear-gradient(180deg, #008afb, #6425d3);
    }
}
.hide {
    max-height: 0%;
    padding: 0 20px;
    overflow: hidden;
}
.link {
    padding: 5px 20px;
    color: #000;
    transition: all 0.4s ease-in-out;
    background-color: transparent;
    text-decoration: none;
    display: flex;
    align-items: center;
    flex-basis: 100%;
    border-radius: 50px;

    &:hover {
        color: $primaryColor;

        .link-icon {
            color: $primaryColor;
        }
    }

    &-icon {
        margin-left: auto;
        color: #fff;
        padding-left: 10px;
        transition: all 0.4s ease-in-out;
        svg {
            pointer-events: none;
        }
    }
}
.router-link-active {
    color: #fff;
    background-color: $primaryColor;

    &:hover {
        color: #fff;

        .link-icon {
            color: #fff;
        }
    }

    .link-icon {
        color: $primaryColor;
        svg {
            pointer-events: none;
        }
    }
}
.list-enter-from,
.list-leave-to {
    opacity: 0;
    transform: translateX(-30px);
}

.list-leave-active {
    position: absolute;
}

.modal__input {
    margin-top: 15px;
}
.modal__btn {
    flex-grow: 1;
    margin: 0 5px;
}
</style>