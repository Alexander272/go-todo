import classes from "./button.module.scss"
import { Circle } from "./Circle"

type Props = {
    children?: React.ReactNode
    onClick?: () => void
    variant?: "primary" | "danger" | "ghost" | "grayPrimary" | "grayDanger"
    size?: "small" | "middle" | "large"
    rounded?: "none" | "low" | "medium" | "high" | "round"
}

const Button = ({
    children,
    onClick,
    variant,
    size,
    rounded,
    ...attr
}: Props & React.ButtonHTMLAttributes<any>) => {
    return (
        <button
            onClick={onClick}
            className={[
                classes.button,
                classes.default,
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

Button.Circle = Circle

export { Button }
