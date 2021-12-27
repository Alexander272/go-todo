import { FC } from "react"
import { List } from "../../types/list"
import { Checkbox } from "../UI/Checkbox/Checkbox"
import classes from "./item.module.scss"

type Props = {
    item: List
    className?: string
}

const options = {
    day: "numeric" as "numeric",
    month: "short" as "short",
    year: "numeric" as "numeric",
}

export const ListScrollItem: FC<Props> = ({ item, className }) => {
    const date = new Date(item.createdAt * 1000).toLocaleString("en-US", options)

    return (
        <div className={`${classes.item} ${className}`}>
            <Checkbox id='cbx' checked={false} />
            <div className={classes.content}>
                <p className={classes.title}>{item.title}</p>
                <p className={classes.date}>{date}</p>
            </div>
        </div>
    )
}
