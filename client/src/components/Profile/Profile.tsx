import React, { FC, useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { Dispatch, RootState, store } from "../../store/store"
import { CategoryList } from "../CategoryList/CategoryList"
import { useModal } from "../Modal/hooks/useModal"
import { Modal } from "../Modal/Modal"
import { Button } from "../UI/Button/Button"
import { Input } from "../UI/Input/Input"
import classes from "./profile.module.scss"

export const Profile: FC = () => {
    const nicname = useSelector((state: RootState) => state.user.nickname)

    const total = useSelector(store.select.category.total)
    const { completed, count } = useSelector(store.select.category.completed)
    const { category } = useDispatch<Dispatch>()

    const [error, setError] = useState(false)
    const [title, setTitle] = useState("")

    const { isOpen, toggle } = useModal()

    const changeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTitle(event.target.value)
    }

    const saveHandler = () => {
        if (!title.trim()) {
            setError(true)
            return
        }
        setError(false)
        category.newCategory({ title })
    }

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
                <div className={classes.projects}>
                    <div className={classes.projects__header}>
                        <p className={classes.projects__title}>Groups</p>
                        <Button.Circle onClick={toggle} size='small' variant='grayPrimary'>
                            +
                        </Button.Circle>
                    </div>
                    <CategoryList />
                </div>
            </div>

            {isOpen && (
                <Modal isOpen={isOpen} toggle={toggle}>
                    <Modal.Header title='Create group' onClose={toggle} />
                    <Modal.Content>
                        <Input
                            name='title'
                            placeholder='title'
                            value={title}
                            error={error}
                            errorText={"Title is required"}
                            onChange={changeHandler}
                        />
                    </Modal.Content>
                    <Modal.Footer>
                        <div className={classes.btns}>
                            <Button size='small' onClick={toggle} variant='grayPrimary'>
                                Cancel
                            </Button>
                            <Button size='small' onClick={saveHandler}>
                                Create
                            </Button>
                        </div>
                    </Modal.Footer>
                </Modal>
            )}
        </div>
    )
}
