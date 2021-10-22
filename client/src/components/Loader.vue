<template>
    <!-- <div class="container"> -->
    <div class="loading" :class="size">
        <div class="arc"></div>
        <div class="arc"></div>
        <div class="arc"></div>
    </div>
    <!-- </div> -->
</template>

<script>
export default {
    name: 'Loader',
    props: {
        size: {
            type: String,
            validator(value) {
                return ['small', 'middle', 'large'].includes(value)
            },
            default: 'middle',
        },
    },
}
</script>

<style lang="scss" scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
}

.small {
    width: 2rem;
    height: 2rem;
}
.middle {
    width: 4rem;
    height: 4rem;
}
.large {
    width: 6rem;
    height: 6rem;
}
.loading {
    position: relative;
    transform-style: preserve-3d;
    perspective: 800px;

    .arc {
        position: absolute;
        content: '';
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        border-radius: 50%;
        border-bottom: 3px solid #1890ff;

        @for $i from 1 through 3 {
            &:nth-child(#{$i}) {
                animation: rotate#{$i} 1.15s linear infinite;
            }
        }

        &:nth-child(1) {
            animation-delay: -0.8s;
        }

        &:nth-child(2) {
            animation-delay: -0.4s;
        }

        &:nth-child(3) {
            animation-delay: 0s;
        }
    }
}

@keyframes rotate1 {
    from {
        transform: rotateX(35deg) rotateY(-45deg) rotateZ(0);
    }

    to {
        transform: rotateX(35deg) rotateY(-45deg) rotateZ(1turn);
    }
}

@keyframes rotate2 {
    from {
        transform: rotateX(50deg) rotateY(10deg) rotateZ(0);
    }

    to {
        transform: rotateX(50deg) rotateY(10deg) rotateZ(1turn);
    }
}

@keyframes rotate3 {
    from {
        transform: rotateX(35deg) rotateY(55deg) rotateZ(0);
    }

    to {
        transform: rotateX(35deg) rotateY(55deg) rotateZ(1turn);
    }
}
</style>