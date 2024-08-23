document.addEventListener('DOMContentLoaded', function() {
    fetchTasks();
});

function fetchTasks() {
    fetch('http://localhost:8080/tasks')
        .then(response => response.json())
        .then(data => {
            const tasksDiv = document.getElementById('tasks');
            tasksDiv.innerHTML = '';
            data.forEach(task => {
                const taskDiv = document.createElement('div');
                taskDiv.className = 'task';
                taskDiv.innerHTML = `
                    <h3>${task.Title}</h3>
                    <p>${task.Description}</p>
                    <p>Ответственный: ${task.UserName}</p>
                    <p>Приоритет: ${task.TaskPriority.Priority}</p>
                    <p>Статус: ${task.IsDone ? 'Сделано' : 'Не сделано'}</p>
                    <button onclick="deleteTask(${task.ID})">Удалить</button>
                    <button onclick="markAsDone(${task.ID})">Отметить как выполненное</button>
                `;
                tasksDiv.appendChild(taskDiv);
            });
        })
        .catch(error => console.error('Ошибка:', error));
}

function addTask() {
    const title = document.getElementById('title').value;
    const description = document.getElementById('description').value;
    const userName = document.getElementById('userName').value;
    const priority = document.getElementById('priority').value;

    const task = {
        Title: title,
        Description: description,
        UserName: userName,
        TaskPriorityID: parseInt(priority),
    };

    fetch('http://localhost:8080/tasks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(task),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Успех:', data);
            fetchTasks();
        })
        .catch(error => console.error('Ошибка:', error));
}

function deleteTask(id) {
    fetch(`http://localhost:8080/tasks/${id}`, {
        method: 'DELETE',
    })
        .then(() => {
            fetchTasks();
        })
        .catch(error => console.error('Ошибка:', error));
}

function markAsDone(id) {
    fetch(`http://localhost:8080/tasks/${id}/done`, {
        method: 'PUT',
    })
        .then(() => {
            fetchTasks();
        })
        .catch(error => console.error('Ошибка:', error));
}