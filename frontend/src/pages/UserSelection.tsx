import { useEffect, useState } from 'react';
import { models } from "../../wailsjs/go/models"
import { Button } from "@/components/ui/button"
import { Input } from '@/components/ui/input';
import useUserStore from '@/store/UserStore'
import { useNavigate } from 'react-router-dom';
import { Label } from '@/components/ui/label'
import { House, Key, Plus, Trash } from 'lucide-react';
import { ScrollArea } from '@/components/ui/scroll-area'
import {
    Register,
    Login,
    GetAllProfiles,
    UpdateProfile,
    DeleteProfile
} from "../../wailsjs/go/services/UserService"
import { mapNumToBudget, mapBudgetToNum } from '@/lib/mapBudgetPeriod';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';

function UserSelection() {
    const navigate = useNavigate(); // use to go to other pages

    const { selectedUser, setUser } = useUserStore();
    const [userList, setUserList] = useState<models.User[] | null>(null);
    const [newUser, setNewUser] = useState({name: "", email: "", budgetPeriod: ""});
    const updateUserList = (result: models.User[]) => setUserList(result);
    const getUsers = () => GetAllProfiles().then(updateUserList);

    const handleLogin = () => {
        if (selectedUser) {
            try {
                navigate("/home")
            } catch (error) {
                console.log(error)
            }
        } else {
            console.log("no user selected")
        }
    }

    const handleDelete = async () => {
        if (selectedUser?._id) {
            try {
                await DeleteProfile(selectedUser?._id);
                await getUsers();
                setUser(null);
            } catch (error) {
                console.log(error);
            }
        } else {
            console.log("how did you get here?");
        }
    }

    const handleSelectUser = (user: models.User) => {
        setUser(user)
    }

    const handleAddUser = async (e: React.FormEvent) => {
        e.preventDefault();
        if (newUser.name && newUser.email && newUser.budgetPeriod) {
            console.log("called handleAddUser")
            const date = new Date()
            const newUserDto = new models.UserDto();

            newUserDto.name = newUser.name;
            newUserDto.email = newUser.email;

            try {
                console.log("Creating user...")
                await Register(newUserDto);
                await getUsers();

                setNewUser({ name: "", email: "", budgetPeriod: "" })
            } catch (error) {
                console.log(error);
            }
        } else {
            console.log("fields empty!");
        }
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
                            <Input 
                                id='name'
                                value={newUser.name}
                                onChange={(e) => setNewUser({ ...newUser, name: e.target.value })}
                            />
                        </div>
                        <div>
                            <Label htmlFor='email'>Email</Label>
                            <Input 
                                id='email'
                                value={newUser.email}
                                onChange={(e) => setNewUser({ ...newUser, email: e.target.value })}
                            />
                        </div>
                        <div>
                            <Label htmlFor='budgetPeriod'>Budget Period</Label>
                            <RadioGroup 
                                id='budgetPeriod' 
                                value={newUser.budgetPeriod}
                                onValueChange={(value) => setNewUser({ ...newUser, budgetPeriod: value })}
                            >
                                <div className='flex flex-row space-x-4'>
                                    <div className="flex items-center space-x-2">
                                        <RadioGroupItem value="Monthly" id="r1" />
                                        <Label htmlFor="r1">Monthly</Label>
                                    </div>
                                    <div className="flex items-center space-x-2">
                                        <RadioGroupItem value="Biweekly" id="r2" />
                                        <Label htmlFor="r2">Biweekly</Label>
                                    </div>
                                    <div className="flex items-center space-x-2">
                                        <RadioGroupItem value="Weekly" id="r3" />
                                        <Label htmlFor="r3">Weekly</Label>
                                    </div>
                                </div>
                            </RadioGroup>
                        </div>
                        <Button type='submit'><Plus />Add User</Button>
                    </form>
                    {selectedUser && (
                        <div className='mt-8'>
                            <h3 className="text-xl font-semibold mb-2">Selected User</h3>
                            <p>Name: {selectedUser.name}</p>
                            <p>Email: {selectedUser.email}</p>
                            <div className='flex flex-row gap-5 mt-4'>
                                <Button className='gap-1 bg-green-700' onClick={handleLogin}><House />Login</Button>
                                <Button className='gap-1 bg-red-800' onClick={handleDelete}><Trash />Delete</Button>
                            </div>   
                        </div>
                    )}
                </div>
            </div>
        </>
    )
}

export default UserSelection
