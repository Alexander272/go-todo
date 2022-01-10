import { useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { Dispatch, RootState } from "../../store/store"
import { Category, CategoryWithLists } from "../../types/category"
import { CategoryItem } from "../CategoryItem/CategoryItem"
import { useModal } from "../Modal/hooks/useModal"
import { Modal } from "../Modal/Modal"
import { Button } from "../UI/Button/Button"
import { Input } from "../UI/Input/Input"
import classes from "./list.module.scss"

export const CategoryList = () => {
    const categories = useSelector((state: RootState) => state.category.categories)
    const selCat = useSelector((state: RootState) => state.category.selectedCategories)

    const [error, setError] = useState(false)
    const [catState, setCatState] = useState<Category>({ id: "", title: "" })

    const { isOpen, toggle } = useModal()

    const { category } = useDispatch<Dispatch>()

    const selectHandler = (id: string) => () => {
        category.selectCategory(id)
    }
    const unselectHandler = (id: string) => () => {
        category.unselectCategory(id)
    }

    const changeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setCatState(prev => ({ ...prev, title: event.target.value }))
    }
    const cancelHandler = () => {
        setCatState({ id: "", title: "" })
        toggle()
    }

    const saveHandler = () => {
        if (!catState.title.trim()) {
            setError(true)
            return
        }
        setError(false)
        category.newCategory({ title: catState.title })
    }

    const editHandler = (c: CategoryWithLists) => () => {
        setCatState({ id: c.id, title: c.title })
        toggle()
    }

    const updateHandler = () => {
        if (!catState.title.trim()) {
            setError(true)
            return
        }
        setError(false)
        category.updateCategory(catState)
    }

    const removeHandler = () => {
        category.removeCategory(catState.id)
    }

    return (
        <div className={classes.projects}>
            <div className={classes.projects__header}>
                <p className={classes.projects__title}>Groups</p>
                <Button.Circle onClick={toggle} size='small' variant='grayPrimary'>
                    +
                </Button.Circle>
            </div>
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
                        onEdit={editHandler(cat)}
                    />
                ))
            ) : (
                <p className={classes.empty}>No groups have been created yet</p>
            )}

            {isOpen && (
                <Modal isOpen={isOpen} toggle={toggle}>
                    <Modal.Header title='Create group' onClose={toggle} />
                    <Modal.Content>
                        <Input
                            name='title'
                            placeholder='title'
                            value={catState.title}
                            error={error}
                            errorText={"Title is required"}
                            onChange={changeHandler}
                        />
                    </Modal.Content>
                    <Modal.Footer>
                        <div className={classes.btns}>
                            <Button size='small' onClick={cancelHandler} variant='grayPrimary'>
                                Cancel
                            </Button>
                            {catState.id !== "" && (
                                <Button size='small' variant='danger' onClick={removeHandler}>
                                    Remove
                                </Button>
                            )}
                            <Button
                                size='small'
                                onClick={catState.id !== "" ? updateHandler : saveHandler}
                            >
                                {catState.id !== "" ? "Update" : "Create"}
                            </Button>
                        </div>
                    </Modal.Footer>
                </Modal>
            )}
        </div>
    )
}
