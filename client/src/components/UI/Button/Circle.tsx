import { FC } from "react"
import classes from "./button.module.scss"

type Props = {
    children?: React.ReactNode
    onClick?: () => void
    variant?: "primary" | "danger" | "ghost" | "grayPrimary" | "grayDanger"
    size?: "small" | "middle" | "large"
}

export const Circle: FC<Props & React.ButtonHTMLAttributes<any>> = ({
    children,
    onClick,
    variant,
    size,
    ...attr
}) => {
    return (
        <button
            onClick={onClick}
            className={[
                classes.button,
                classes.circle,
                classes[variant || "primary"],
                classes[size || "middle"],
            ].join(" ")}
            {...attr}
        >
            {children}
        </button>
    )
}
