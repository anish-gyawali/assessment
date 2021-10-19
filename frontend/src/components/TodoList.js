import React from "react";
import "./TodoList.css";
export default function TodoList() {
  function submit(e) {}
  function handle(e) {}

  return (
    <div className="conatainer">
      <form onSubmit={(e) => submit(e)}>
        <div className="form-group" style={{ paddingBottom: "1rem" }}>
          <label htmlFor="items"> Add Todo Items</label>
          <input
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
    </div>
  );
}
