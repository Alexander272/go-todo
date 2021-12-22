import { UseFormRegister } from "react-hook-form"
import classes from "./input.module.scss"

type Props = {
    id?: string
    label?: string
    name: string
    orentation?: "horizontal" | "vertical"
    inputType?: "round" | "rounded"
    onChange?: any
    register?: UseFormRegister<any>
    error?: any
    errorText?: string
}

export const Input = ({
    id,
    label,
    name,
    orentation,
    inputType,
    register,
    onChange,
    error,
    errorText,
    ...attr
}: Props & React.InputHTMLAttributes<HTMLInputElement>) => {
    return (
        <div className={`${classes.field} ${classes[orentation || "vertical"]}`}>
            {label && (
                <label className={classes.label} htmlFor={id}>
                    {label}
                </label>
            )}
            {register ? (
                <input
                    className={`${classes.input} ${classes[inputType || "rounded"]}`}
                    id={id}
                    {...attr}
                    {...register(name)}
                />
            ) : (
                <input
                    className={`${classes.input} ${classes[inputType || "rounded"]}`}
                    id={id}
                    name={name}
                    onChange={onChange}
                    {...attr}
                />
            )}

            {error && <p className={classes.error}>{errorText}</p>}
        </div>
    )
}
