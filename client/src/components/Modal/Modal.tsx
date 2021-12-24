import { PropsWithChildren } from "react"
import { createPortal } from "react-dom"
import { Backdrop } from "../Backdrop/Backdrop"
import { Content } from "./Content"
import { Footer } from "./Footer"
import { Header } from "./Header"
import classes from "./modal.module.scss"

type Props = {
    isOpen: boolean
    toggle: () => void
}

const Modal = ({ children, isOpen, toggle }: PropsWithChildren<Props>) => {
    if (!isOpen) return null

    return createPortal(
        <>
            <Backdrop onClick={toggle} />
            <div className={classes.modal}>{children}</div>
        </>,
        document.getElementById("root") as HTMLElement
    )
}

Modal.Header = Header
Modal.Content = Content
Modal.Footer = Footer

export { Modal }
