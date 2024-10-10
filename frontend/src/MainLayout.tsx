import { Outlet } from "react-router-dom"
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/components/ui/tabs"
import { Home } from "lucide-react"

function MainLayout() {
    return(
        <div className="grid grid-cols-10 grid-rows-[auto_1fr] h-screen bg-gray-200">
            <div className="col-span-1 row-span-1 bg-gray-200 border-b-slate-300"></div>
            <div className="col-span-9 row-span-1 bg-gray-200 border-b-slate-300 min-h-[50px]">

            </div>
            <div className="col-span-1 bg-gray-200 border-b-slate-300">
                <Tabs defaultValue="home" orientation="vertical" className="w-full">
                    <TabsList className="flex flex-col h-full space-y-2 bg-gray-100 p-2">
                        <TabsTrigger value="home" className="w-full flex items-center justify-start gap-2">
                            <Home size={16} />
                            Home
                        </TabsTrigger>
                    </TabsList>
                </Tabs>
            </div>
            <div className="col-span-9 border-b-slate-500 bg-white">
                <Outlet />
            </div>
        </div>
    )
}

export default MainLayout