<template>
    <div class="backdrop">
        <div class="modal">
            <div class="header">
                <p class="title">{{ title }}</p>
                <app-icon class="icon" icon="times" @click="$emit('close')" />
            </div>
            <div class="body">
                <slot></slot>
            </div>
            <div class="footer">
                <slot name="footer"></slot>
            </div>
        </div>
    </div>
</template>

<script>
import { FontAwesomeIcon as AppIcon } from '@fortawesome/vue-fontawesome'
export default {
    name: 'Modal',
    components: { AppIcon },
    emits: ['close'],
    props: {
        title: String,
    },
    setup() {},
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
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
}
.modal {
    min-width: 450px;
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
    justify-content: space-between;
    padding: 0 5px 12px 5px;
    // padding-bottom: 7px;
    border-bottom: 1px solid #cdcdcd;
}
.title {
    font-size: 1.5em;
    color: $primaryColor;
}
.icon {
    cursor: pointer;
    transition: 0.3s all ease-in-out;

    &:hover {
        color: $primaryColor;
    }
}

.footer {
    border-top: 1px solid #cdcdcd;
    padding-top: 12px;
    display: flex;
    justify-content: space-between;
}
</style>