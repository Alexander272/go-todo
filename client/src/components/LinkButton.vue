<template>
    <div class="block" :class="[variant, size, rounded]">
        <div class="dropdown"></div>
        <router-link :to="path" class="link">
            {{ text }}
        </router-link>
    </div>
</template>

<script>
export default {
    name: "LinkButton",
    props: {
        path: {
            type: String,
            default: "/",
        },
        text: {
            type: String,
            required: true,
        },
        disabled: {
            type: Boolean,
            default: false,
        },
        variant: {
            type: String,
            validator(value) {
                return ["primary", "danger", "danger"].includes(value);
            },
            default: "primary",
        },
        size: {
            type: String,
            validator(value) {
                return ["small", "middle", "large"].includes(value);
            },
            default: "middle",
        },
        rounded: {
            type: String,
            validator(value) {
                return [
                    "none",
                    "low",
                    "medium",
                    "high",
                    "round",
                    "circle",
                ].includes(value);
            },
            default: "medium",
        },
    },
    setup() {
        return {};
    },
};
</script>

<style lang="scss" scoped>
.block {
    transition: 0.3s all ease-out;
    border-radius: 16px;
    display: flex;
    overflow: hidden;
    position: relative;

    &:hover {
        .link {
            color: #fff;
        }
        .dropdown {
            width: 100%;
        }
    }
}

.link {
    text-decoration: none;
    transition: 0.3s all ease-out;
    text-transform: uppercase;
    display: flex;
    width: 100%;
    min-height: 100%;
    justify-content: center;
    font-weight: bold;
    z-index: 5;
}

.dropdown {
    transition: 0.3s all ease-out;
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 0%;
}

.primary {
    border: 2px solid #673ab7;
    .dropdown {
        background-color: #673ab7;
    }
}

.small .link {
    border-width: 1px;
    font-size: 0.8rem;
    padding: 6px 10px;
}
.middle .link {
    font-size: 1rem;
    padding: 10px 15px;
}
.large .link {
    font-size: 1.2rem;
    padding: 14px 20px;
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