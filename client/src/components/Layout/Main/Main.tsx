import { FC } from 'react'
// import { Outlet } from 'react-router-dom'
import { Profile } from '../../Profile/Profile'
import classes from './main.module.scss'

export const MainLayout: FC = () => {
    return (
        <div className={classes.wrapper}>
            <Profile />
            <div className={classes.container}>
                {/* Search */}
                {/* <Outlet /> */}
            </div>
        </div>
    )
}
