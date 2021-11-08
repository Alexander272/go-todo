<template>
    <button
        class="button"
        @click="$emit('onClick')"
        :disabled="disabled"
        :class="[variant, size, rounded]"
    >
        <app-icon v-if="icon" :icon="icon" class="icon" />
        {{ text }}
    </button>
</template>

<script>
import { FontAwesomeIcon as AppIcon } from '@fortawesome/vue-fontawesome'
export default {
    name: 'Button',
    components: { AppIcon },
    emits: ['onClick'],
    props: {
        text: {
            type: String,
            required: true,
        },
        disabled: {
            type: Boolean,
            default: false,
        },
        icon: String,
        variant: {
            type: String,
            validator(value) {
                return ['primary', 'danger', 'ghost'].includes(value)
            },
            default: 'primary',
        },
        size: {
            type: String,
            validator(value) {
                return ['small', 'middle', 'large'].includes(value)
            },
            default: 'middle',
        },
        rounded: {
            type: String,
            validator(value) {
                return ['none', 'low', 'medium', 'high', 'round', 'circle'].includes(value)
            },
            default: 'medium',
        },
    },
    setup() {},
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
$dangerColor: #e32525;
.button {
    transition: 0.2s all ease-out;
    border-radius: 12px;
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
    position: relative;
    background-color: transparent;
    border: none;
    cursor: pointer;
    // text-align: center;

    &:active {
        box-shadow: 5px 5px 12px 0px #00030675 inset;
    }

    &:disabled {
        background-color: #c3c3c3;
        color: #919191;
        cursor: not-allowed;
    }
}
.icon {
    margin-right: 10px;
    pointer-events: none;
}

.primary {
    border: 2px solid $primaryColor;
    background: $primaryColor;
    color: #fff;

    &:hover {
        background: #782bff;
        border-color: #782bff;
    }
}
.danger {
    border: 2px solid $dangerColor;
    background: $dangerColor;
    color: #fff;
}
.ghost {
    border: 2px solid $primaryColor;
    background: transparent;
    color: $primaryColor;

    &:hover {
        border-color: #8946ff;
        color: #8946ff;
        background: #7a84f51a;
    }
}

.small {
    border-width: 1px;
    font-size: 0.8rem;
    padding: 5px 10px;
}
.middle {
    font-size: 1rem;
    padding: 7px 15px;
}
.large {
    font-size: 1.2rem;
    padding: 11px 20px;
}

.none {
    border-radius: 0;
}
.low {
    border-radius: 8px;
}
.medium {
    border-radius: 12px;
}
.high {
    border-radius: 16px;
}
.round {
    border-radius: 90px;
}
.circle {
    border-radius: 50%;
}
</style>