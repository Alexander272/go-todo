import { useForm, SubmitHandler } from "react-hook-form"
import { Button } from "../UI/Button/Button"
import { Input } from "../UI/Input/Input"
import classes from "./form.module.scss"

export interface IFormSignIn {
    email: string
    password: string
}

export const SignInForm = () => {
    const { register, handleSubmit } = useForm<IFormSignIn>()

    const signInHandler: SubmitHandler<IFormSignIn> = data => {
        console.log("signInHandler", data)
    }

    return (
        <form className={classes.tab} onSubmit={handleSubmit(signInHandler)}>
            <Input
                key='in-email'
                id='in-email'
                label='Email'
                name='email'
                inputType='round'
                type='email'
                register={register}
            />
            <Input
                key='in-password'
                id='in-password'
                label='Password'
                name='password'
                inputType='round'
                type='password'
                register={register}
            />
            <Button rounded='round' type='submit'>
                Sign in
            </Button>
        </form>
    )
}
