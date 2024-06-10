
// 在這裡實作所有和後端溝通的 API
// getExpenses, listexpense, createExpense, updateExpense, deleteExpense
// 後端 host 為 localhost:6741
// 後端 endpoint prefix 為 /api/v1/expense
// 用 fetch 來發送 request

import { Expense } from '../types/Expense';

// getExpenses 用來取得所有 expenses
export const getExpenses = async (id: number): Promise<Expense[]> => {
    const res = await fetch(`http://localhost:6741/api/v1/expenses/${id}`, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('ai-training-token')}`,
        }
    });
    const data = await res.json();
    return data;
};

// listexpense 用來取得符合條件的 expenses
export const listExpenses = async (search: string, filterDate: string): Promise<Expense[]> => {
    const res = await fetch(`http://localhost:6741/api/v1/expenses?search=${search}&filterDate=${filterDate}`, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('ai-training-token')}`,
        }
    });
    if (!res.ok) {
        return [];
    }
    const data = await res.json();
    return data;
};

// createExpense 用來新增一筆 expense
export const createExpense = async (expense: Omit<Expense, 'id'>): Promise<void> => {
    await fetch('http://localhost:6741/api/v1/expenses', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem('ai-training-token')}`,
        },
        body: JSON.stringify(expense),
    });
};

// updateExpense 用來更新一筆 expense
export const updateExpense = async (id: number, expense: Omit<Expense, 'id'>): Promise<void> => {
    await fetch(`http://localhost:6741/api/v1/expenses/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${localStorage.getItem('ai-training-token')}`,
        },
        body: JSON.stringify(expense),
    });
};

// deleteExpense 用來刪除一筆 expense
export const deleteExpense = async (id: number): Promise<void> => {
    await fetch(`http://localhost:6741/api/v1/expenses/${id}`, {
        method: 'DELETE',
        headers: {
            Authorization: `Bearer ${localStorage.getItem('ai-training-token')}`,
        }
    });
};
