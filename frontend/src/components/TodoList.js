import React, { useEffect, useState } from "react";
import ApiServices from "../services/api-services";
import "./TodoList.css";

export default function TodoList(props) {
  const [todoList, setTodoList] = useState([]);
  const [data, setData] = useState("");

  useEffect(() => {
    //List all the items that are posted
    ApiServices.getAllData()
      .then((res) => {
        setTodoList(res.data);
      })
      .catch((err) => console.error(err));
  }, []);

  //Function to remove the itemn
  function remove(id) {
    ApiServices.removeData(id)
      .then((res) => {
        console.log(res.data);
        const myalldata = todoList.filter((item) => item.id !== id);
        setTodoList(myalldata);
      })
      .catch((err) => console.error(err));
  }

  //pass id as props to update the specific item
  function Update(id) {
    console.log(id);
    props.history.push("/Update/" + id);
  }

  //function to submit the item to add on the list
  function submit(e) {
    e.preventDefault();
    ApiServices.postData({ title: data }).then((res) => {
      const mydata = [...todoList, res.data];
      setTodoList(mydata);
    });
  }

  function handle(e) {
    setData(e.target.value);
  }

  //map the array response to display all the items
  const display = todoList.map((item) => (
    <tr key={item.id}>
      <td>{item.title}</td>
      <td>{new Date(item.CreatedAt).toUTCString()}</td>

      <td>
        <button onClick={() => Update(item.id)} className="btn btn-success">
          Update
        </button>
      </td>
      <td>
        <button onClick={() => remove(item.id)} className="btn btn-danger">
          Delete
        </button>
      </td>
    </tr>
  ));

  //sort data to display latest created data first
  const sortData = display.sort((a, b) => {
    if (new Date(a.createdAt) > new Date(b.createdAt)) {
      return 1;
    }
    return -1;
  });

  return (
    <div className="conatainer">
      <form onSubmit={(e) => submit(e)}>
        <div className="form-group" style={{ paddingBottom: "1rem" }}>
          <label htmlFor="items"> Add Todo Items</label>
          <input
            value={data}
            type="text"
            name="items"
            id="items"
            className="form-control"
            placeholder="add todo items here..."
            onChange={(e) => handle(e)}
          />
        </div>
        <button className="btn btn-primary" style={{ marginBottom: "1rem" }}>
          Submit
        </button>
      </form>
      <table className="table table-stripped">
        <tbody>{sortData}</tbody>
      </table>
    </div>
  );
}
