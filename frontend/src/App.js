import { BrowserRouter as Router, Route } from "react-router-dom";
import TodoList from "./components/TodoList";
import "./App.css";
function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/" exact component={TodoList} />
      </Router>
    </div>
  );
}

export default App;
