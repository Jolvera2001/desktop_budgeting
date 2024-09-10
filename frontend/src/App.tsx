import {useState} from 'react';
import {Greet} from "../wailsjs/go/main/App";
import {
    Card,
    CardContent,
    CardHeader,
    CardFooter,
    CardDescription,
    CardTitle
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from '@/components/ui/input';

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(name).then(updateResultText);
    }

    return (
        <>
            <div className='flex flex-col space-y-4 items-center justify-center h-screen w-screen'>
                <Card>
                    <CardHeader>
                        <CardTitle>Example Page</CardTitle>
                    </CardHeader>
                    <CardContent>
                        <div className='space-y-4'>
                            <p>{resultText}</p>
                            <Input onChange={updateName}></Input>
                            <Button onClick={greet}>Submit</Button>
                        </div>
                    </CardContent>
                </Card>
                <Card>
                    <CardHeader>
                        <CardTitle>Select User:</CardTitle>
                    </CardHeader>
                    <CardContent>

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
