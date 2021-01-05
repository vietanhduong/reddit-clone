import React from "react";
import { browserHistory } from "utils/history";
import { Router, Route, Switch, Redirect } from "react-router-dom";
import Home from "components/home";
import Notfound from "components/404";
import Login from "components/login";
import Create from "components/create";

const App = () => {
  return (<Router history={browserHistory}>
    <Switch>
    <Route exact path="/create" component={Create} />
      <Route exact path="/login" component={Login} />
      <Route exact path="/404" component={Notfound} />
      <Route exact path="/" component={Home} />
      <Redirect to="/404" />
    </Switch>
  </Router>);
}

export default App;
