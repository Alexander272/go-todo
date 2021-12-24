import { lazy, Suspense } from "react"
import { Route, Routes } from "react-router-dom"

import { MainLayout } from "./Layout/Main/Main"
import AuthPage from "../pages/Auth/Auth"
import ListsPage from "../pages/Lists/Lists"

const NotFoundPage = lazy(() => import("../pages/NotFound/NotFound"))

export const MyRoutes = () => {
    return (
        <Suspense fallback='loading...'>
            <Routes>
                <Route path='/' element={<MainLayout />}>
                    <Route index element={<ListsPage />} />
                </Route>
                <Route path='/auth/' element={<AuthPage />} />
                <Route path='*' element={<NotFoundPage />} />
            </Routes>
        </Suspense>
    )
}
