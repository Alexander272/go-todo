<template>
    <div class="page">
        <div class="container" :class="{ 'sign-in': isSignIn, 'sign-up': !isSignIn }">
            <div class="form">
                <ul class="nav">
                    <li class="nav__item" :class="{ active: isSignIn }">
                        <span @click="changeTab('sign-in')">Sign In</span>
                    </li>
                    <li class="nav__item" :class="{ active: !isSignIn }">
                        <span @click="changeTab('sign-up')">Sign Up</span>
                    </li>
                </ul>
                <transition name="slide-fade" mode="out-in">
                    <form class="tab" v-if="isSignIn" @submit.prevent="SignIn(signIn)">
                        <input-field
                            id="email"
                            name="email"
                            type="email"
                            labelText="Email"
                            :errorText="signIn.errorEmail"
                            v-model="signIn.email"
                        />
                        <input-field
                            id="password"
                            name="password"
                            type="password"
                            labelText="Пароль"
                            :errorText="signIn.errorPassword"
                            v-model="signIn.password"
                        />
                        <button class="submit" @click="checkFormSignIn">Sign in</button>
                    </form>
                    <form class="tab" v-else @submit.prevent="SignUp(signUp)">
                        <input-field
                            id="name"
                            name="name"
                            labelText="Логин"
                            v-model="signUp.login"
                        />
                        <input-field
                            id="email"
                            name="email"
                            type="email"
                            labelText="Email"
                            v-model="signUp.email"
                        />
                        <input-field
                            id="password"
                            name="password"
                            type="password"
                            labelText="Пароль"
                            v-model="signUp.password"
                        />
                        <button class="submit" @click="checkformSignUp">Sign up</button>
                    </form>
                </transition>
            </div>
            <!-- <a href='#' class='forgot'>
                    Forgot Password?
                </a>  -->
        </div>
    </div>
</template>

