import { useState } from "react";
import { Check, ChevronsUpDown } from "lucide-react";
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils";
import {
    Command,
    CommandEmpty,
    CommandGroup,
    CommandInput,
    CommandItem,
    CommandList, 
} from "@/components/ui/command"
import {
    Popover,
    PopoverContent,
    PopoverTrigger, 
} from "@/components/ui/popover"

interface ComboBoxProps {
    className: string;
    id: string;
    onChange: (value: string) => void;
}

const budgetPeriods = [
    {
        value: "Monthly",
        label: "Monthly",
    },
    {
        value: "Biweekly",
        label: "Biweekly",
    },
    {
        value: "Weekly",
        label: "Weekly",
    }
]

const ComboBox: React.FC<ComboBoxProps> = ({ className, id, onChange}) => {
    const [open, setOpen] = useState<boolean>(false)
    const [value, setValue] = useState<string>("")

    return(
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button 
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className={className}
                >
                    {value 
                    ? budgetPeriods.find((period) => period.value === value)?.label
                    : "Select Budget Period..."}
                    <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="w-[200px] p-0">
                <Command>
                    <CommandInput placeholder="Search framework..." />
                    <CommandList>
                        <CommandEmpty>No framework found.</CommandEmpty>
                        <CommandGroup>
                            {budgetPeriods.map((period) => (
                                <CommandItem
                                    key={period.value}
                                    value={period.value}
                                    onSelect={(currentValue) => {
                                        setValue(currentValue === value ? "" : currentValue)
                                        onChange(currentValue)
                                        setOpen(false)
                                    }}
                                >
                                    <Check
                                        className={cn(
                                            "mr-2 h-4 w-4",
                                            value === period.value ? "opacity-100" : "opacity-0"
                                        )}
                                    />
                                    {period.label}
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}

export default ComboBox