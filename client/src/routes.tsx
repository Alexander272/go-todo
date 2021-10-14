import { Route, Switch } from "react-router-dom"
import { LoginPage } from "./pages/Login/Login"
import { NotFound } from "./pages/NotFound/NotFound"

export const Routes = () => {
    let routes = (
        <Switch>
            <Route path='/' exact>
                <LoginPage />
            </Route>
            <Route path='*'>
                <NotFound />
            </Route>
        </Switch>
    )

    return routes
}
