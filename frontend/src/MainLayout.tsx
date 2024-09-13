import { Outlet } from "react-router-dom"

function MainLayout() {
    return(
        <div className="grid grid-cols-4 grid-rows-[auto_1fr] h-screen">
            <div className="col-span-1 row-span-1 bg-gray-200 border-b-slate-500"></div>
            <div className="col-span-3 row-span-1 bg-gray-200 border-b-slate-500">

            </div>
            <div className="col-span-1 bg-gray-200 border-b-slate-500"></div>
            <div className="col-span-3 border-b-slate-500">
                <Outlet />
            </div>
        </div>
    )
}

export default MainLayout