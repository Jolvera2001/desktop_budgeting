import { useEffect, useState } from 'react';
import { users } from "../../wailsjs/go/models"
import { Button } from "@/components/ui/button"
import { Input } from '@/components/ui/input';
import useUserStore from '@/store/UserStore'
import { useNavigate } from 'react-router-dom';
import ComboBox from '@/components/helpers/ComboBox'
import { Label } from '@/components/ui/label'
import { 
    Popover,
    PopoverTrigger,
    PopoverContent
 } from '@/components/ui/popover'
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
} from "../../wailsjs/go/users/UserService"
import { Plus } from 'lucide-react';
import { Toaster } from '@/components/ui/sonner';
import { toast } from 'sonner';

function UserSelection() {
    const navigate = useNavigate();
    const { selectedUser, setUser } = useUserStore();
    const [formBudgetPeriod, setFormBudgetPeriod] = useState<string>("")
    const [userList, setUserList] = useState<users.User[] | null>(null);
    const [chosenUser, setSelectedUser] = useState<users.User | null>(null);
    const updateUserList = (result: users.User[]) => setUserList(result);
    const selectUser = (user: users.User) => setSelectedUser(user);
    const getUsers = () => GetUsers().then(updateUserList);
    const deleteUser = async (userId: number | undefined) => {
        if (userId !== undefined) {
            try {
                await DeleteUser(userId);
                await getUsers();
                setSelectedUser(null);
            } catch (error) {
                console.log("error:", error)
            }
        } else {
            console.log("no user selected to complete process");
        }
    };

    const login = () => {
        if (chosenUser == null) {
            console.log("No user chosen");
            toast("No user has been chosen");
        } else {
            setUser(chosenUser);
            navigate("/home");
        }
    };

    const addUser = () => {
        let user: users.User;

        
    }
    
    useEffect(() => {
        getUsers();
        console.log(userList);
    }, []);
        
        function mapBudgetCycle(cycle: number | undefined): string {
        let cycleString: string;
        
        switch (cycle) {
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
    };

    return (
        <>
            <div className='flex flex-col space-y-4 items-center justify-center h-screen w-screen'>
                <p>Debug area</p>
                <p>Selected user: {chosenUser ? chosenUser._id : "none selected"}</p>
                <Card className='w-1/2'>
                    <CardHeader>
                        <CardTitle>Select User:</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <div className='flex flex-col space-y-4 max-h-64 overflow-y-auto'>
                            {userList ? (
                                userList.map((item, index) => (
                                    <Card key={index} onClick={() => selectUser(item)} className='hover:bg-slate-100 click:bg-slate-200 transition duration-150 ease-in-out'>
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
                        <div className='flex flex-row gap-4 w-full'>
                            <Popover>
                                <PopoverTrigger asChild>
                                    <Button>Add User</Button>
                                </PopoverTrigger>
                                <PopoverContent className='w-80'>
                                    <div className="grid gap-4">
                                        <div className='space-y-2'>
                                            <h1 className='font-medium leading-none'>Fill out the fields below</h1>
                                        </div>
                                        <div className='grid gap-2'>
                                            <div className='grid grid-cols-3 items-center gap-4'>
                                                <Label htmlFor='name'>Name</Label>
                                                <Input
                                                    id="name" 
                                                    className="col-span-2 h-8"
                                                />
                                            </div>
                                            <div className='grid grid-cols-3 items-center gap-4'>
                                                <Label htmlFor='email'>Email</Label>
                                                <Input
                                                    id="email"
                                                    className="col-span-2 h-8"
                                                />
                                            </div>
                                            <div className='grid grid-cols-3 items-center gap-4'>
                                                <Label htmlFor='budgetPeriod'>Budget Period</Label>
                                                <ComboBox
                                                    id="budgetPeriod"
                                                    className="col-span-2 h-8"
                                                    onChange={(value) => setFormBudgetPeriod(value)}
                                                />
                                            </div>
                                            <div className='grid grid-cols-3 items-center gap-4'>
                                                <Button>Add <Plus /></Button>
                                            </div>
                                        </div>
                                    </div>
                                </PopoverContent>
                            </Popover>
                            <Button className='bg-red-950' onClick={() => deleteUser(chosenUser?._id)}>Delete User</Button>
                            <Button className='bg-green-900 ml-auto' onClick={login}>Login</Button>
                        </div>
                    </CardFooter>
                </Card>
            </div>
            <Toaster />
        </>
    )
}

export default UserSelection
