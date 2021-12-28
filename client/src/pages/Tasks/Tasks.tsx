import { useSelector } from "react-redux"
import { RootState } from "../../store/store"
import classes from "./tasks.module.scss"

export default function Tasks() {
    const selectList = useSelector((state: RootState) => state.todo.selectList)

    return (
        <div className={`${classes.tasks} scroll`}>
            {!selectList ? <p className={classes.empty}>No list selected</p> : null}
        </div>
    )
}
