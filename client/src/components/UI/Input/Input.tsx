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
    rule?: Partial<any>
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
    rule,
    onChange,
    error,
    errorText,
    ...attr
}: Props & React.InputHTMLAttributes<HTMLInputElement>) => {
    // const partial = rule ? rule : []

    return (
        <div className={`${classes.field} ${classes[orentation || "vertical"]}`}>
            {label && (
                <label className={classes.label} htmlFor={id}>
                    {label}
                </label>
            )}
            {register ? (
                <input
                    className={`${classes.input} ${error && classes.invalid} ${
                        classes[inputType || "rounded"]
                    }`}
                    id={id}
                    {...attr}
                    {...register(name, rule)}
                />
            ) : (
                <input
                    className={`${classes.input} ${error && classes.invalid} ${
                        classes[inputType || "rounded"]
                    }`}
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
