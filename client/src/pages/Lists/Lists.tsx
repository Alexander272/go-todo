import { useSelector } from "react-redux"
import { ListsScroll } from "../../components/ListsScroll/ListsScroll"
import { RootState } from "../../store/store"
import classes from "./lists.module.scss"

export default function ListsPage() {
    const selCat = useSelector((state: RootState) => state.category.selectedCategories)
    const categories = useSelector((state: RootState) => state.category.categories)

    return (
        <div className={classes.page}>
            <div className={classes.container}>
                {selCat.length > 0 ? (
                    categories
                        .filter(cat => selCat.includes(cat.id))
                        .map(cat => <ListsScroll key={cat.id} title={cat.title} data={cat.lists} />)
                ) : (
                    <p className={classes.empty}>No lists selected</p>
                )}
            </div>
        </div>
    )
}
