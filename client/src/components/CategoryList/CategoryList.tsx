import { useDispatch, useSelector } from "react-redux"
import { Dispatch, RootState } from "../../store/store"
import { CategoryItem } from "../CategoryItem/CategoryItem"
import classes from "./list.module.scss"

export const CategoryList = () => {
    const categories = useSelector((state: RootState) => state.category.categories)
    const selCat = useSelector((state: RootState) => state.category.selectedCategories)

    const { category } = useDispatch<Dispatch>()

    const selectHandler = (id: string) => () => {
        category.selectCategory(id)
    }
    const unselectHandler = (id: string) => () => {
        category.unselectCategory(id)
    }

    return (
        <>
            {categories.length > 0 ? (
                categories.map(cat => (
                    <CategoryItem
                        key={cat.id}
                        category={cat}
                        active={selCat.includes(cat.id)}
                        onClick={
                            selCat.includes(cat.id)
                                ? unselectHandler(cat.id)
                                : selectHandler(cat.id)
                        }
                    />
                ))
            ) : (
                <p className={classes.empty}>No groups have been created yet</p>
            )}
        </>
    )
}
