<template>
    <p class="tag" :class="[size, rounded]">
        {{ title }}
    </p>
</template>

<script>
export default {
    name: 'Tag',
    props: {
        title: {
            type: String,
            required: true,
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
                return ['none', 'low', 'medium', 'high', 'round'].includes(value)
            },
            default: 'round',
        },
    },
    setup() {},
}
</script>

<style lang="scss" scoped>
.tag {
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(135deg, #008afb, #6425d3);
    color: #fff;
    cursor: pointer;
}

.small {
    font-size: 0.8rem;
    padding: 5px 10px;
}
.middle {
    font-size: 1rem;
    padding: 7px 14px;
}
.large {
    font-size: 1.2rem;
    padding: 9px 18px;
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
</style>