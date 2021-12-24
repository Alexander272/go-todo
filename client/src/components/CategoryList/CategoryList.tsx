import { useSelector } from "react-redux"
import { RootState } from "../../store/store"
import { CategoryItem } from "../CategoryItem/CategoryItem"
import classes from "./list.module.scss"

export const CategoryList = () => {
    const categories = useSelector((state: RootState) => state.category.categories)

    // const mockCategories: CategoryWithLists[] = [
    //     { id: "123", title: "first category", lists: [] },
    //     { id: "234", title: "second category", lists: [] },
    //     { id: "345", title: "third category", lists: [] },
    // ]

    return (
        <>
            {categories.length > 0 ? (
                categories.map(cat => <CategoryItem key={cat.id} category={cat} active={false} />)
            ) : (
                <p className={classes.empty}>No groups have been created yet</p>
            )}
        </>
    )
}
