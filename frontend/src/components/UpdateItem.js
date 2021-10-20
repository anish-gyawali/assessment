import React, { useState, useEffect } from "react";
import ApiServices from "../services/api-services";
import "./TodoList.css";

export default function UpdateItem(props) {
  const [data, setData] = useState("");
  const [date, setDate] = useState("");

  //get the id of item that is needed to update
  useEffect(() => {
    const id = props.match.params.id;
    ApiServices.getData(id)
      .then((res) => {
        console.log(res.data);
        setData(res.data);
        setDate(res.data.CreatedAt);
      })
      .catch((err) => console.error(err));
  }, []);

  //submit the item after updating it
  function submit(e) {
    e.preventDefault();
    const id = props.match.params.id;
    ApiServices.updateData(id, data, date).then((res) => {
      props.history.push("/");
    });
  }

  function handle(e) {
    setData(e.target.value);
  }

  return (
    <div className="conatainer">
      <form onSubmit={(e) => submit(e)}>
        <div className="form-group" style={{ paddingBottom: "1rem" }}>
          <label htmlFor="items"> Add Todo Items</label>
          <input
            defaultValue={data.title}
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
    </div>
  );
}
