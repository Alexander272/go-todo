import { BrowserRouter } from "react-router-dom"
import { MyRoutes } from "./components/routes"

function App() {
    return (
        <BrowserRouter>
            <MyRoutes />
        </BrowserRouter>
    )
}

export default App
