import { users } from "wailsjs/go/models";

export function convertDateToBudgetStart(date: Date): string {
    return date.toISOString()
}

export function convertBudgetStartToDate(user: users.User): Date | null {
    let budgetDate: Date | null = null;

    if (user.budget_start) {
        budgetDate = new Date(user.budget_start);
    } else {
        budgetDate = null;
    }

    return budgetDate;
}
