import { FC, useState } from "react"
import { List } from "../../types/list"
import { ListScrollItem } from "../ListsScrollItem/ListsScrollItem"
import classes from "./scroll.module.scss"

type Props = {
    title: string
    data: List[]
}

export const ListsScroll: FC<Props> = ({ title, data }) => {
    const [isOpen, setIsOpen] = useState(false)

    const toggle = () => {
        setIsOpen(prev => !prev)
    }

    return (
        <div className={`${classes.scroll} ${isOpen && classes.active}`}>
            <div className={`${classes.header} ${classes.item}`} onClick={toggle}>
                <p className={classes.title}>
                    {title}
                    <span className={classes.icon}>&#8910;</span>
                </p>
            </div>
            <div className={classes.content}>
                {isOpen && data.map(d => <ListScrollItem item={d} className={classes.item} />)}
            </div>
        </div>
    )
}
