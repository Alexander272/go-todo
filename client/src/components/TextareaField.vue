<template>
    <div class="filed" :class="styleInput">
        <label :for="id" class="label"> {{ labelText }} </label>
        <textarea
            :id="id"
            class="area"
            :value="modelValue"
            :cols="cols"
            :rows="rows"
            @input="$emit('update:modelValue', $event.target.value)"
        />
        <p v-if="errorText" class="error">{{ errorText }}</p>
    </div>
</template>

<script>
export default {
    name: 'TextareaField',
    emits: ['update:modelValue'],
    props: {
        modelValue: String,
        labelText: {
            type: String,
            required: true,
        },
        cols: Number,
        rows: Number,
        id: {
            type: String,
            required: true,
        },
        name: {
            type: String,
            required: true,
        },
        styleInput: {
            type: String,
            validator(value) {
                return ['white', 'black'].includes(value)
            },
            default: 'white',
        },
        errorText: String,
    },
}
</script>

<style lang="scss" scoped>
$primaryColor: #6425d3;
$dangerColor: #e32525;

.filed {
    margin-bottom: 1.5rem;
    position: relative;
}

.white {
    .label {
        color: rgba(255, 255, 255, 0.5);
    }
    .area {
        color: white;
        background-color: rgba(255, 255, 255, 0.25);

        &:hover,
        &:focus {
            color: white;
            border: 2px solid rgba(255, 255, 255, 0.5);
            background-color: transparent;
        }
    }
}

.black {
    .label {
        color: #000;
        margin-bottom: 0.4rem;
    }
    .area {
        color: #000;
        background-color: #ededed;
        border: 1px solid #ededed;
        transition: 0.3s all ease-in-out;

        &:hover,
        &:focus {
            color: #000;
            border-color: $primaryColor;
            background-color: transparent;
        }
    }
}

.label {
    display: block;
    padding-left: 1rem;
    text-transform: uppercase;
    font-size: 0.75rem;
    margin-bottom: 1rem;
}

.area {
    font-size: 1.15rem;
    width: 100%;
    padding: 0.5rem 1rem;
    border: 2px solid transparent;
    outline: none;
    border-radius: 1.1rem;
    background-color: rgba(255, 255, 255, 0.25);
    letter-spacing: 1px;
    transition: 0.3s all ease-in-out;
    resize: none;
}

.error {
    font-size: 0.8rem;
    color: $dangerColor;
    padding: 3px 15px;
    position: absolute;
}

// .input + .label {
//     margin-top: 1.5rem;
// }
</style>