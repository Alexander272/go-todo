import { FC } from "react"
import { Category } from "../../types/category"
import classes from "./item.module.scss"

type Props = {
    category: Category
    active: boolean
    onClick: () => void
}

export const CategoryItem: FC<Props> = ({ category, active, onClick }) => {
    return (
        <div className={`${classes.category} ${active && classes.active}`} onClick={onClick}>
            <p className={classes.title}>
                <span className={classes.icon}>&#8827;</span>
                {category.title}
            </p>
            {/* btn */}
        </div>
    )
}
