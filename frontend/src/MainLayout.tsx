import { Outlet, useLocation, useNavigate } from "react-router-dom"
import { Tabs, TabsList, TabsContent, TabsTrigger } from "@/components/ui/tabs"
import { Home, CircleDollarSign, Wallet } from "lucide-react"

function MainLayout() {
    const navigate = useNavigate();
    const location = useLocation();

    const getCurrentTab = () => {
        const path = location.pathname;
        if (path === '/home') return 'home';
        if (path === '/home/income') return 'income';
        if (path === '/home/expenses') return 'expenses';
        return 'home'; // default to home if path doesn't match
    }

    const handleTabChange = (value: string) => {
        switch(value) {
            case 'home':
                navigate('/home');
                break;
            case 'income':
                navigate('/home/income');
                break;
            case 'expenses':
                navigate('/home/expenses');
                break;
            default:
                navigate('/home');
        }
    }

    return(
        <>
        <style>
            {`
                @media (max-width: 1040px) {
                    .responsive-tab {
                        justify-content: center;
                    }
                    
                    .responsive-tab .tab-text {
                        display: none;
                    }
                }
            `}
        </style>
            <div className="grid grid-cols-10 grid-rows-[auto_1fr] h-screen bg-gray-200">
                <div className="col-span-1 row-span-1 bg-gray-200 border-b-slate-300"></div>
                <div className="col-span-9 row-span-1 bg-gray-200 border-b-slate-300 min-h-[50px]">

                </div>
                <div className="col-span-1 bg-gray-200 border-b-slate-300">
                    <Tabs value={getCurrentTab()} onValueChange={handleTabChange} orientation="vertical" className="w-full">
                        <TabsList className="flex flex-col h-full space-y-2 bg-gray-100 p-2">
                            <TabsTrigger value="home" className="w-full flex items-center justify-start gap-2 responsive-tab">
                                <Home size={16} />
                                <span className="tab-text">Home</span>
                            </TabsTrigger>
                            <TabsTrigger value="expenses" className="w-full flex items-center justify-start gap-2 responsive-tab">
                                <CircleDollarSign size={16} />
                                <span className="tab-text">Expenses</span>
                            </TabsTrigger>
                            <TabsTrigger value="income" className="w-full flex items-center justify-start gap-2 responsive-tab">
                                <Wallet size={16} />
                                <span className="tab-text">Income</span>
                            </TabsTrigger>
                        </TabsList>
                    </Tabs>
                </div>
                <div className="col-span-9 border-b-slate-500 bg-white">
                    <Outlet />
                </div>
            </div>
        </>
    )
}

export default MainLayout