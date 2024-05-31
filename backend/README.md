## Assignment
* Please implement an Expense API that can perform CRUD of an "Expense"
  * An "Expense" should include 
    * Title: what is this expense for (string)
    * Amount: how much is this expense (number, not negative)
    * Date: when is this expense (date, max 1 year from past to now)
    * Category: what kind of expense is this (string, can only be "food", clothing", "housing", "transportation")
  * An InMemory database is enough for this assignment
  * The API should include
    * Create an expense
    * List all expenses
    * Delete an expense
    * Update an expense
    * Search expenses by title (can be a ambiguous search)
    * Filter expenses by date range (date range should be within 30 days)
    * login
  * Each user can only see their own expenses