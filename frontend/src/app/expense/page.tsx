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
import React, { useEffect, useState } from 'react';
import { Button, Collapse, Form } from 'react-bootstrap';
import { Expense } from '../../types/Expense';
import { getExpenses, listExpenses, createExpense, updateExpense, deleteExpense } from '../../apis/expenseApis';
import { formatDateTime } from '../../utils/dateTime';


const ExpensePage: React.FC = () => {
    const [expenses, setExpenses] = useState<Expense[]>([{ expenseId: 1, title: 'test', amount: 100, date: '2021-10-01', updater: 'test' }, { expenseId: 2, title: 'test2', amount: 200, date: '2021-10-02', updater: 'test2' }]);
    const [name, setName] = useState<string>('');
    const [amount, setAmount] = useState<number>(0);
    const [date, setDate] = useState<string>('');
    const [search, setSearch] = useState<string>('');
    const [filterDate, setFilterDate] = useState<string>('');
    const [open, setOpen] = useState(false);
    const [editId, setEditId] = useState<number | null>(null);
    
    const fetchData = async () => {
        // const data = await getExpenses();
        // setExpenses(data);
        return;
    };
    
    useEffect(() => {
        // fetchData();
    }, []);
    
    const handleCreate = async () => {
        // await createExpense({ name, amount, date });
        // fetchData();
        // setShow(false);
        return;
    };
    
    const handleUpdate = async (id: number) => {
        // await updateExpense(id, { name, amount, date });
        // fetchData();
        // setShow(false);
        return;
    };
    
    const handleDelete = async (id: number) => {
        // await deleteExpense(id);
        // fetchData();
        return;
    };
    
    const handleSearch = (e: React.ChangeEvent<HTMLInputElement>) => {
        // setSearch(e.target.value);
        return;
    };
    
    const handleFilter = (e: React.ChangeEvent<HTMLInputElement>) => {
        // setFilterDate(e.target.value);
        return;
    };
    
    const handleEdit = (id: number) => {
        // const expense = expenses.find((expense) => expense.id === id);
        // if (expense) {
        // setName(expense.name);
        // setAmount(expense.amount);
        // setDate(expense.date);
        // setEditId(id);
        // setShow(true);
        // }
        return;
    };
    
    const handleToggle = () => {
        setOpen(!open);
    };
    
    return (
        <div className="container">
            <div className="content">
                <h1 className='title'>支出管理</h1>
                <Form className='form'>
                    <Form.Group className='form-group'>
                        <Form.Control className='form-control-input' type="text" placeholder="搜尋支出" value={search} onChange={handleSearch} />
                        <Form.Control className='form-control-date' type="date" value={filterDate} onChange={handleFilter} />
                    </Form.Group>
                </Form>
                <Button
                    className='button'
                    variant="primary"
                    onClick={handleToggle}
                >新增支出</Button>
                {open && (
                    <div id="example-collapse-text">
                        <form>
                            <div className="form-group">
                                <label htmlFor="name">名稱</label>
                                <input type="text" className="form-control" id="name" value="" onChange={(e) => setName(e.target.value)} />
                            </div>
                            <div className="form-group">
                                <label htmlFor="amount">金額</label>
                                <input type="number" className="form-control" id="amount" value="" onChange={(e) => setAmount(Number.parseInt(e.target.value))} />
                            </div>
                            <div className="form-group">
                                <label htmlFor="date">日期</label>
                                <input type="date" className="form-control" id="date" value="" onChange={(e) => setDate(e.target.value)} />
                            </div>
                            <button type="submit" className="btn btn-primary" onClick={handleCreate}>確認</button>
                        </form>
                    </div>
                )}
                <table className='table'>
                    <thead>
                        <tr>
                            <th>名稱</th>
                            <th>金額</th>
                            <th>日期</th>
                            <th>操作</th>
                        </tr>
                    </thead>
                    <tbody>
                        {expenses.map((expense) => (
                            <tr key={expense.expenseId}>
                                <td>{expense.title}</td>
                                <td>{expense.amount}</td>
                                <td>{formatDateTime(expense.date)}</td>
                                <td>
                                    <Button className='button td-btn' variant="primary" onClick={() => handleEdit(expense.expenseId)}>編輯</Button>
                                    <Button className='button td-btn' variant="danger" onClick={() => handleDelete(expense.expenseId)}>刪除</Button>
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