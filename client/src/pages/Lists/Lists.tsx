import { useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { ListsScroll } from "../../components/ListsScroll/ListsScroll"
import { useModal } from "../../components/Modal/hooks/useModal"
import { Modal } from "../../components/Modal/Modal"
import { Button } from "../../components/UI/Button/Button"
import { Input } from "../../components/UI/Input/Input"
import { Select } from "../../components/UI/Select/Select"
import { Dispatch, RootState } from "../../store/store"
import { NewList } from "../../types/list"
import Tasks from "../Tasks/Tasks"
import classes from "./lists.module.scss"

export default function ListsPage() {
    const selCat = useSelector((state: RootState) => state.category.selectedCategories)
    const categories = useSelector((state: RootState) => state.category.categories)

    const [error, setError] = useState(false)
    const [listState, setListState] = useState<NewList>({
        categoryId: selCat[0],
        title: "",
        description: "",
    })

    const { isOpen, toggle } = useModal()

    const { category } = useDispatch<Dispatch>()

    const selectHandler = (value: string) => {
        setListState(prev => ({ ...prev, categoryId: value }))
    }

    const changeHandler = (event: React.ChangeEvent<HTMLInputElement>) => {
        setListState(prev => ({ ...prev, [event.target.name]: event.target.value }))
    }

    const openHandler = () => {
        setListState({ categoryId: selCat[0], title: "", description: "" })
        toggle()
    }
    const saveHandler = () => {
        if (!listState.title.trim()) {
            setError(true)
            return
        }
        setError(false)
        category.newList(listState)
    }

    return (
        <div className={classes.page}>
            {isOpen && (
                <Modal isOpen={isOpen} toggle={toggle}>
                    <Modal.Header title='Create list' onClose={toggle} />
                    <Modal.Content>
                        <Select value={listState.categoryId} onChange={selectHandler}>
                            {categories.map(cat => (
                                <Select.Option key={cat.id} value={cat.id}>
                                    {cat.title}
                                </Select.Option>
                            ))}
                        </Select>
                        <Input
                            name='title'
                            placeholder='title'
                            value={listState.title}
                            error={error}
                            errorText={"Title is required"}
                            onChange={changeHandler}
                        />
                        <Input
                            name='description'
                            placeholder='description'
                            value={listState.description}
                            onChange={changeHandler}
                        />
                    </Modal.Content>
                    <Modal.Footer>
                        <div className={classes.btns}>
                            <Button size='small' onClick={toggle} variant='grayPrimary'>
                                Cancel
                            </Button>
                            <Button size='small' onClick={saveHandler}>
                                Create
                            </Button>
                        </div>
                    </Modal.Footer>
                </Modal>
            )}

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
                    <Button size='small' rounded='round' onClick={openHandler}>
                        Add list
                    </Button>
                </div>
            </div>
            <Tasks />
        </div>
    )
}
