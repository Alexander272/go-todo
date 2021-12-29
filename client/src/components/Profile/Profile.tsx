import { FC } from "react"
import { useSelector } from "react-redux"
import { RootState, store } from "../../store/store"
import { CategoryList } from "../CategoryList/CategoryList"
import classes from "./profile.module.scss"

export const Profile: FC = () => {
    const nicname = useSelector((state: RootState) => state.user.nickname)

    const total = useSelector(store.select.category.total)
    const { completed, count } = useSelector(store.select.category.completed)

    return (
        <div className={`${classes.profile} scroll`}>
            <h3 className={classes.appName}>Task Manager</h3>
            <div className={classes.wrapper}>
                <div className={classes.profileBlock}>
                    <img
                        className={classes.image}
                        src='https://assets.codepen.io/3364143/Screen+Shot+2020-08-01+at+12.24.16.png'
                        alt='user'
                    />
                    <p className={classes.username}>{nicname || "User Name"}</p>
                </div>
                <div className={classes.progress}>
                    <p className={classes.progress__count}>
                        {completed}/{count}
                    </p>
                    <div className={classes.progress__bar}>
                        <div
                            style={{ width: `${(completed / (count || 1)) * 100}%` }}
                            className={classes.progress__line}
                        ></div>
                    </div>
                </div>
                <div className={classes.status}>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>{completed}</p>
                        <p className={classes.status__text}>Completed</p>
                        <p className={classes.status__tasks}>tasks</p>
                    </div>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>{count - completed}</p>
                        <p className={classes.status__text}>To do</p>
                        <p className={classes.status__tasks}>tasks</p>
                    </div>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>{total}</p>
                        <p className={classes.status__text}>All</p>
                        <p className={classes.status__tasks}>completed</p>
                    </div>
                </div>
            </div>

            <div className={classes.wrapper}>
                <CategoryList />
            </div>
        </div>
    )
}
