import { useEffect, useState } from "react"
import { useSelector } from "react-redux"
import { useLocation, useNavigate } from "react-router-dom"
import { SignInForm } from "../../components/AuthForms/SignIn"
import { SingUpForm } from "../../components/AuthForms/SignUp"
import { Loader } from "../../components/UI/Loader/Loader"
import { RootState } from "../../store/store"
import classes from "./auth.module.scss"

export default function AuthPage() {
    const [isSignIn, setIsSignIn] = useState(true)

    const navigate = useNavigate()
    const location = useLocation()

    const loading = useSelector((state: RootState) => state.user.loading)
    const token = useSelector((state: RootState) => state.user.token.accessToken)

    const from: string = (location.state as any)?.from?.pathname || "/"

    useEffect(() => {
        if (token !== "") {
            navigate(from, { replace: true })
        }
    }, [token, navigate, from])

    const changeTabHandler = (value: boolean) => () => {
        setIsSignIn(value)
    }

    return (
        <div className={classes.page}>
            <div className={`${classes.container} ${isSignIn ? classes.signIn : classes.signUp}`}>
                {loading && <Loader />}
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
