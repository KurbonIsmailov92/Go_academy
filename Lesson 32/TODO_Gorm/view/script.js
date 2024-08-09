document.getElementById('todo-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const taskInput = document.getElementById('new-task');
    const taskText = taskInput.value.trim();

    if (taskText !== '') {
        addTaskToList(taskText);
        taskInput.value = '';
    }
});

function addTaskToList(taskText) {
    const taskList = document.getElementById('task-list');

    const li = document.createElement('li');
    li.textContent = taskText;

    const deleteButton = document.createElement('button');
    deleteButton.textContent = 'Delete';
    deleteButton.className = 'delete';
    deleteButton.onclick = function() {
        taskList.removeChild(li);
    };

    const doneButton = document.createElement('button');
    doneButton.textContent = 'Done';
    doneButton.className = 'done';
    doneButton.onclick = function() {
        li.classList.toggle('completed');
    };

    li.appendChild(doneButton);
    li.appendChild(deleteButton);

    taskList.appendChild(li);

    fetch('http://localhost:8080/tasks')
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error('Error:', error));

// Пример создания новой задачи
    fetch('http://localhost:8080/tasks', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            title: 'Новая задача',
            description: 'Описание новой задачи',
            task_priority_id: 2,
            user_name: 'Vasya',
        }),
    })
        .then(response => response.json())
        .then(data => console.log(data))
        .catch(error => console.error('Error:', error));
}
