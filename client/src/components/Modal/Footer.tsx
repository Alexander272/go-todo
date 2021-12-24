import { PropsWithChildren } from "react"
import classes from "./modal.module.scss"

export const Footer = ({ children }: PropsWithChildren<any>) => {
    return <footer className={classes.footer}>{children}</footer>
}
