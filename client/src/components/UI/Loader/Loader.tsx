import { FC } from "react"
import classes from "./loader.module.scss"

type Props = {
    size?: "small" | "middle" | "large"
    background?: "none" | "fill"
}

export const Loader: FC<Props> = ({ size, background }) => {
    return (
        <div className={`${classes.container} ${classes[background || "fill"]}`}>
            <div className={`${classes.loader} ${classes[size || "middle"]}`}></div>
        </div>
    )
}
