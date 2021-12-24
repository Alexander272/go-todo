import { FC } from "react"
import classes from "./backdrop.module.scss"

type Props = {
    onClick: () => void
}

export const Backdrop: FC<Props> = ({ onClick }) => {
    return <div className={classes.backdrop} onClick={onClick} />
}
