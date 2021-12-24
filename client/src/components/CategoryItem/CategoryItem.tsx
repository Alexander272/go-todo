import { FC } from "react"
import { Category } from "../../types/category"
import classes from "./item.module.scss"

type Props = {
    category: Category
    active: boolean
}

export const CategoryItem: FC<Props> = ({ category, active }) => {
    return (
        <div className={`${classes.category} ${active && classes.active}`}>
            <p className={classes.title}>
                <span className={classes.icon}>&#8827;</span>
                {category.title}
            </p>
            {/* btn */}
        </div>
    )
}
