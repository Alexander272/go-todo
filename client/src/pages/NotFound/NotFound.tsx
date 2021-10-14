import { LinkButton } from "../../components/Button/Button"
import classes from "./notFound.module.scss"

export const NotFound = () => {
    return (
        <div className='wrapper'>
            <div className={classes.center}>
                <div className={classes.number}>404</div>
                <div className={classes.text}>
                    <span>Ooops...</span>
                    <br />
                    page not found
                </div>
                <LinkButton
                    text='Home'
                    path='/'
                    className={classes.button}
                    testStyle={classes.button__text}
                    colorClass={classes.button__color}
                />
            </div>
        </div>
    )
}
