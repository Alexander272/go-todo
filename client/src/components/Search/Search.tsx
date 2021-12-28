import classes from "./search.module.scss"

export const Search = () => {
    return (
        <div className={classes.search}>
            <input className={classes.input} placeholder='Search...' />
        </div>
    )
}
