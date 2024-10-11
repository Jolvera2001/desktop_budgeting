import { Button } from '@/components/ui/button'
import { Card, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Dialog, DialogTrigger } from '@/components/ui/dialog'
import useUserStore from '@/store/UserStore'
import { PlusCircle } from 'lucide-react'

function Expenses() {

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
                            <Dialog>
                                <DialogTrigger asChild>
                                    <Button variant="outline" size="icon">
                                        <PlusCircle className='h-4 w-4' />
                                    </Button>
                                </DialogTrigger>
                            </Dialog>
                        </CardHeader>
                    </Card>
                </div>
            </div>
        </>
    )
}

export default Expenses