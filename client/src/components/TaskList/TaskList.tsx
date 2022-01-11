import { FC, useCallback, useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { toast } from "react-toastify"
import TodoService from "../../service/todo"
import { RootState } from "../../store/store"
import { Todo, UpdateTodo } from "../../types/todo"
import { TaskItem } from "../TaskItem/TaskItem"
import { Loader } from "../UI/Loader/Loader"
import classes from "./taskList.module.scss"

export const TaskList: FC = () => {
    const selectList = useSelector((state: RootState) => state.todo.selectList)

    const [isLoading, setIsLoading] = useState(true)
    const [list, setList] = useState<Todo[]>([
        {
            id: "1",
            listId: "",
            title: "mock title",
            description: "some description",
            createdAt: 0,
            completedAt: 0,
            startAt: 1641882614,
            done: true,
            priority: 1,
        },
        {
            id: "2",
            listId: "",
            title: "mock title 2",
            description: "some description",
            createdAt: 0,
            completedAt: 0,
            startAt: 1649882614,
            done: false,
            priority: 1,
        },
    ])

    const getTask = useCallback(async () => {
        if (selectList) {
            try {
                const res = await TodoService.get(selectList?.id)
                setList(res.data)
            } catch (error: any) {
                toast.error(error.message)
            }
        }
    }, [selectList])

    useEffect(() => {
        setIsLoading(true)
        getTask()
        setIsLoading(false)
    }, [getTask])

    // const updateTask = async (task: UpdateTodo) => {
    //     try {
    //         await TodoService.update(task)
    //     } catch (error: any) {
    //         toast.error(error.message)
    //     }
    // }

    if (isLoading)
        return (
            <div className={classes.list}>
                <Loader background='none' />
            </div>
        )

    return (
        <div className={classes.list}>
            {list.map(task => (
                <TaskItem key={task.id} task={task} />
            ))}
        </div>
    )
}
