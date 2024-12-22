function TaskList({ tasks }) {
  return (
    <ul>
      {tasks.map(task => (
        <li key={task.id}>
          {task.title} - {task.done ? 'Done' : 'Pending'}
        </li>
      ))}
    </ul>
  );
}

export default TaskList;
