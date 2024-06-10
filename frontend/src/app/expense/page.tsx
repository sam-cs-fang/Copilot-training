'use client'
// 請參考 README.md 的說明，在這個頁面上，你需要完成以下功能：
// 1. 顯示所有的支出資料
// 2. 透過 API 新增支出資料
// 3. 透過 API 刪除支出資料
// 4. 透過 API 更新支出資料
// 5. 搜尋欄位，可以透過名稱搜尋支出資料，且支援模糊搜尋
// 6. 篩選欄位，可以透過時間篩選支出資料
// style 請使用 bootstrap like
// 以單純 css 的方式撰寫即可，不需要使用其他 css framework
// 以下是我想要的 style，請在產生 HTML 時同步考慮這些 style
// 1. 背景是一個純白背景
// 2. 頁面中間有一個區域是主要內容，最大寬度為視窗的 80%，置中
// 3. 主要內容區域包含一個標題，標題的上下 padding 為 20px，左右 padding 為 0，文字置中，內容是「支出管理」
// 4. 主要內容區域包含一個按鈕，按鈕的上下 padding 為 20px，左右 padding 為 0，文字置中，內容是「新增支出」
// 5. 主要內容區域包含一個表格，表格的上下 padding 為 20px，左右 padding 為 0，文字置中，會根據支出資料的數量動態增加列數
// 6. 表格的標題列背景色是 #f0f0f0，文字顏色是黑色，字體大小是 16px，字體粗細是 bold
// 7. 表格的內容列背景色是白色，文字顏色是黑色，字體大小是 14px，字體粗細是 normal
// 8. 表格的內容列中，最後一個欄位是操作欄，會包含三個按鈕，分別是「編輯」、「刪除」
// 9. 按鈕的背景色是 #007bff，文字顏色是白色，字體大小是 14px，字體粗細是 normal
// 10. 表格的內容列中，日期格式為 yyyy-mm-dd
// 11. 表格的內容列中，金額格式為 $1,000.00
// 12. 表格與標題之間有一個 input 欄位，用來搜尋支出資料，input 欄位的上下 padding 為 10px，左右 padding 為 20px
// 13. input 欄位的 placeholder 是「搜尋支出」
// 14. input 欄位右邊有一個日期選擇器，用來篩選支出資料，日期選擇器的上下 padding 為 10px，左右 padding 為 20px

import './style.css';
import React, { ChangeEventHandler, useEffect, useState } from 'react';
import { Button, Collapse, Form } from 'react-bootstrap';
import { Expense, CreateExpensePayload, UpdateExpensePayload } from '../../types/Expense';
import { getExpense, listExpenses, createExpense, updateExpense, deleteExpense } from '../../apis/expenseApis';
import { formatDateTime } from '../../utils/dateTime';


