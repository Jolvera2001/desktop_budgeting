import { Router, Route, Routes, createBrowserRouter, RouterProvider } from "react-router-dom";
import UserSelection from "./pages/UserSelection";

function App() {
    return(
        <Routes>
            <Route path="/" element={<UserSelection />} />
            <Route path="/home" element={<MainLayout />}>
                <Route index element={<Home />} />
                <Route path="budgets" element={<Budgets />} />
                <Route path="transactions" element={<Transactions />} />
            </Route>
        </Routes>
    )
}

export default App
