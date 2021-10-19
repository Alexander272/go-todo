<template>
    <div class="filed">
        <label :for="id" class="label"> {{ labelText }} </label>
        <input
            :id="id"
            class="input"
            :type="type"
            :value="modelValue"
            @input="$emit('update:modelValue', $event.target.value)"
        />
        <p v-if="errorText" class="error">{{ errorText }}</p>
    </div>
</template>

<script>
export default {
    name: 'InputFiled',
    emits: ['update:modelValue'],
    props: {
        modelValue: String,
        type: {
            type: String,
            default: 'text',
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
        errorText: String,
    },
}
</script>

<style lang="scss" scoped>
.filed {
    margin-bottom: 1.5rem;
    position: relative;
}

.label {
    display: block;
    padding-left: 1rem;
    color: rgba(255, 255, 255, 0.5);
    text-transform: uppercase;
    font-size: 0.75rem;
    margin-bottom: 1rem;
}

.input {
    color: white;
    font-size: 1.15rem;
    width: 100%;
    padding: 0.5rem 1rem;
    border: 2px solid transparent;
    outline: none;
    border-radius: 1.5rem;
    background-color: rgba(255, 255, 255, 0.25);
    letter-spacing: 1px;
    transition: 0.3s all ease-in-out;

    &:hover,
    &:focus {
        color: white;
        border: 2px solid rgba(255, 255, 255, 0.5);
        background-color: transparent;
    }
}

.error {
    font-size: 0.8rem;
    color: red;
    padding: 3px 15px;
    position: absolute;
}

// .input + .label {
//     margin-top: 1.5rem;
// }
</style>