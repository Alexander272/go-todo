import { useSelector } from "react-redux"
import { RootState } from "../../store/store"
import classes from "./tasks.module.scss"

const options = {
    day: "numeric" as "numeric",
    month: "short" as "short",
    year: "numeric" as "numeric",
}

export default function Tasks() {
    const selectList = useSelector((state: RootState) => state.todo.selectList)

    return (
        <div className={`${classes.tasks} scroll`}>
            {!selectList ? (
                <p className={classes.empty}>No list selected</p>
            ) : (
                <div className={classes.container}>
                    <h3 className={classes.title}>{selectList.title}</h3>
                    <div className={classes.dateContainer}>
                        <svg
                            xmlns='http://www.w3.org/2000/svg'
                            viewBox='0 0 24 24'
                            fill='none'
                            stroke='currentColor'
                            strokeWidth='2'
                            strokeLinecap='round'
                            strokeLinejoin='round'
                        >
                            <circle cx='12' cy='12' r='10'></circle>
                            <path d='M12 6v6l4 2'></path>
                        </svg>
                        <p className={classes.date}>
                            {new Date(selectList?.createdAt * 1000).toLocaleString(
                                "en-US",
                                options
                            )}
                        </p>
                    </div>

                    <p className={classes.description}>{selectList.description}</p>
                </div>
            )}
        </div>
    )
}
