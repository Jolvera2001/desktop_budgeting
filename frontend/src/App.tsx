import { useEffect, useState } from 'react';
import { users } from "../wailsjs/go/models" 
import { Button } from "@/components/ui/button"
import { Input } from '@/components/ui/input';
import {
    Card,
    CardContent,
    CardHeader,
    CardFooter,
    CardDescription,
    CardTitle
} from "@/components/ui/card"
import { 
    CreateUser,
    UpdateUser,
    GetUser,
    GetUsers,
    DeleteUser
 } from "../wailsjs/go/users/UserService"

function App() {
    const [userForm, setUserForm] = useState<users.UserDto | null>(null);
    const [userList, setUserList] = useState<users.User[] | null>(null);
    const updateUserList = (result: users.User[]) => setUserList(result)
    const getUsers = () => GetUsers().then(updateUserList)

    useEffect(() => {
        getUsers();
        console.log(userList);
    }, []);

    function mapBudgetCycle(cycle: number | undefined): string {
        let cycleString: string;

        switch(cycle) {
            case 1:
                cycleString = "Monthly";
                break;
            case 2:
                cycleString = "Biweekly";
                break;
            case 3: 
                cycleString = "Weekly";
                break;
            default:
                cycleString = "No cycle set";
        }

        return cycleString;
    }


    return (
        <>
            <div className='flex flex-col space-y-4 items-center justify-center h-screen w-screen'>
                <Card className='w-1/3'>
                    <CardHeader>
                        <CardTitle>Select User:</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <div className='flex flex-col space-y-4 max-h-64 overflow-y-auto'>
                            {userList ? (
                                userList.map((item, index) => (
                                    <Card key={index}>
                                        <CardHeader>
                                            <CardTitle>{item.name}</CardTitle>
                                        </CardHeader>
                                        <CardContent>
                                            <p>Budget Cycle: {mapBudgetCycle(item.budget_period)}</p>
                                        </CardContent>
                                    </Card>
                                ))
                            ) : (
                                <p>No users yet</p>
                            )}
                        </div>
                    </CardContent>
                    <CardFooter>
                        <Button>Add User</Button>
                    </CardFooter>
                </Card>
            </div>
        </>
    )
}

export default App
