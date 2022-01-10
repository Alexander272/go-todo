import { FC, PropsWithChildren } from "react"
import classes from "./select.module.scss"

type Props = {
    value: string
    disabled?: boolean
    onClick?: () => void
}

export const Option: FC<PropsWithChildren<Props>> = ({ children, disabled, onClick }) => {
    return (
        <p className={`${classes.option} ${disabled ? classes.disabled : ""}`} onClick={onClick}>
            {children}
        </p>
    )
}
