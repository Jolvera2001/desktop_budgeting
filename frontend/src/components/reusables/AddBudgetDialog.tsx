import { useState } from "react";
import { Dialog, DialogHeader, DialogTitle, DialogTrigger, DialogContent, DialogDescription } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import { PlusCircle } from "lucide-react";
import useUserStore from "../../store/UserStore"

import { Make } from "../../../wailsjs/go/services/BudgetService"
// import { BudgetDto } from "../../../wailsjs/go/models"

function BudgetAddDialog() {
    const { selectedUser } = useUserStore();
    const [inputName, setInputName] = useState<string>("");
    const [inputAmount, setInputAmount] = useState<string>("");

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const budgetDto: BudgetDto = new BudgetDto({
            user_id: selectedUser?.id,
            name: inputName,
            amount: parseFloat(inputAmount)
        });

        try {
            const result = await Make(budgetDto);
            console.log("Budget added", result);
            setInputAmount("");
            setInputName("");
        } catch (e) {
            console.error("error adding budget: ", e);
        }
    }

    return(
        <Dialog>
            <DialogTrigger asChild>
                <Button variant="outline" size="icon">
                    <PlusCircle className="h-4 w-4" />
                </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Add New Budget</DialogTitle>
                    <DialogDescription>Create a new budget category and set a spending limit</DialogDescription>
                </DialogHeader>
                <form onSubmit={handleSubmit}>
                    <div className="grip gap-4 py-4">
                        <div className="grid grid-cols-4 items-center gap-4">
                            <Label htmlFor="name">Name</Label>
                            <div className="col-span-3">
                                <Input 
                                    id="name"
                                    value={inputName}
                                    onChange={(e) => setInputName(e.target.value)} 
                                />
                            </div>
                            <Label htmlFor="amount">Amount</Label>
                            <div className="col-span-3">
                                <Input 
                                    id="amount" 
                                    type="number" 
                                    min={0} 
                                    max={100000000} 
                                    step={0.01} 
                                    placeholder="0.00"
                                    value={inputAmount}
                                    onChange={(e) => setInputAmount(e.target.value)} 
                                />
                            </div>
                        </div>
                    </div>
                </form>
            </DialogContent>
        </Dialog>
    )
}

export default BudgetAddDialog