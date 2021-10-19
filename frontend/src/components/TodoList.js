import React, { useEffect, useState } from "react";
import ApiServices from "../services/api-services";
import "./TodoList.css";
export default function TodoList() {
  const [todoList, setTodoList] = useState([]);
  const [data, setData] = useState("");
  useEffect(() => {
    ApiServices.getAllData()
      .then((res) => {
        console.log(res.data);
        setTodoList(res.data);
      })
      .catch((err) => console.error(err));
  }, []);
  function submit(e) {
    e.preventDefault();
    ApiServices.postData({ title: data }).then((res) => {
      console.log(res.data);
      const mydata = [...todoList, res.data];
      setTodoList(mydata);
    });
  }
  function handle(e) {
    setData(e.target.value);
  }

  const display = todoList.map((item) => (
    <tr key={item.id}>
      <td>{item.title}</td>
      <td>{new Date(item.CreatedAt).toUTCString()}</td>

      <td>
        <button className="btn btn-success">Update</button>
      </td>
      <td>
        <button className="btn btn-Danger">Delete</button>
      </td>
    </tr>
  ));
  return (
    <div className="conatainer">
      <form onSubmit={(e) => submit(e)}>
        <div className="form-group" style={{ paddingBottom: "1rem" }}>
          <label htmlFor="items"> Add Todo Items</label>
          <input
            defaultValue={data}
            type="text"
            name="items"
            id="items"
            className="form-control"
            placeholder="add todo items here..."
            onChange={(e) => handle(e)}
          />
        </div>
        <button className="btn btn-primary">Submit</button>
      </form>
      <table className="table table-stripped">
        <tbody>{display}</tbody>
      </table>
    </div>
  );
}
