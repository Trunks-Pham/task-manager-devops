import React, { useEffect, useState } from 'react';
import TaskList from './components/TaskList';
import TaskForm from './components/TaskForm';

function App() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then(res => res.json())
      .then(data => setTasks(data));
  }, []);

  return (
    <div>
      <h1>Quản Lý Công Việc</h1>
      <TaskForm />
      <TaskList tasks={tasks} />
    </div>
  );
}

export default App;