const ExpensePage: React.FC = () => {
    const [expenses, setExpenses] = useState<Expense[]>([]);
    const [name, setName] = useState<string>('');
    const [amount, setAmount] = useState<number>(0);
    const [date, setDate] = useState<string>('');
    const [category, setCategory] = useState('food');
    const [search, setSearch] = useState<string>('');
    const [filterDate, setFilterDate] = useState<string>('');
    const [open, setOpen] = useState(false);
    const [editId, setEditId] = useState<number | null>(null);
    const [minDate, setMinDate] = useState('');
    const [maxDate, setMaxDate] = useState('');
    
    const fetchExpenses = async () => {
        const data = await listExpenses(search, filterDate);
        setExpenses(data ?? []);
    };
    
    useEffect(() => {
        fetchExpenses();
    }, []);

    useEffect(() => {
        const today = new Date();
        const oneYearAgo = new Date(today.getFullYear() - 1, today.getMonth(), today.getDate());
        const formattedMinDate = oneYearAgo.toISOString().split('T')[0];
        const formattedMaxDate = today.toISOString().split('T')[0];
        console.log(formattedMinDate, formattedMaxDate);
        setMinDate(formattedMinDate);
        setMaxDate(formattedMaxDate);
    }, []);
    
    const handleCreate = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const expenseName = formData.get('expense-name') as string;
        const amountValue = Number.parseInt(formData.get('expense-amount') as string);
        const dateValue = formData.get('expense-date') as string;
        const categoryValue = formData.get('expense-category') as string;

        const payload: CreateExpensePayload = {
            title: expenseName,
            amount: amountValue,
            category: categoryValue,
            updater: 'aloha',
            // 時間要轉換為包含時區的 ISO 8601 格式
            createdAt: new Date(dateValue).toISOString(),
        };
        try {
            const res = await createExpense(payload);
            setExpenses([...expenses, res]);
            resetForm();
        } catch (error) {
            console.error(error);
        }
    };
    
    const handleUpdate = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (editId === null) return;

        const formData = new FormData(e.currentTarget);
        const expenseName = formData.get('expense-name') as string;
        const amountValue = Number.parseInt(formData.get('expense-amount') as string);
        const categoryValue = formData.get('expense-category') as string;

        const payload: UpdateExpensePayload = {
            title: expenseName,
            amount: amountValue,
            category: categoryValue,
            updater: 'aloha',
        };

        try {
            const updatedExpense = await updateExpense(editId, payload);
            setExpenses(expenses.map(expense => expense.id === editId ? updatedExpense : expense));
            setEditId(null);
            setOpen(false);
        } catch (error) {
            console.error(error);
        }
    };
    
    const handleDelete = async (id: number) => {
        try {
            await deleteExpense(id);
            setExpenses(expenses.filter(expense => expense.id !== id));
        } catch (error) {
            console.error(error);
        }
    };
    
    const handleSearch = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        const searchValue = formData.get('search-input-name') as string;
        const dateValue = formData.get('search-input-date') as string;
        console.log(searchValue, dateValue);

        const data = await listExpenses(searchValue, dateValue);
        console.log(data);
        setExpenses(data ?? []);
    };
    
    const handleFilter = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
       if (name === 'search-input-name') {
           setSearch(value);
       } else if (name === 'search-input-date') {
           setFilterDate(value);
       }
    };
    
    const handleEdit = (expense: Expense) => {
        setName(expense.title);
        setAmount(expense.amount);
        setCategory(expense.category);
        setEditId(expense.id);
        setOpen(true);
    };

    const handleDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const selectedDate = e.target.value;
        const oneYearAgo = new Date(new Date().setFullYear(new Date().getFullYear() - 1));
        const selectedDateObj = new Date(selectedDate);
        if (selectedDateObj < oneYearAgo || selectedDateObj > new Date()) {
            setDate(''); // 如果日期不在范围内，重置日期
        } else {
            setDate(selectedDate);
        }
    };
    
    const handleToggle = () => {
        setOpen(!open);
        if (!open) {
            resetForm();
        }
    };

    const resetForm = () => {
        setName('');
        setAmount(0);
        setDate('');
        setCategory('food');
        setEditId(null);
    };
    
    return (
        <div className="container">
            <div className="content">
                <h1 className='title'>支出管理</h1>
                <Form className='form' onSubmit={handleSearch}>
                    <Form.Group className='form-group'>
                        <Form.Control name='search-input-name' className='form-control-input' type="text" placeholder="搜尋支出" value={search} onChange={handleFilter} />
                        <Form.Control name='search-input-date' className='form-control-date' type="date" value={filterDate} onChange={handleFilter} />
                        <Form.Control className='button form-control-btn' type="submit" value="搜尋" />
                    </Form.Group>
                </Form>
                <Button
                    className='button'
                    variant="primary"
                    onClick={handleToggle}
                >{open ? (editId === null ? '取消新增' : '取消編輯') : '新增支出'}</Button>
                {open && (
                    <div id="example-collapse-text">
                        <Form onSubmit={editId === null ? handleCreate : handleUpdate}>
                            <Form.Group className="form-group">
                                <Form.Label htmlFor="name">名稱</Form.Label>
                                <Form.Control
                                    name="expense-name"
                                    type="text"
                                    className="form-control"
                                    id="expense-name"
                                    value={name}
                                    onChange={(e) => setName(e.target.value)}
                                />
                            </Form.Group>
                            <Form.Group className="form-group">
                                <Form.Label htmlFor="amount">金額</Form.Label>
                                <Form.Control
                                    name="expense-amount"
                                    type="number"
                                    className="form-control"
                                    id="expense-amount"
                                    min={0}
                                    value={amount}
                                    onChange={(e) => setAmount(Number.parseInt(e.target.value))}
                                />
                            </Form.Group>
                            <Form.Group className="form-group">
                                <Form.Label htmlFor="date">日期</Form.Label>
                                <Form.Control
                                    name="expense-date"
                                    type="date"
                                    className="form-control"
                                    id="expense-date"
                                    max={maxDate}
                                    min={minDate}
                                    value={date}
                                    onChange={handleDateChange}
                                />
                            </Form.Group>
                            <Form.Group className="form-group">
                                <Form.Label htmlFor="category">類別</Form.Label>
                                <Form.Control
                                    name="expense-category"
                                    as="select"
                                    className="form-control"
                                    id="category"
                                    value={category}
                                    onChange={(e) => setCategory(e.target.value)}
                                >
                                    <option value="food">食物</option>
                                    <option value="clothing">衣物</option>
                                    <option value="housing">住房</option>
                                    <option value="transportation">交通</option>
                                </Form.Control>
                            </Form.Group>
                            <Button type="submit" className="btn btn-primary">
                                {editId === null ? '確認' : '更新'}
                            </Button>
                        </Form>
                    </div>
                )}
                <table className='table'>
                    <thead>
                        <tr>
                            <th>名稱</th>
                            <th>金額</th>
                            <th>類別</th>
                            <th>日期</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        {expenses.map((expense) => (
                            <tr key={expense.id}>
                                <td>{expense.title}</td>
                                <td>{expense.amount}</td>
                                <td>{expense.category}</td>
                                <td>{formatDateTime(expense.date)}</td>
                                <td>
                                    <Button className='button td-btn' variant="primary" onClick={() => handleEdit(expense)}>編輯</Button>
                                    <Button className='button td-btn' variant="danger" onClick={() => handleDelete(expense.id)}>刪除</Button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default ExpensePage;