import { useEffect, useState } from 'react';
import { users } from "../../wailsjs/go/models"
import { Button } from "@/components/ui/button"
import { Input } from '@/components/ui/input';
import useUserStore from '@/store/UserStore'
import { useNavigate } from 'react-router-dom';
import { Label } from '@/components/ui/label'
import { Plus } from 'lucide-react';
import { ScrollArea } from '@/components/ui/scroll-area'
import {
    CreateUser,
    UpdateUser,
    GetUser,
    GetUsers,
    DeleteUser
} from "../../wailsjs/go/users/UserService"
import { mapNumToBudget, mapBudgetToNum } from '@/lib/mapBudgetPeriod';
import { RadioGroup } from '@/components/ui/radio-group';

function UserSelection() {
    const navigate = useNavigate(); // use to go to other pages

    const { selectedUser, setUser } = useUserStore();
    const [userList, setUserList] = useState<users.User[] | null>(null);
    const [newUser, setSelectedUser] = useState({name: "", email: "", budgetPeriod: ""});
    const updateUserList = (result: users.User[]) => setUserList(result);
    const getUsers = () => GetUsers().then(updateUserList);

    const handleSelectUser = (user: users.User) => {
        setUser(user)
    }

    const handleAddUser = (e: React.FormEvent) => {
        console.log("added user...")
    }
    
    useEffect(() => {
        getUsers();
        console.log(userList);
    }, []);

    return (
        <>
            <div className='flex h-screen bg-gray-100'>
                <div className='flex-1 flex flex-col p-8'>
                    <h2 className="text-2xl font-bold mb-4">User List</h2>
                    <ScrollArea className="flex-1 border rounded-md bg-white">
                        {userList ? userList?.map((user) => (
                            <div 
                                key={user._id}
                                className={`p-4 cursor-pointer hover:bg-gray-100 ${selectedUser?._id === user._id ? "bg-blue-100" : ""}`}
                                onClick={() => handleSelectUser(user)}
                            >
                                <div className='flex items-center'>
                                    <div>
                                        <p>{user.name}</p>
                                        <p>{user.email}</p>
                                        <p>Budget Period: {mapNumToBudget(user.budget_period)}</p>
                                    </div>
                                </div>
                            </div>
                        )) : (
                            <div className='flex items-center'>
                               
                            </div>
                        )}
                    </ScrollArea>
                </div>
                <div className='flex-1 p-8'>
                    <h2 className="text-2xl font-bold mb-4">Add New User</h2>
                    <form onSubmit={handleAddUser} className='space-y-4'>
                        <div>
                            <Label htmlFor='name'>Name</Label>
                            <Input id='name'/>
                        </div>
                        <div>
                            <Label htmlFor='email'>Email</Label>
                            <Input id='email'/>
                        </div>
                        <div>
                            <Label htmlFor='budgetPeriod'>Budget Period</Label>
                            <RadioGroup id='budgetPeriod'>

                            </RadioGroup>
                        </div>
                        <Button type='submit'><Plus />Add User</Button>
                    </form>
                    {selectedUser && (
                        <div className='mt-8'>
                            <h3 className="text-xl font-semibold mb-2">Selected User</h3>
                            <p>Name: {newUser.name}</p>
                            <p>Email: {newUser.email}</p>
                        </div>
                    )}
                </div>
            </div>
        </>
    )
}

export default UserSelection
