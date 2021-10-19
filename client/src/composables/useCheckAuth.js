import { useStore } from 'vuex'
import { computed } from '@vue/reactivity'
import { useRouter } from 'vue-router'
import { watch } from '@vue/runtime-core'

export default function useCheckAuth() {
    const router = useRouter()
    const store = useStore()
    const loading = computed(() => store.state.loading)
    const ready = computed(() => store.state.ready)

    watch(ready, newValue => {
        if (newValue && !store.getters['auth/isAuth']) {
            router.push('/auth')
        }
    })

    return { loading, ready }
}
