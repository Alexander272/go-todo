import { useState } from "react"
import { SignInForm } from "../../components/AuthForms/SignIn"
import { SingUpForm } from "../../components/AuthForms/SignUp"
import classes from "./auth.module.scss"

export default function AuthPage() {
    const [isSignIn, setIsSignIn] = useState(true)

    const changeTabHandler = (value: boolean) => () => {
        setIsSignIn(value)
    }

    return (
        <div className={classes.page}>
            <div className={`${classes.container} ${isSignIn ? classes.signIn : classes.signUp}`}>
                <div className={classes.form}>
                    <ul className={classes.nav}>
                        <li className={`${classes.link} ${isSignIn && classes.active}`}>
                            <span onClick={changeTabHandler(true)}>Sign In</span>
                        </li>
                        <li className={`${classes.link} ${!isSignIn && classes.active}`}>
                            <span onClick={changeTabHandler(false)}>Sign Up</span>
                        </li>
                    </ul>
                    {isSignIn ? <SignInForm /> : <SingUpForm />}
                </div>
            </div>
        </div>
    )
}
