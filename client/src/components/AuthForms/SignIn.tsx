import { useForm, SubmitHandler } from "react-hook-form"
import { useDispatch } from "react-redux"
import { Dispatch } from "../../store/store"
import { ISignIn } from "../../types/user"
import { Button } from "../UI/Button/Button"
import { Input } from "../UI/Input/Input"
import classes from "./form.module.scss"

export const SignInForm = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<ISignIn>()

    const { user } = useDispatch<Dispatch>()

    const signInHandler: SubmitHandler<ISignIn> = data => {
        user.signIn(data)
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
                error={errors.email}
                errorText='Email is incorrect'
                rule={{
                    required: true,
                    pattern:
                        /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
                }}
            />
            <Input
                key='in-password'
                id='in-password'
                label='Password'
                name='password'
                inputType='round'
                type='password'
                register={register}
                error={errors.password}
                errorText='The minimum password length is 6 characters, the maximum is 64'
                rule={{ required: true, minLength: 6, maxLength: 64 }}
            />
            <Button rounded='round' type='submit'>
                Sign in
            </Button>
        </form>
    )
}
