import { Router, Route, Routes, createBrowserRouter, RouterProvider } from "react-router-dom";
import UserSelection from "./pages/UserSelection";
import MainLayout from "./MainLayout";
import Home from "./pages/Home";
import Income from "./pages/Income";
import Expenses from "./pages/Expenses";

function App() {
    return(
        <Routes>
            <Route path="/" element={<UserSelection />} />
            <Route path="/home" element={<MainLayout />}>
                <Route index element={<Home />} />
                <Route path="income" element={<Income />} />
                <Route path="expenses" element={<Expenses />} />
            </Route>
        </Routes>
    )
}

export default App
