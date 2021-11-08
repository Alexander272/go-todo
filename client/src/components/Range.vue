<template>
    <div class="range">
        <label :for="id" class="label">{{ labelText }}</label>
        <label :for="id" class="value">{{ modelValue }}</label>
        <input
            :id="id"
            :name="name"
            :value="modelValue"
            :max="max"
            :min="min"
            :step="step"
            type="range"
            class="input"
            @input="$emit('update:modelValue', $event.target.value)"
            @change="$emit('update:modelValue', $event.target.value)"
        />
        <p class="desc">
            <span>{{ minText || min }}</span>
            <span>{{ maxText || max }}</span>
        </p>
    </div>
</template>

<script>
import { toRefs } from '@vue/reactivity'
export default {
    name: 'Range',
    emits: ['update:modelValue'],
    props: {
        min: Number,
        minText: String,
        max: Number,
        maxText: String,
        modelValue: String,
        step: {
            type: Number,
            default: 1,
        },
        labelText: {
            type: String,
            required: true,
        },
        id: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
    },
    setup(props) {
        // const { modelValue } = toRefs(props)
        // return { modelValue }
    },
}
</script>

<style lang="scss" scoped>
.range {
    display: flex;
    flex-direction: column;
}
.label {
    padding-left: 1rem;
    text-transform: uppercase;
    font-size: 0.75rem;
}
.value {
    font-size: 0.8rem;
}
.input {
    outline: none;
}

.desc {
    display: flex;
    justify-content: space-between;
    font-size: 0.9em;
    color: rgb(46, 45, 45);
}
</style>