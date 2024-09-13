export function mapNumToBudget(cycle: number | undefined): string {
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

export function mapBudgetToNum(cycle: string): number {
    let cycleNum: number;

    switch (cycle) {
        case "Monthly":
            cycleNum = 1;
            break;
        case "Biweekly":
            cycleNum = 2;
            break;
        case "Weekly":
            cycleNum = 3;
            break;
        default:
            cycleNum = 0;
    }

    return cycleNum;
}