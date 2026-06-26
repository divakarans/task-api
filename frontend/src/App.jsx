import { useState, useEffect } from "react";
import "./index.css";

function Modal({tasks,onClose,deleteTask,setTitle,setDescription,setEditingId,setStatus}){
    return(
      <div className="overlay">
        <div className="modal">
          <button onClick={onClose} className="close-btn">
            ✖
          </button>
          <h4>Task List</h4>
          {tasks.map(task=>(
            <div key={task.id}>
              <h4>[Task Id:{task.id}] Title :{task.title}</h4>
              <p>Description :{task.description}</p>
              <p>Status :{task.status}</p>
              <button className="edit-btn" 
                onClick={()=>{
                  setTitle(task.title);
                  setDescription(task.description);
                  setStatus(task.status);
                  setEditingId(task.id);
                  onClose()
                }}
              >
                Edit</button>
              <button className="del-btn" onClick={()=>deleteTask(task.id)}>Delete</button>
              <hr/>
            </div>
          ))}
        </div>
      </div>
    );
}

function App() {

  const [title, setTitle]=useState("");
  const [description, setDescription]=useState("");
  const [status, setStatus]=useState("pending");
  const [tasks, setTasks] = useState([]);
  const [showModal,setshowModal]=useState(false);
  const [editingId, setEditingId]=useState(null);

  const fetchTasks = async () => {
    fetch("http://localhost:8080/tasks")
      .then(response => response.json())
      .then(data => {
        setTasks(data);
      });
  };

  const addTask = async () => {
    const newtask={
      title,
      description,
      status
    }
    await fetch("http://localhost:8080/tasks",{
      method:"POST",
      headers:{
        "Content-Type":"application/json"
      },
      body:JSON.stringify(newtask),
    })
    .then(response=>{
      if(!response.ok){
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then(json=>console.log("success:",json))
    .catch(error=>console.error("Error:",error))

    await fetchTasks();
    setTitle("");
    setDescription("");
    setStatus("pending");

  };

  const editTask = async (id) => {
    try{
      const edittask={
        title,
        description,
        status
      }
      const response=await fetch(`http://localhost:8080/tasks/${id}`,{
        method:"PUT",
        headers:{
          "Content-Type":"application/json"
        },
        body:JSON.stringify(edittask),
      })
      if(!response.ok){
        throw new Error("failed to update task")
      }
      await fetchTasks()
      setTitle("");
      setDescription("");
      setStatus("pending");
      setEditingId(null);
    } catch(err){
      console.log(err)
    }
  };

  const deleteTask = async (id) => {
    try{
      const response= await fetch(`http://localhost:8080/tasks/${id}`,{
        method:"DELETE",
      })

      if (!response.ok){
        throw new Error("failed to delete task")
      }
      await fetchTasks()
    } catch(err){
      console.log(err)
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  return (
    <div className="container">
      <h3>Welcome to To-Do App</h3>

      <input
        type="text"
        placeholder="Enter Title"
        value={title}
        onChange={(e)=>setTitle(e.target.value)}
      />

      <br />

      <input
        type="text"
        placeholder="Enter Description"
        value={description}
        onChange={(e)=>setDescription(e.target.value)}
      />

      <br />

      <select value={status} onChange={(e)=>setStatus(e.target.value)}>
        <option>pending</option>
        <option>in_progress</option>
        <option>done</option>
      </select>

      <button className="Addbutton" 
        onClick={editingId ? ()=> editTask(editingId) : addTask}>
        {editingId ? "Edit Task" : "Add Task"}
      </button>

      <button className="Viewbutton" 
        onClick={()=>setshowModal(true)}>
        View Tasks
      </button>
      
      {showModal && (
        <Modal  tasks={tasks}
        onClose={()=>setshowModal(false)}
        deleteTask={deleteTask}
        setTitle={setTitle}
        setDescription={setDescription}
        setStatus={setStatus}
        setEditingId={setEditingId}
      />)}
    </div>
  );
}

export default App;