import React, {useState, useEffect} from "react";

function ToDoList(){
    const[todos,setTodos] = useState([]);
    const[newTodo,setNewTodo] = useState("");

    useEffect(() => {
        fetch("/todos", {
            method: "GET"
        })
        .then((response) => response.json())
        .then((data) => setTodos(data));
    }, [])

    const addTodo = () => {
        fetch("/todo", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({title: newTodo, completed: false})
        })
        .then((response) => response.json())
        .then((data) => setTodos([...todos,data]))
        setNewTodo("")
    }

    const toggleComplete = (todoId, status) => {
        fetch(`/todo/${todoId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ completed: status })
        })
        .then(response => response.json())
        .then(updatedTodo => {
            setTodos(prevTodos => 
                prevTodos.map(t => t.Id === todoId ? updatedTodo : t));
        })
        .catch(error => {
            console.error("Error updating todo:", error);
        });
    };

    const updateTodo = (todoId, todoTitle) => {
        fetch(`todo/${todoId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ title: todoTitle })
        })
        .then(response => response.json())
        .then(updatedTodo => {
            setTodos(prevTodos => prevTodos.map(t => t.Id === todoId ? updatedTodo : t))
        })
        .catch(error => {
            console.error("Error on updating todo: ", error)
        })

    }

    const deleteTodo = (todoId) => {
        fetch(`todo/${todoId}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            }
        })
        .then(response => response.json())
        .then(data  => {
            setTodos(data)
        })
        .catch(error => {
            console.error("Error deleting todo: ", error)
        })
            
        
    }

    return (
        <div style={styles.container}>
            <h1 style={styles.header}>Todo List</h1>
            <div style={styles.inputContainer}>
                <input
                    value={newTodo}
                    onChange={(e) => setNewTodo(e.target.value)}
                    placeholder="Add new todo ..."
                    style={styles.input}
                />
                <button onClick={addTodo} style={styles.addButton}>Add</button>
            </div>
            <ul style={styles.list}>
                
                {todos.map((todo) => (
                    <li key={todo.Id} style={styles.listItem}>
                        <input
                            type="checkbox"
                            checked={todo.completed}
                            onChange={() => toggleComplete(todo.Id, !todo.completed)}
                            style={styles.checkbox}
                        />
                        <input
                            value={todo.title}
                            onChange={(e) => updateTodo(todo.Id, e.target.value)}
                            style={styles.todoInput}
                        />
                        <button 
                        onClick={() => deleteTodo(todo.Id)} 
                        style={styles.deleteButton}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
}

const styles = {
    container: {
        maxWidth: "400px",
        margin: "60px auto",
        padding: "20px",
        backgroundColor: "#f4f4f4",
        borderRadius: "8px",
        boxShadow: "0 0 10px rgba(0, 0, 0, 0.1)"
    },
    header: {
        textAlign: "center",
        color: "#333",
    },
    inputContainer: {
        display: "flex",
        justifyContent: "space-between",
        marginBottom: "20px",
    },
    input: {
        flex: "1",
        padding: "10px",
        marginRight: "10px",
        borderRadius: "4px",
        border: "1px solid #ccc",
    },
    addButton: {
        padding: "10px 20px",
        backgroundColor: "#28a745",
        color: "#fff",
        border: "none",
        borderRadius: "4px",
        cursor: "pointer",
    },
    list: {
        listStyleType: "none",
        padding: "0",
    },
    listItem: {
        display: "flex",
        alignItems: "center",
        marginBottom: "10px",
        backgroundColor: "#fff",
        padding: "10px",
        borderRadius: "4px",
        boxShadow: "0 1px 3px rgba(0, 0, 0, 0.1)",
    },
    checkbox: {
        marginRight: "10px",
    },
    todoInput: {
        flex: "1",
        padding: "5px",
        borderRadius: "4px",
        border: "1px solid #ccc",
        marginRight: "10px",
    },
    deleteButton: {
        padding: "5px 10px",
        backgroundColor: "#dc3545",
        color: "#fff",
        border: "none",
        borderRadius: "4px",
        cursor: "pointer",
    },
};

export default ToDoList;
