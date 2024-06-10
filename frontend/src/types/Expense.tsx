
// 在這裡實作 expense/page.tsx 要用到的 Expense type
export type Expense = {
    id: number;
    title: string;
    amount: number;
    date: string;
    updater: string;
    category: string;
  };
  
  export type CreateExpensePayload = {
    title: string;
    amount: number;
    category: string;
    updater: string;
    createdAt: string;
  }
  export type UpdateExpensePayload = {
    title: string;
    amount: number;
    category: string;
    updater: string;
  }
