<template>
    <aside class="sidebar">
        <h4 class="title">
            <app-icon icon="list-ul" class="icon" /> Мои списки
            <app-icon v-if="!isEmptyList" icon="caret-down" class="arrow" />
        </h4>
        <nav class="nav" v-if="!isEmptyList">
            <router-link
                v-for="item in lists"
                :key="item.id"
                :to="`/list/${item.id}/`"
                class="link"
                >{{ item.title }}</router-link
            >
        </nav>
        <p class="empty" v-else>Ни одного списка еще не создано</p>
        <teleport to="#app">
            <modal v-if="isOpenModal" title="Создать список" @close="toggleOpenModal">
                <template v-slot:default>
                    <input-field
                        id="title"
                        labelText="Название"
                        name="title"
                        styleInput="black"
                        errorText="error"
                        class="modal__input"
                    />
                    <input-field
                        id="description"
                        labelText="Описание"
                        name="description"
                        styleInput="black"
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
                    <app-button text="Сохранить" rounded="round" class="modal__btn" icon="save" />
                </template>
            </modal>
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
import { computed, ref } from '@vue/reactivity'
import { useStore } from 'vuex'
import { FontAwesomeIcon as AppIcon } from '@fortawesome/vue-fontawesome'
import AppButton from './Button.vue'
import Modal from './Modal.vue'
import InputField from './InputField.vue'
export default {
    components: { AppButton, AppIcon, Modal, InputField },
    name: 'SideBar',
    setup() {
        const store = useStore()
        const lists = computed(() => store.state.lists.lists)
        const isEmptyList = computed(() =>
            store.state.lists.lists.value ? store.state.lists.lists.length === 0 : true
        )

        const isOpenModal = ref(false)
        const toggleOpenModal = () => (isOpenModal.value = !isOpenModal.value)

        return { lists, isEmptyList, isOpenModal, toggleOpenModal }
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
}
.icon {
    margin-right: 10px;
}
.arrow {
    margin-left: auto;
    cursor: pointer;
}

.title {
    color: $primaryColor;
    display: flex;
    align-items: center;
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
}
.link {
    padding: 5px 20px;
    color: #000;
    transition: all 0.4s ease-in-out;
    background-color: transparent;
    text-decoration: none;
    display: flex;
    flex-basis: 100%;
    border-radius: 50px;

    &:hover {
        color: $primaryColor;
    }
}
.router-link-active {
    color: #fff;
    background-color: $primaryColor;

    &:hover {
        color: #fff;
    }
}

.modal__input {
    margin-top: 15px;
}
.modal__btn {
    flex-grow: 1;
    margin: 0 5px;
}
</style>