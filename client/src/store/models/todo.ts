import { createModel } from "@rematch/core"
import { toast } from "react-toastify"
import TodoService from "../../service/todo"
import { Todo } from "../../types/todo"
import { RootModel } from "./index"

export interface ITodo {}

type TodoState = {
    loading: boolean
    todos: Todo[]
    completed: number
    remaining: number
}

export const todo = createModel<RootModel>()({
    state: {
        loading: false,
        todos: [],
        completed: 12,
        remaining: 34,
    } as TodoState,

    //? А стоит ли хранить это здесь

    reducers: {
        setLoading(state, payload: boolean) {
            state.loading = payload
            return state
        },
        setTodos(state, payload: Todo[]) {
            state.todos = payload
            return state
        },
        add(state, payload: Todo) {
            state.todos.push(payload)
            return state
        },
        update(state, payload: Todo) {
            state.todos = state.todos.map(todo => {
                if ((todo.id = payload.id)) todo = payload
                return todo
            })
            return state
        },
    },

    effects: dispatch => {
        const { todo } = dispatch

        return {
            async getTodos(listId: string) {
                todo.setLoading(true)
                try {
                    const res = await TodoService.get(listId)
                    todo.setTodos(res.data)
                } catch (error: any) {
                    toast.error(error.message)
                } finally {
                    todo.setLoading(false)
                }
            },
        }
    },
})
