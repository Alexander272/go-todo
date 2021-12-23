import { useForm, SubmitHandler } from "react-hook-form"
import { useDispatch } from "react-redux"
import { Button } from "../../components/UI/Button/Button"
import { Input } from "../../components/UI/Input/Input"
import { Dispatch } from "../../store/store"
import { ISignUp } from "../../types/user"
import classes from "./form.module.scss"

export const SingUpForm = () => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<ISignUp>()

    const { user } = useDispatch<Dispatch>()

    const signUpHandler: SubmitHandler<ISignUp> = data => {
        user.singUp(data)
    }

    return (
        <form className={classes.tab} onSubmit={handleSubmit(signUpHandler)}>
            <Input
                key='up-nickname'
                id='up-nickname'
                label='Nickname'
                name='nickname'
                inputType='round'
                register={register}
                error={errors.nickname}
                errorText='Nicname is required'
                rule={{ required: true }}
            />
            <Input
                key='up-email'
                id='up-email'
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
                key='up-password'
                id='up-password'
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
                Sign up
            </Button>
        </form>
    )
}
