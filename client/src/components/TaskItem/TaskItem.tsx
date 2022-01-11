import { FC } from "react"
import { Todo } from "../../types/todo"
import { Checkbox } from "../UI/Checkbox/Checkbox"
import classes from "./taskItem.module.scss"

type Props = {
    task: Todo
    onToggle?: () => void
}

const options = {
    day: "numeric" as "numeric",
    month: "short" as "short",
    year: "numeric" as "numeric",
    hour: "2-digit" as "2-digit",
    minute: "2-digit" as "2-digit",
}

export const TaskItem: FC<Props> = ({ task, onToggle }) => {
    return (
        <div className={classes.item}>
            <Checkbox id={task.id} checked={task.done} onClick={onToggle} />
            <div className={classes.info}>
                {/* //todo надо сюда запихать приоритет задачи */}
                <p className={classes.title}>{task.title}</p>
                <p className={classes.description}>{task.description}</p>
            </div>
            <p
                className={`${classes.date} ${
                    task.startAt * 1000 < Date.now() && !task.done ? classes.danger : ""
                }`}
            >
                {new Date(task.startAt * 1000).toLocaleString("en-US", options)}
            </p>
            {/* //todo можно выводить строчку выполнено тогда-то */}
        </div>
    )
}
