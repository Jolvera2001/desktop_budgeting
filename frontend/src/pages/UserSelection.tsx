import { useEffect, useState } from 'react';
import { users } from "../../wailsjs/go/models"
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
} from "../../wailsjs/go/users/UserService"

function UserSelection() {
    const [userForm, setUserForm] = useState<users.UserDto | null>(null);
    const [userList, setUserList] = useState<users.User[] | null>(null);
    const [selectedUser, setSelectedUser] = useState<users.User | null>(null);
    const updateUserList = (result: users.User[]) => setUserList(result)
    const selectUser = (user: users.User) => setSelectedUser(user)

    const getUsers = () => GetUsers().then(updateUserList)
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
    }


    return (
        <>
            <div className='flex flex-col space-y-4 items-center justify-center h-screen w-screen'>
                <p>Debug area</p>
                <p>Selected user: {selectedUser ? selectedUser._id : "none selected"}</p>
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
                        <div className='flex flex-row space-x-4'>
                            <Button>Add User</Button>
                            <Button className='bg-red-950' onClick={() => deleteUser(selectedUser?._id)}>Delete User</Button>
                        </div>
                    </CardFooter>
                </Card>
            </div>
        </>
    )
}

export default UserSelection
