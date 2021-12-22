import { FC } from "react"
import classes from "./button.module.scss"

type Props = {
    children?: React.ReactNode
    onClick?: () => void
    variant?: "primary" | "danger" | "ghost" | "grayPrimary" | "grayDanger"
    size?: "small" | "middle" | "large"
    rounded?: "none" | "low" | "medium" | "high" | "round" | "circle"
}

export const Button: FC<Props & React.ButtonHTMLAttributes<any>> = ({
    children,
    onClick,
    variant,
    size,
    rounded,
    ...attr
}) => {
    return (
        <button
            onClick={onClick}
            className={[
                classes.button,
                classes[variant || "primary"],
                classes[size || "middle"],
                classes[rounded || "medium"],
            ].join(" ")}
            {...attr}
        >
            {children}
        </button>
    )
}
