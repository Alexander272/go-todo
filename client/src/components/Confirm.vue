<template>
    <transition name="fade">
        <div v-if="condition" class="backdrop">
            <div class="modal">
                <div class="header">
                    <div class="warn-icon" :class="type">
                        <app-icon icon="exclamation-triangle" />
                    </div>
                    <p class="title">Потверждение</p>
                    <app-icon class="icon" icon="times" @click="$emit('close')" />
                </div>
                <div class="body">
                    <p>{{ description }}</p>
                </div>
                <div class="footer">
                    <app-button
                        text="Отмена"
                        rounded="round"
                        variant="ghost"
                        class="btn"
                        @on-click="$emit('close')"
                    />
                    <app-button
                        :text="textBtn"
                        rounded="round"
                        :variant="variantBtn"
                        class="btn"
                        :icon="iconBtn"
                        @click="$emit('confirm')"
                    />
                </div>
            </div>
        </div>
    </transition>
</template>

<script>
import { FontAwesomeIcon as AppIcon } from '@fortawesome/vue-fontawesome'
import AppButton from '@/components/Button.vue'
export default {
    name: 'Modal',
    components: { AppIcon, AppButton },
    emits: ['close', 'confirm'],
    props: {
        type: {
            type: String,
            default: 'warn',
            validator(value) {
                return ['warn', 'danger'].includes(value)
            },
        },
        description: String,
        condition: {
            type: Boolean,
            default: false,
        },
        textBtn: String,
        iconBtn: String,
        variantBtn: String,
    },
    setup() {},
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

.backdrop {
    position: fixed;
    top: 0;
    left: 0;
    height: 100vh;
    width: 100%;
    background-color: #837f9157;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 70;
}
.modal {
    min-width: 550px;
    z-index: 500;
    background-color: #fff;
    border-radius: 12px;
    padding: 25px;
    padding-bottom: 15px;
    box-shadow: 3px 3px 8px #6d6978;
}
.header {
    display: flex;
    align-items: center;
    padding: 0 5px 12px 5px;
    // padding-bottom: 7px;
    // border-bottom: 1px solid #cdcdcd;
}
.warn-icon {
    font-size: 1.6em;
    margin-right: 20px;
}
.warn.warn-icon {
    color: #fabf14;
}

.danger.warn-icon {
    color: $dangerColor;
}

.title {
    font-size: 1.5em;
    color: $primaryColor;
    margin-right: auto;
}
.icon {
    cursor: pointer;
    transition: 0.3s all ease-in-out;
    margin-left: 10px;

    &:hover {
        color: $primaryColor;
    }
}

.footer {
    margin-top: 15px;
    border-top: 1px solid #cdcdcd;
    padding-top: 12px;
    display: flex;
    justify-content: space-between;
}
.btn {
    flex-grow: 1;
    margin: 0 5px;
}
</style>