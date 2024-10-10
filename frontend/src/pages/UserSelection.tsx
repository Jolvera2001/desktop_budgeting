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
    // use to go to other pages
    const navigate = useNavigate(); 

    // user store operations
    const { selectedUser, setUser } = useUserStore(); 

    // react state
    const [userList, setUserList] = useState<models.User[]>([]);
    const [newUser, setNewUser] = useState({name: "", email: "", password: ""});
    const [loginPassword, setLoginPassword] = useState("");

    const updateUserList = (result: models.User[]) => {
        console.log(result)
        if (Array.isArray(result)) {
            setUserList(result)
        } else {
            console.error("Expected an array of users, but received:", result);
            setUserList([]); // Fallback to an empty array if the result is invalid
        }
    };
    const getUsers = () => GetAllProfiles().then(updateUserList);

    const handleLogin = async () => {
        if (selectedUser) {
            try {
                const loggedInUser = await Login(selectedUser.id, loginPassword, selectedUser.password)
                setUser(loggedInUser)
                navigate("/home")
            } catch (error) {
                console.log(error)
            }
        } else {
            console.log("no user selected")
        }
    }

    const handleDelete = async () => {
        if (selectedUser?.id) {
            try {
                await DeleteProfile(selectedUser?.id);
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
        if (newUser.name && newUser.email && newUser.password) {
            console.log("called handleAddUser")
            const date = new Date()
            const newUserDto = new models.UserDto();

            newUserDto.name = newUser.name;
            newUserDto.email = newUser.email;
            newUserDto.password = newUser.password

            try {
                console.log("Creating user...")
                await Register(newUserDto);
                await getUsers();

                setNewUser({ name: "", email: "", password: "" })
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
                                key={user.id}
                                className={`p-4 cursor-pointer hover:bg-gray-100 ${selectedUser?.id === user.id ? "bg-blue-100" : ""}`}
                                onClick={() => handleSelectUser(user)}
                            >
                                <div className='flex items-center'>
                                    <div>
                                        <p>{user.name}</p>
                                        <p>{user.email}</p>
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
                            <Label htmlFor='password'>Password</Label>
                            <Input 
                                id='password'
                                value={newUser.password}
                                type='password'
                                onChange={(e) => setNewUser({ ...newUser, password: e.target.value })}
                            />
                        </div>
                        <Button type='submit'><Plus />Add User</Button>
                    </form>
                    {selectedUser && (
                        <div className='mt-8'>
                            <h3 className="text-xl font-semibold mb-2">Selected User</h3>
                            <p>Name: {selectedUser.name}</p>
                            <p>Email: {selectedUser.email}</p>
                            <div> 
                                <Label htmlFor="loginPassword">Login</Label>
                                <Input
                                    id='loginPassword'
                                    value={loginPassword}
                                    type='password'
                                    onChange={(e) => setLoginPassword(e.target.value)} 
                                />
                            </div>
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