<script>
import { ref, computed, reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import InputField from '../components/InputField.vue'

export default {
    components: { InputField },
    name: 'Auth',
    setup() {
        const activeTab = ref('sign-in')
        const isSignIn = computed(() => activeTab.value === 'sign-in')

        const changeTab = value => {
            activeTab.value = value
        }

        const signIn = reactive({
            email: '',
            errorEmail: '',
            password: '',
            errorPassword: '',
        })
        const signUp = reactive({
            login: '',
            errorLogin: '',
            email: '',
            errorEmail: '',
            password: '',
            errorPassword: '',
        })

        const checkFormSignIn = () => {
            if (!validEmail(signIn.email)) signIn.errorEmail = 'Email некорректен'
            else signIn.errorEmail = ''

            if (signIn.password.trim().length === 0)
                signIn.errorPassword = 'Пароль не должен быть пустым'
            else signIn.errorPassword = ''
        }
        const checkformSignUp = () => {
            if (signUp.login.trim().length < 2 || signUp.login.trim().length > 64)
                signUp.errorLogin = 'Длина логина должна составлять от 2 до 64 символов'
            else signUp.errorLogin = ''
            if (!validEmail(signUp.email)) signUp.errorEmail = 'Email некорректен'
            else signUp.errorEmail = ''
            if (signUp.password.trim().length < 8 || signUp.password.trim().length > 64)
                signUp.errorPassword = 'Длина пароля должна составлять от 8 до 64 символов'
            else signUp.errorPassword = ''
        }

        const validEmail = email => {
            var re =
                /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            return re.test(email)
        }

        const store = useStore()
        const router = useRouter()

        // {email: signIn.email, password: signIn.password}
        const SignIn = signIn => {
            if (signIn.errorEmail || signIn.errorPassword) return
            store.dispatch('auth/signIn', signIn).then(() => router.push('/'))
        }
        const SignUp = signUp => {
            if (signUp.errorEmail || signUp.errorPassword || signUp.errorLogin) return
            store.dispatch('auth/signUp', signUp)
        }

        return {
            activeTab,
            isSignIn,
            changeTab,
            signIn,
            signUp,
            checkFormSignIn,
            checkformSignUp,
            SignIn,
            SignUp,
        }
    },
}
</script>

<style lang="scss" scoped>
.page {
    background-color: #493968;
    width: 100%;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
}

.container {
    display: block;
    position: relative;
    z-index: 0;
    margin: auto;
    padding: 5rem 4rem 4rem 4rem;
    width: 100%;
    max-width: 525px;
    min-height: 580px;
    background-image: url(https://s3-us-west-2.amazonaws.com/s.cdpn.io/283591/login-background.jpg);
    box-shadow: 0 50px 70px -20px rgba(0, 0, 0, 0.85);
    border-radius: 16px;
    overflow: hidden;
    transition: 0.5s min-height ease-in-out;

    &:after {
        content: '';
        display: inline-block;
        position: absolute;
        z-index: 0;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        background-image: radial-gradient(
            ellipse at left bottom,
            rgba(22, 24, 47, 1) 0%,
            rgba(38, 20, 72, 0.9) 59%,
            rgba(17, 27, 75, 0.9) 100%
        );
        box-shadow: 0 -20px 150px -20px rgba(0, 0, 0, 0.5);
    }
}
.sign-in {
    min-height: 540px;
}
.sign-up {
    min-height: 635px;
}

.form {
    position: relative;
    z-index: 1;
    // padding-bottom: 3.5rem;
    // border-bottom: 1px solid rgba(255, 255, 255, 0.25);
}
.tab {
    padding-bottom: 3.5rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.25);
}

.nav {
    position: relative;
    padding: 0;
    margin: 0 0 3em 1rem;

    &__item {
        list-style: none;
        display: inline-block;
    }

    &__item + &__item {
        margin-left: 2.25rem;
    }

    &__item span {
        position: relative;
        color: rgba(255, 255, 255, 0.5);
        text-decoration: none;
        text-transform: uppercase;
        font-weight: 500;
        font-size: 1.25rem;
        padding-bottom: 0.5rem;
        transition: 0.2s all ease;
        cursor: pointer;

        &:after {
            content: '';
            display: inline-block;
            height: 10px;
            background-color: rgb(255, 255, 255);
            position: absolute;
            right: 100%;
            bottom: -1px;
            left: 0;
            border-radius: 50%;
            transition: 0.15s all ease;
        }
    }

    &__item.active span,
    &__item span:hover {
        color: #ffffff;
        transition: 0.15s all ease;
    }

    &__item span:hover:after,
    &__item.active span:after {
        background-color: rgb(17, 97, 237);
        height: 2px;
        right: 0;
        bottom: 2px;
        border-radius: 0;
        transition: 0.2s all ease;
    }
}

.submit {
    color: #ffffff;
    font-size: 1rem;
    font-family: 'Montserrat', sans-serif;
    text-transform: uppercase;
    letter-spacing: 1px;
    margin-top: 3rem;
    padding: 0.75rem;
    border-radius: 2rem;
    display: block;
    width: 100%;
    background-color: rgba(17, 97, 237, 0.75);
    border: none;
    cursor: pointer;

    &:hover {
        background-color: rgba(17, 97, 237, 1);
    }
}

// .forgot {
//     display: block;
//     margin-top: 3rem;
//     text-align: center;
//     color: rgba(255, 255, 255, 0.75);
//     font-size: 0.75rem;
//     text-decoration: none;
//     position: relative;
//     z-index: 1;

//     &:hover {
//         color: rgb(17, 97, 237);
//     }
// }

.slide-fade-enter-active {
    transition: all 0.4s ease-out;
}

.slide-fade-leave-active {
    transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from {
    transform: translateX(-120px);
    opacity: 0;
}
.slide-fade-leave-to {
    transform: translateX(120px);
    opacity: 0;
}
</style>
