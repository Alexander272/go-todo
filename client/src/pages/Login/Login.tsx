import React, { useState } from "react"
import classes from "./login.module.scss"

export const LoginPage = () => {
    const [active, setActive] = useState("in")
    const [login, setLogin] = useState({
        email: "",
        password: "",
    })

    const onChangeTab = (event: React.MouseEvent<HTMLSpanElement>) => {
        const { textContent } = event.target as HTMLSpanElement
        if (textContent === "Sign In") {
            setActive("in")
        }
        if (textContent === "Sign Up") {
            setActive("up")
        }
    }

    return (
        <div className={classes.page}>
            <div className={classes.container}>
                <form action='' className={classes.form}>
                    <ul className={classes.nav}>
                        <li
                            className={`${classes.nav__item} ${
                                active === "in" ? classes.active : null
                            }`}
                        >
                            <span onClick={onChangeTab}>Sign In</span>
                        </li>
                        <li
                            className={`${classes.nav__item} ${
                                active === "up" ? classes.active : null
                            }`}
                        >
                            <span onClick={onChangeTab}>Sign Up</span>
                        </li>
                    </ul>
                    <div className={classes.tabs}>
                        <div
                            className={[
                                classes.tab,
                                active === "up" ? classes.tab__hidden : null,
                            ].join(" ")}
                        >
                            <label htmlFor='email' className={classes.label}>
                                Email
                            </label>
                            <input id='email' className={classes.input} type='email' />
                            <label htmlFor='password' className={classes.label}>
                                Пароль
                            </label>
                            <input id='password' className={classes.input} type='password' />
                            <button className={classes.submit} disabled>
                                Sign in
                            </button>
                        </div>
                        <div
                            className={[
                                classes.tab,
                                active === "in" ? classes.tab__hidden : null,
                            ].join(" ")}
                        >
                            <label htmlFor='name' className={classes.label}>
                                Логин
                            </label>
                            <input id='name' className={classes.input} type='text' />
                            <label htmlFor='email' className={classes.label}>
                                Email
                            </label>
                            <input id='email' className={classes.input} type='email' />
                            <label htmlFor='password' className={classes.label}>
                                Пароль
                            </label>
                            <input id='password' className={classes.input} type='password' />
                            <button className={classes.submit} disabled>
                                Sign up
                            </button>
                        </div>
                    </div>
                    {/* <label htmlFor='email' className={classes.label}>
                        Email
                    </label>
                    <input id='email' className={classes.input} type='email' />
                    <label htmlFor='password' className={classes.label}>
                        Пароль
                    </label>
                    <input id='password' className={classes.input} type='password' />
                    <button className={classes.submit} disabled>
                        Sign in
                    </button> */}
                </form>
                {/* <a href='#' className='classes.forgot'>
                    Forgot Password?
                </a> */}
            </div>
        </div>
    )
}
