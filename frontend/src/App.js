import { BrowserRouter as Router, Route } from "react-router-dom";
import TodoList from "./components/TodoList";
import "./App.css";
import UpdateItem from "./components/UpdateItem";
function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/" exact component={TodoList} />
        <Route path="/Update/:id" exact component={UpdateItem} />
      </Router>
    </div>
  );
}

export default App;
