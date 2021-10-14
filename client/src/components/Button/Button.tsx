import React, { FC } from "react"
import { Link } from "react-router-dom"
import classes from "./button.module.scss"

type Props = {
    text: string
    className?: string
    colorClass?: string
    testStyle?: string
    path?: string
    onClick?: (event: React.MouseEvent) => void
}

export const Button: FC<Props> = ({ text, className, onClick }) => {
    return (
        <button className={[classes.button, className].join(" ")} onClick={onClick}>
            {text}
        </button>
    )
}

export const LinkButton: FC<Props> = ({ text, path, className, colorClass, testStyle }) => {
    return (
        <div className={[classes.block, className].join(" ")}>
            <div className={[classes.dropdown, colorClass].join(" ")}></div>
            <Link className={[classes.link, testStyle].join(" ")} to={path || "#"}>
                {text}
            </Link>
        </div>
    )
}
