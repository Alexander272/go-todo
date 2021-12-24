import { PropsWithChildren } from "react"
import classes from "./modal.module.scss"

export const Content = ({ children }: PropsWithChildren<any>) => {
    return <main className={classes.content}>{children}</main>
}
