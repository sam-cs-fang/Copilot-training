
// 在這裡實作 app/expense/page.tsx 中要用到的 formatDate 函數
// 這個函數會將日期字串轉換成 yyyy-mm-dd 格式
// 例如將 '2021-10-01T00:00:00.000Z' 轉換成 '2021-10-01'
// 這個函數會被用在 ExpenseTable 的日期欄位


export const formatDateTime = (date: string): string => {
    return date.split('T')[0];
};
