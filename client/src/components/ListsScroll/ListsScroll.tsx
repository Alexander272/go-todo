import { FC, useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { Dispatch, RootState } from "../../store/store"
import { List } from "../../types/list"
import { ListScrollItem } from "../ListsScrollItem/ListsScrollItem"
import classes from "./scroll.module.scss"

type Props = {
    title: string
    open?: boolean
    data: List[]
}

export const ListsScroll: FC<Props> = ({ title, data, open }) => {
    const [isOpen, setIsOpen] = useState(open || false)

    const selectlist = useSelector((state: RootState) => state.todo.selectList)
    const { todo } = useDispatch<Dispatch>()

    const toggle = () => {
        setIsOpen(prev => !prev)
    }

    const selectHandler = (list: List) => () => {
        todo.setSelectList(list)
    }

    return (
        <div className={`${classes.scroll} scroll ${isOpen ? classes.active : ""}`}>
            <div className={`${classes.header} ${classes.item}`} onClick={toggle}>
                <p className={classes.title}>
                    {title}
                    <span className={classes.icon}>&#8910;</span>
                </p>
            </div>
            <div className={classes.content}>
                {isOpen &&
                    data.map(d => (
                        <ListScrollItem
                            key={d.id}
                            item={d}
                            className={classes.item}
                            select={selectlist?.id === d.id}
                            onSelect={selectHandler(d)}
                        />
                    ))}
                {/* {data.map(d => (
                    <ListScrollItem item={d} className={classes.item} />
                ))} */}
            </div>
        </div>
    )
}
