import { useForm, SubmitHandler } from "react-hook-form"
import { Button } from "../../components/UI/Button/Button"
import { Input } from "../../components/UI/Input/Input"
import classes from "./form.module.scss"

export interface IFormSignUp {
    nicname: string
    email: string
    password: string
}

export const SingUpForm = () => {
    const { register, handleSubmit } = useForm<IFormSignUp>()

    const signUpHandler: SubmitHandler<IFormSignUp> = data => {
        console.log("signUpHandler", data)
    }

    return (
        <form className={classes.tab} onSubmit={handleSubmit(signUpHandler)}>
            <Input
                key='up-nicname'
                id='up-nicname'
                label='Nicname'
                name='nicname'
                inputType='round'
                register={register}
            />
            <Input
                key='up-email'
                id='up-email'
                label='Email'
                name='email'
                inputType='round'
                type='email'
                register={register}
            />
            <Input
                key='up-password'
                id='up-password'
                label='Password'
                name='password'
                inputType='round'
                type='password'
                register={register}
            />
            <Button rounded='round' type='submit'>
                Sign up
            </Button>
        </form>
    )
}
