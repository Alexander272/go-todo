import { FC } from "react"
import classes from "./loader.module.scss"

type Props = {
    size?: "small" | "middle" | "large"
}

export const Loader: FC<Props> = ({ size }) => {
    return (
        <div className={classes.container}>
            <div className={`${classes.loader} ${classes[size || "middle"]}`}></div>
        </div>
    )
}
