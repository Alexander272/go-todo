import { FC } from 'react'
import { useSelector } from 'react-redux'
import { RootState } from '../../store/store'
import classes from './profile.module.scss'

export const Profile: FC = () => {
    const completed = useSelector((state: RootState) => state.todo.completed)
    const remaining = useSelector((state: RootState) => state.todo.remaining)

    return (
        <div className={classes.profile}>
            <h3 className={classes.appName}>Task Manager</h3>
            <div className={classes.wrapper}>
                <div className={classes.profileBlock}>
                    <img
                        className={classes.image}
                        src='https://assets.codepen.io/3364143/Screen+Shot+2020-08-01+at+12.24.16.png'
                        alt='user'
                    />
                    <p className={classes.username}>User Name</p>
                </div>
                <div className={classes.progress}>
                    <p className={classes.progress__count}>
                        {completed}/{remaining}
                    </p>
                    <div className={classes.progress__bar}>
                        <div
                            style={{ width: `${(completed / remaining) * 100}%` }}
                            className={classes.progress__line}></div>
                    </div>
                </div>
                <div className={classes.status}>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>12</p>
                        <p className={classes.status__text}>Completed</p>
                        <p className={classes.status__tasks}>tasks</p>
                    </div>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>22</p>
                        <p className={classes.status__text}>To do</p>
                        <p className={classes.status__tasks}>tasks</p>
                    </div>
                    <div className={classes.status__item}>
                        <p className={classes.status__count}>243</p>
                        <p className={classes.status__text}>All</p>
                        <p className={classes.status__tasks}>completed</p>
                    </div>
                </div>
            </div>

            <div className={classes.projects}></div>
        </div>
    )
}
