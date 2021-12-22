import { lazy } from "react"
import { Route, Routes } from "react-router-dom"

import AuthPage from "../pages/Auth/Auth"

const NotFoundPage = lazy(() => import("../pages/NotFound/NotFound"))

export const MyRoutes = () => {
    return (
        <Routes>
            <Route path='/auth/' element={<AuthPage />} />
            <Route path='*' element={<NotFoundPage />} />
        </Routes>
    )
}
