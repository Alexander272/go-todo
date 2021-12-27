import { FC } from "react"
import classes from "./checkbox.module.scss"

type Props = {
    label?: string
    id: string
    checked: boolean
    name?: string
    readOnly?: boolean
    onClick?: (event: React.MouseEvent<HTMLInputElement>) => void
}

export const Checkbox: FC<Props> = ({ label, id, checked, name, readOnly, onClick }) => {
    return (
        <div className={classes.checkbox}>
            <input
                className={classes.input}
                id={id}
                type='checkbox'
                checked={checked}
                name={name}
                onClick={onClick}
                readOnly={readOnly}
            />
            <label className={classes.label} htmlFor={id}>
                <span>
                    <svg width='12px' height='9px' viewBox='0 0 12 9'>
                        <polyline points='1 5 4 8 11 1'></polyline>
                    </svg>
                </span>
                {label && <span className={classes.text}>{label}</span>}
            </label>
        </div>
    )
}
