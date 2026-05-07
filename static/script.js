async function addTask() {
 
    const title = document.getElementById("title").value
 
    const description = document.getElementById("description").value
 
    await fetch("/tasks", {
 
        method: "POST",
 
        headers: {
            "Content-Type": "application/json"
        },
 
        body: JSON.stringify({
 
            title: title,
            description: description
        })
    })
 
    loadTasks()
}
 
async function deleteTask(id) {
 
    await fetch(`/tasks/${id}`, {
 
        method: "DELETE"
    })
 
    loadTasks()
}
 
async function updateTask(id) {
 
    const newTitle = prompt("Enter New Title")
 
    const newDescription = prompt("Enter New Description")
 
    await fetch(`/tasks/${id}`, {
 
        method: "PUT",
 
        headers: {
            "Content-Type": "application/json"
        },
 
        body: JSON.stringify({
 
            title: newTitle,
            description: newDescription
        })
    })
 
    loadTasks()
}
 
async function loadTasks() {
 
    const response = await fetch("/tasks")
 
    const data = await response.json()
 
    const taskList = document.getElementById("taskList")
 
    taskList.innerHTML = ""
 
    data.data.forEach(task => {
 
        const li = document.createElement("li")
 
        li.innerHTML = `
 
            <b>${task.title}</b>
 
            <br>
 
            ${task.description}
 
            <br><br>
 
            <button onclick="updateTask(${task.id})">
                Edit
            </button>
 
            <button onclick="deleteTask(${task.id})">
                Delete
            </button>
 
            <hr>
        `
 
        taskList.appendChild(li)
    })
}
 
loadTasks()