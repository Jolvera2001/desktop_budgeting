import BudgetAddDialog from '@/components/reusables/AddBudgetDialog'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import useUserStore from '@/store/UserStore'

function Expenses() {
    const { selectedUser } = useUserStore();
    return (
        <>
            <div className="container mx-auto p-4 mt-4">
                <h1 className='text-3xl font-bold mb-6'>Expenses Overview</h1>
                <div className="grid gap-6 md:grid-cols-2">
                    <Card>
                        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                            <div>
                                <CardTitle>Budgets</CardTitle>
                                <CardDescription>Your spending limits and current status</CardDescription>
                            </div>
                            <BudgetAddDialog />
                        </CardHeader>
                        <CardContent>
                            {selectedUser?.budgets 
                            ? selectedUser?.budgets.map((budget, index) => (
                                <div key={budget.id} className='mb-4'>
                                    <div className='flex justify-between items-center mb-2'>
                                        <span className='font-medium'>{budget.name}</span>
                                        <span className='text-sm text-gray-500'>
                                            ${budget.amount.toFixed(2)}
                                        </span>
                                    </div>
                                </div>
                            ))
                            :<p>No Budgets Created Yet</p>}
                        </CardContent>
                    </Card>
                </div>
            </div>
        </>
    )
}

export default Expenses