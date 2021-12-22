import { Provider } from 'react-redux'
import { BrowserRouter } from 'react-router-dom'
import { MyRoutes } from './components/routes'
import { store } from './store/store'

function App() {
    return (
        <Provider store={store}>
            <BrowserRouter>
                <MyRoutes />
            </BrowserRouter>
        </Provider>
    )
}

export default App
