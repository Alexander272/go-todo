import { Children, cloneElement, PropsWithChildren, useEffect, useState } from "react"
import { Option } from "./Option"
import classes from "./select.module.scss"

type Props = {
    value: string
    onChange: (value: string) => void
}

const Select = ({ children, value, onChange }: PropsWithChildren<Props>) => {
    const [isOpen, setIsOpen] = useState(false)
    const [title, setTitle] = useState(value)

    useEffect(() => {
        Children.forEach(children as React.ReactElement[], (child: React.ReactElement) => {
            if (child.props.value === value) setTitle(child.props.children)
        })
    }, [value, children])

    const changeHandler = (curValue: string) => () => {
        setIsOpen(false)
        if (value === curValue) return
        onChange(curValue)
    }

    return (
        <div className={`${classes.select} ${isOpen ? "" : classes.close}`}>
            <p className={classes.selected} onClick={() => setIsOpen(prev => !prev)}>
                {title}
                <span className={classes.icon}>&#8910;</span>
            </p>
            <div className={`${classes.options} scroll`}>
                {Children.map(children as React.ReactElement[], (child: React.ReactElement) =>
                    cloneElement(child, { onClick: changeHandler(child.props.value) })
                )}
            </div>
        </div>
    )
}

Select.Option = Option

export { Select }
