import { useSelector } from "react-redux"
import { ListsScroll } from "../../components/ListsScroll/ListsScroll"
import { Button } from "../../components/UI/Button/Button"
import { RootState } from "../../store/store"
import Tasks from "../Tasks/Tasks"
import classes from "./lists.module.scss"

export default function ListsPage() {
    const selCat = useSelector((state: RootState) => state.category.selectedCategories)
    const categories = useSelector((state: RootState) => state.category.categories)

    return (
        <div className={classes.page}>
            <div className={`${classes.container} scroll`}>
                {selCat.length > 0 ? (
                    categories
                        .filter(cat => selCat.includes(cat.id))
                        .map((cat, index) => (
                            <ListsScroll
                                key={cat.id}
                                title={cat.title}
                                data={cat.lists}
                                open={index === 0}
                            />
                        ))
                ) : (
                    <p className={classes.empty}>No group selected</p>
                )}
                <div className={classes.add}>
                    <Button size='small' rounded='round'>
                        Add list
                    </Button>
                </div>
            </div>
            <Tasks />
        </div>
    )
}
