import { Link } from "react-router-dom"
import classes from "./notfound.module.scss"

export default function NotFoundPage() {
    return (
        <div className={classes.page}>
            <h3 className={classes.title}>404</h3>
            <p className={classes.text}>
                Maybe this page moved? Got deleted? Is hiding out in quarantine? Never existed in
                the first place?
            </p>

            <div className={classes.buttons}>
                <Link className={classes.linkButton} to='/'>
                    Home
                </Link>
                <p className={classes.linkButton}>Back</p>
            </div>
        </div>
    )
}
